// Package uci chess uci.go (partly) implements the UCI protocol for communicating with
// chess engines.
package uci

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"unicode"

	chess "github.com/notnil/chess"
)

// CommunicationTimeout is the time to wait for a response from the engine. If
// the engine fails to respond, it is terminated.
var CommunicationTimeout time.Duration = 15 * time.Second

// process implements io.Closer for a running process.
type process struct {
	cmd *exec.Cmd
}

// Close waits for the process to stop.
func (p *process) Close() error {
	if p.cmd == nil {
		return nil
	}
	waited := make(chan bool)
	go func() {
		p.cmd.Wait()
		waited <- true
	}()
	select {
	case <-waited:
		// nothing
	case <-time.After(CommunicationTimeout):
		p.cmd.Process.Kill()
		<-waited
	}
	p.cmd = nil
	return nil
}

// UciEngine represents a running UCI engine.
// and takes the UciEngine interface declared in uciengine.go
type uciEngine struct {
	cmdc chan<- interface{}
	errc <-chan error
}

var _ UciEngine = &uciEngine{}

// Run starts an engine executable, with the given arguments. If logger is not
// nil, it will be used to log all communication to and from the engine.
func Run(exe string, args []string, logger *log.Logger) (UciEngine, error) {
	cmd := exec.Command(exe, args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("start engine: %s", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("start engine: %s", err)
	}
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("%s %v: %s", exe, args, err)
	}
	return initialise(stdout, stdin, &process{cmd}, logger)
}

func initialise(stdout io.Reader, stdin io.Writer, proc io.Closer, logger *log.Logger) (*uciEngine, error) {
	var (
		cmdc  = make(chan interface{})
		errc  = make(chan error)
		linec = make(chan string)
	)
	c := &comm{
		cmdc:    cmdc,
		errc:    errc,
		linec:   linec,
		stdin:   stdin,
		process: proc,
		log:     logger,
	}
	go c.run()
	go readLines(stdout, linec, &c.readError)

	e := &uciEngine{
		cmdc: cmdc,
		errc: errc,
	}
	if err := e.Send("uci"); err != nil {
		return nil, err
	}
	return e, nil
}

// Send sends a command to the engine.
func (e *uciEngine) Send(cmd string) error {
	e.cmdc <- cmd
	return <-e.errc
}

// Stop implements engine.Engine.
func (e *uciEngine) Stop() {
	e.Send("stop")
}

// Ping implements engine.Engine.
func (e *uciEngine) Ping() error {
	return e.Send("isready")
}

// Quit implements engine.Engine.
func (e *uciEngine) Quit() {
	e.Send("quit")
	close(e.cmdc)
}

// Search

// SetPosition implements engine.Engine.
func (e *uciEngine) SetPosition(position *chess.Position) {
	e.Send("ucinewgame")
	fen, err := position.MarshalText()
	if err != nil {
		// nothing?
	}
	fenstring := string(fen)
	e.Send(fmt.Sprintf("position fen %s", fenstring))
	e.cmdc <- position
	<-e.errc
}

// Search implements engine.Engine.
func (e *uciEngine) Search() <-chan UciEngineInfo {
	return e.search("go infinite")
}

// SearchDepth implements engine.Engine.
func (e *uciEngine) SearchDepth(depth int) <-chan UciEngineInfo {
	cmd := "go depth %d"
	return e.search(fmt.Sprintf(cmd, depth))
}

// SearchTime implements engine.Engine.
func (e *uciEngine) SearchTime(t time.Duration) <-chan UciEngineInfo {
	cmd := "go movetime %d"
	return e.search(fmt.Sprintf(cmd, t/time.Millisecond))
}

// SearchClock implements engine.Engine.
func (e *uciEngine) SearchClock(wtime, btime, winc, binc time.Duration, movesToGo int) <-chan UciEngineInfo {
	return e.search(fmt.Sprintf(
		"go wtime %d btime %d winc %d binc %d movestogo %d",
		wtime/time.Millisecond, btime/time.Millisecond,
		winc/time.Millisecond, binc/time.Millisecond,
		movesToGo))
}

func (e *uciEngine) search(cmd string) <-chan UciEngineInfo {
	infoc := make(chan UciEngineInfo, 1)
	if err := e.initSearch(cmd, infoc); err != nil {
		infoc <- Info{err: err}
		close(infoc)
	}
	return infoc
}

func (e *uciEngine) initSearch(cmd string, infoc chan UciEngineInfo) error {
	// Sync to ensure that no debris is sent on the Info channel.
	e.Ping()
	// Tell the communicator to send info lines on infoc.
	e.cmdc <- infoc
	if err := <-e.errc; err != nil {
		return err
	}
	// Start the search.
	e.Send(cmd)
	return nil
}

// Options implements engine.Engine.
func (e *uciEngine) Options() map[string]UciEngineOption {
	optc := make(chan map[string]UciEngineOption)
	e.cmdc <- optc
	if err := <-e.errc; err != nil {
		return nil
	}
	return <-optc
}

// Communicator.

type comm struct {
	cmdc      chan interface{}           // request channel
	errc      chan error                 // response channel
	err       error                      // error state of the communication
	linec     <-chan string              // engine output lines
	infoc     chan<- UciEngineInfo       // for sending out "info ..." lines
	position  *chess.Position            // position being searched
	process   io.Closer                  // the thing to close on error
	stdin     io.Writer                  // for sending commands
	log       *log.Logger                // communication log
	name      string                     // engine name
	author    string                     // engine author(s)
	options   map[string]UciEngineOption // engine options
	readError error                      // error returned by readLines
}

func readLines(stdout io.Reader, linec chan<- string, perr *error) {
	bufrd := bufio.NewReader(stdout)
	for {
		line, isprefix, err := bufrd.ReadLine()
		for err == nil && isprefix {
			// ignore rest of line
			_, isprefix, err = bufrd.ReadLine()
		}
		if err != nil {
			*perr = err
			break
		}
		linec <- strings.TrimSpace(string(line))
	}
	close(linec)
}

func timeoutWrite(w io.Writer, line string) error {
	errc := make(chan error)
	go func() {
		_, err := fmt.Fprintln(w, line)
		errc <- err
	}()
	select {
	case err := <-errc:
		return err
	case <-time.After(CommunicationTimeout):
		return ErrTimeout
	}
	panic("unreachable")
}

func (c *comm) close(err error) {
	c.err = err
	c.process.Close()
	if c.infoc != nil {
		c.infoc <- Info{err: err}
		close(c.infoc)
		c.infoc = nil
	}
}

func (c *comm) run() {
	var timeout <-chan time.Time
	initialised := false
	c.options = make(map[string]UciEngineOption)

loop:
	select {
	case in, ok := <-c.cmdc:
		if !ok {
			return
		}
		errc := c.errc
		if c.err == nil {
			switch v := in.(type) {
			case string:
				if c.log != nil {
					c.log.Println(">", v)
				}
				c.err = timeoutWrite(c.stdin, v)
				switch {
				case c.err != nil:
					c.close(c.err)
				case v == "uci" || v == "isready" || v == "quit":
					timeout = time.After(CommunicationTimeout)
					errc = nil
				}
			case *chess.Position:
				c.position = v
			case chan UciEngineInfo:
				if c.position == nil {
					c.err = errors.New("SetPosition not called before search")
				} else {
					c.infoc = v
				}
			case chan map[string]UciEngineOption:
				errc <- nil
				errc = nil
				v <- c.options
			case chan string:
				v <- c.name
				v <- c.author
			}
		}
		if errc != nil {
			errc <- c.err
		}
	case line, ok := <-c.linec:
		if !ok {
			c.linec = nil
			if c.err == nil {
				c.close(c.readError)
			}
			if timeout != nil {
				c.errc <- c.err
				timeout = nil
			}
			break
		}
		if c.log != nil {
			log.Println("|", line)
		}
		switch field := tokenise(line); field.next() {
		case "id":
			switch field.next() {
			case "name":
				c.name = field.remainder()
			case "author":
				c.author = field.remainder()
			}
		case "option":
			if !initialised {
				c.parseOption(line)
			}
		case "uciok":
			if !initialised && timeout != nil {
				c.errc <- nil
				initialised = true
				timeout = nil
			}
		case "readyok":
			if initialised && timeout != nil {
				c.errc <- nil
				timeout = nil
			}
		case "info":
			if c.infoc != nil {
				c.infoc <- Info{Line: line, position: c.position}
			}
		case "bestmove":
			if c.infoc != nil {
				c.infoc <- Info{Line: line, position: c.position}
				close(c.infoc)
				c.infoc = nil
			}
		}
	case <-timeout:
		c.close(ErrTimeout)
		c.errc <- c.err
		timeout = nil
	}

	goto loop
}

func (c *comm) parseOption(line string) {
	var err error

	name, ok := fieldValue(line, "name", optionKeywords)
	if !ok {
		return
	}
	typ, ok := fieldValue(line, "type", optionKeywords)
	if !ok {
		return
	}
	def, haveDefault := fieldValue(line, "default", optionKeywords)

	opt := option{name, c.cmdc, c.errc}

	switch typ {
	case "string":
		c.options[name] = &sStringOption{
			option: opt,
			def:    def,
			value:  def,
		}
	case "check":
		defbool := false
		if haveDefault {
			if defbool, err = strconv.ParseBool(def); err != nil {
				defbool = false
			}
		}
		c.options[name] = &sBoolOption{
			option: opt,
			def:    defbool,
			value:  defbool,
		}
	case "spin":
		minint, maxint := 0, 0
		if min, ok := fieldValue(line, "min", optionKeywords); ok {
			if minint, err = strconv.Atoi(min); err != nil {
				minint = 0
			}
		}
		if max, ok := fieldValue(line, "max", optionKeywords); ok {
			if maxint, err = strconv.Atoi(max); err != nil {
				maxint = 0
			}
		}
		defint := minint
		if haveDefault {
			if defint, err = strconv.Atoi(def); err != nil {
				defint = minint
			}
		}
		c.options[name] = &sIntOption{
			option: opt,
			def:    defint,
			value:  defint,
			min:    minint,
			max:    maxint,
		}
	case "combo":
		// TODO
	case "button":
		// TODO
	default:
		return
	}
}

// Info
type Info struct {
	Line     string
	position *chess.Position
	err      error
}

func (i Info) Err() error { return i.err }

func (i Info) BestMove() (chess.Move, bool) {
	if move, ok := i.Value("bestmove"); ok {
		lan := chess.LongAlgebraicNotation{}
		m, err := lan.Decode(i.position, move) //i. board. ParseMove(move)
		if err != nil {
			m = &chess.Move{} // NullMove ?
		}
		return *m, true
	}
	return chess.Move{}, false
}

func ParseMove(pos *chess.Position, move string) (chess.Move, error) {
	lan := chess.LongAlgebraicNotation{}
	m, err := lan.Decode(pos, move)
	return *m, err
}

func (i Info) Pv() *Pv {
	pv, ok := i.Value("pv")
	if !ok {
		return nil
	}

	// score
	s, mate := i.Value("mate")
	if !mate {
		s, ok = i.Value("cp")
		if !ok {
			return nil
		}
	}
	score, _ := strconv.Atoi(s)
	//if i.board.SideToMove == chess.Black {
	if i.position.Turn() == chess.Black {
		score = -score
	}
	_, upper := i.Value("upperbound")
	_, lower := i.Value("lowerbound")

	// principal variation
	b := i.position
	fields := strings.Fields(pv)
	moves := make([]chess.Move, 0, len(fields))
	for _, move := range fields {
		m, err := ParseMove(b, move)
		if err != nil {
			break
		}
		moves = append(moves, m)
		//b = b.MakeMove(m)
		b = b.Update(&m)
	}
	multipv, ok := i.Value("multipv")
	if !ok {
		multipv = "0"
	}
	rank, _ := strconv.Atoi(multipv)
	rank--
	if rank < 0 {
		rank = 0
	}

	return &Pv{
		Moves:      moves,
		Score:      score,
		Mate:       mate,
		Upperbound: upper,
		Lowerbound: lower,
		Rank:       rank,
	}
}

func (i Info) Stats() *Stats {
	return &Stats{
		Depth:    i.intval("depth"),
		SelDepth: i.intval("seldepth"),
		Nodes:    i.intval("nodes"),
		Time:     time.Duration(i.intval("time")) * time.Millisecond,
	}
}

// Value returns the value of the given keyword. It returns !ok if the keyword
// is not present in this info.
func (i Info) Value(key string) (v string, ok bool) {
	return fieldValue(i.Line, key, infoKeywords)
}

func (i Info) intval(key string) int {
	v, ok := i.Value(key)
	if !ok {
		return 0
	}
	x, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return x
}

// Options

type option struct {
	name string
	cmdc chan<- interface{}
	errc <-chan error
}

func (o *option) send(cmd string) {
	o.cmdc <- cmd
	<-o.errc
}

type sStringOption struct {
	option
	def   string
	value string
}

func (s *sStringOption) StringDefault() string { return s.def }
func (s *sStringOption) String() string        { return s.value }
func (s *sStringOption) Set(value string) {
	s.value = value
	s.send(fmt.Sprintf("setoption name %s value %s", s.name, s.value))
}

type sIntOption struct {
	option
	def   int
	value int
	min   int
	max   int
}

func (i *sIntOption) StringDefault() string { return fmt.Sprint(i.def) }
func (i *sIntOption) String() string        { return fmt.Sprint(i.value) }
func (i *sIntOption) Default() int          { return i.def }
func (i *sIntOption) Int() int              { return i.value }
func (i *sIntOption) Min() int              { return i.min }
func (i *sIntOption) Max() int              { return i.max }

func (i *sIntOption) Set(value string) {
	v, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	i.SetInt(v)
}

func (i *sIntOption) SetInt(v int) {
	i.value = v
	i.send(fmt.Sprintf("setoption name %s value %d", i.name, i.value))
}

type sBoolOption struct {
	option
	def   bool
	value bool
}

func (b *sBoolOption) StringDefault() string { return fmt.Sprint(b.def) }
func (b *sBoolOption) String() string        { return fmt.Sprint(b.value) }
func (b *sBoolOption) Default() bool         { return b.def }
func (b *sBoolOption) Bool() bool            { return b.value }

func (b *sBoolOption) Set(value string) {
	v, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}
	b.SetBool(v)
}

func (b *sBoolOption) SetBool(v bool) {
	b.value = v
	b.send(fmt.Sprintf("setoption name %s value %v", b.name, b.value))
}

var _ StringOption = &sStringOption{}
var _ BoolOption = &sBoolOption{}
var _ IntOption = &sIntOption{}

// fields

var infoKeywords = map[string]bool{
	"bestmove":       true,
	"cp":             true,
	"cpuload":        true,
	"currline":       true,
	"currmove":       true,
	"currmovenumber": true,
	"depth":          true,
	"hashfull":       true,
	"info":           true,
	"lowerbound":     true,
	"mate":           true,
	"multipv":        true,
	"nodes":          true,
	"nps":            true,
	"ponder":         true,
	"pv":             true,
	"refutation":     true,
	"score":          true,
	"seldepth":       true,
	"string":         true,
	"tbhits":         true,
	"time":           true,
	"upperbound":     true,
}

var optionKeywords = map[string]bool{
	"name":    true,
	"type":    true,
	"default": true,
	"min":     true,
	"max":     true,
	"var":     true,
}

type fields struct {
	line string
	pos  int
}

func tokenise(line string) *fields {
	return &fields{line, 0}
}

func (f *fields) next() string {
	notIsSpace := func(r rune) bool { return !unicode.IsSpace(r) }

	l := f.line[f.pos:]
	i := strings.IndexFunc(l, notIsSpace)
	if i < 0 {
		return ""
	}
	j := i + strings.IndexFunc(l[i:], unicode.IsSpace)
	if j < i {
		j = len(l)
	}
	f.pos += j
	return l[i:j]
}

func (f *fields) hasNext() bool {
	return f.pos < len(f.line)
}

func (f *fields) remainder() string {
	return strings.TrimSpace(f.line[f.pos:])
}

func fieldValue(line, key string, keyword map[string]bool) (v string, ok bool) {
	field := &fields{line, 0}
	for field.next() != key {
		if !field.hasNext() {
			return "", false
		}
	}
	if key == "string" {
		// after the "string" keyword ignore other keywords
		return field.remainder(), true
	}
	p, q := field.pos, field.pos
	for field.hasNext() {
		f := field.next()
		if keyword[f] {
			break
		}
		q = field.pos
	}
	return strings.TrimSpace(line[p:q]), true
}
