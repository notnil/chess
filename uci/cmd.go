package uci

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/notnil/chess"
)

// Cmd is a UCI compliant command
type Cmd interface {
	fmt.Stringer
	ProcessResponse(e *Engine) error
}

type cmdNoOptions struct {
	Name string
	F    func(e *Engine) error
}

func (cmd cmdNoOptions) String() string {
	return cmd.Name
}

func (cmd cmdNoOptions) ProcessResponse(e *Engine) error {
	return cmd.F(e)
}

var (
	// CmdUCI corresponds to the "uci" command:
	// tell engine to use the uci (universal chess interface),
	// this will be send once as a first command after program boot
	// to tell the engine to switch to uci mode.
	// After receiving the uci command the engine must identify itself with the "id" command
	// and sent the "option" commands to tell the GUI which engine settings the engine supports if any.
	// After that the engine should sent "uciok" to acknowledge the uci mode.
	// If no uciok is sent within a certain time period, the engine task will be killed by the GUI.
	CmdUCI = cmdNoOptions{Name: "uci", F: func(e *Engine) error {
		e.id = map[string]string{}
		e.options = map[string]Option{}
		scanner := bufio.NewScanner(e.out)
		for scanner.Scan() {
			text := e.readLine(scanner)
			k, v, err := parseIDLine(text)
			if err == nil {
				e.id[k] = v
				continue
			}
			o := &Option{}
			err = o.UnmarshalText([]byte(text))
			if err == nil {
				e.options[o.Name] = *o
				continue
			}
			if text == "uciok" {
				break
			}
		}
		return nil
	}}

	// CmdIsReady corresponds to the "isready" command:
	// this is used to synchronize the engine with the GUI. When the GUI has sent a command or
	// multiple commands that can take some time to complete,
	// this command can be used to wait for the engine to be ready again or
	// to ping the engine to find out if it is still alive.
	// E.g. this should be sent after setting the path to the tablebases as this can take some time.
	// This command is also required once before the engine is asked to do any search
	// to wait for the engine to finish initializing.
	// This command must always be answered with "readyok" and can be sent also when the engine is calculating
	// in which case the engine should also immediately answer with "readyok" without stopping the search.
	CmdIsReady = cmdNoOptions{Name: "isready", F: func(e *Engine) error {
		scanner := bufio.NewScanner(e.out)
		for scanner.Scan() {
			text := e.readLine(scanner)
			if text == "readyok" {
				break
			}
		}
		return nil
	}}

	// CmdUCINewGame corresponds to the "ucinewgame" command:
	// this is sent to the engine when the next search (started with "position" and "go") will be from
	// a different game. This can be a new game the engine should play or a new game it should analyse but
	// also the next position from a testsuite with positions only.
	// If the GUI hasn't sent a "ucinewgame" before the first "position" command, the engine shouldn't
	// expect any further ucinewgame commands as the GUI is probably not supporting the ucinewgame command.
	// So the engine should not rely on this command even though all new GUIs should support it.
	// As the engine's reaction to "ucinewgame" can take some time the GUI should always send "isready"
	// after "ucinewgame" to wait for the engine to finish its operation.
	CmdUCINewGame = cmdNoOptions{Name: "ucinewgame", F: func(e *Engine) error {
		return nil
	}}

	// CmdPonderHit corresponds to the "ponderhit" command:
	// the user has played the expected move. This will be sent if the engine was told to ponder on the same move
	// the user has played. The engine should continue searching but switch from pondering to normal search.
	CmdPonderHit = cmdNoOptions{Name: "ponderhit", F: func(e *Engine) error {
		return nil
	}}

	// CmdStop corresponds to the "stop" command:
	// stop calculating as soon as possible,
	// don't forget the "bestmove" and possibly the "ponder" token when finishing the search
	CmdStop = cmdNoOptions{Name: "stop", F: func(e *Engine) error {
		return nil
	}}

	// CmdQuit (shouldn't be used directly as its handled by Engine.Close()) corresponds to the "quit" command:
	// quit the program as soon as possible
	CmdQuit = cmdNoOptions{Name: "quit", F: func(e *Engine) error {
		return nil
	}}
)

// CmdSetOption corresponds to the "setoption" command:
// this is sent to the engine when the user wants to change the internal parameters
// of the engine. For the "button" type no value is needed.
// One string will be sent for each parameter and this will only be sent when the engine is waiting.
// The name of the option in  should not be case sensitive and can inludes spaces like also the value.
// The substrings "value" and "name" should be avoided in  and  to allow unambiguous parsing,
// for example do not use  = "draw value".
// Here are some strings for the example below:
//    "setoption name Nullmove value true\n"
//   "setoption name Selectivity value 3\n"
//    "setoption name Style value Risky\n"
//    "setoption name Clear Hash\n"
//    "setoption name NalimovPath value c:\chess\tb\4;c:\chess\tb\5\n"
type CmdSetOption struct {
	Name  string
	Value string
}

func (cmd CmdSetOption) String() string {
	return fmt.Sprintf("setoption name %s value %s", cmd.Name, cmd.Value)
}

// ProcessResponse implements the Cmd interface
func (cmd CmdSetOption) ProcessResponse(e *Engine) error {
	return nil
}

// CmdPosition corresponds to the "position" command:
// set up the position described in fenstring on the internal board and
// play the moves on the internal chess board.
// if the game was played  from the start position the string "startpos" will be sent
// Note: no "new" command is needed. However, if this position is from a different game than
// the last position sent to the engine, the GUI should have sent a "ucinewgame" inbetween.
type CmdPosition struct {
	Position *chess.Position
	Moves    []*chess.Move
}

func (cmd CmdPosition) String() string {
	if cmd.Position == nil {
		cmd.Position = chess.StartingPosition()
	}
	if len(cmd.Moves) == 0 {
		return "position fen " + cmd.Position.String()
	}
	moveStrs := []string{}
	for _, m := range cmd.Moves {
		mStr := chess.UCINotation{}.Encode(nil, m)
		moveStrs = append(moveStrs, mStr)
	}
	return fmt.Sprintf("position fen %s moves %s", cmd.Position, strings.Join(moveStrs, " "))
}

// ProcessResponse implements the Cmd interface
func (CmdPosition) ProcessResponse(e *Engine) error {
	return nil
}

// CmdGo corresponds to the "go" command:
// start calculating on the current position set up with the "position" command.
// There are a number of commands that can follow this command, all will be sent in the same string.
// If one command is not send its value should be interpreted as it would not influence the search.
// * searchmoves  ....
// 	restrict search to this moves only
// 	Example: After "position startpos" and "go infinite searchmoves e2e4 d2d4"
// 	the engine should only search the two moves e2e4 and d2d4 in the initial position.
// * ponder
// 	start searching in pondering mode.
// 	Do not exit the search in ponder mode, even if it's mate!
// 	This means that the last move sent in in the position string is the ponder move.
// 	The engine can do what it wants to do, but after a "ponderhit" command
// 	it should execute the suggested move to ponder on. This means that the ponder move sent by
// 	the GUI can be interpreted as a recommendation about which move to ponder. However, if the
// 	engine decides to ponder on a different move, it should not display any mainlines as they are
// 	likely to be misinterpreted by the GUI because the GUI expects the engine to ponder
//    on the suggested move.
// * wtime
// 	white has x msec left on the clock
// * btime
// 	black has x msec left on the clock
// * winc
// 	white increment per move in mseconds if x > 0
// * binc
// 	black increment per move in mseconds if x > 0
// * movestogo
//   there are x moves to the next time control,
// 	this will only be sent if x > 0,
// 	if you don't get this and get the wtime and btime it's sudden death
// * depth
// 	search x plies only.
// * nodes
//    search x nodes only,
// * mate
// 	search for a mate in x moves
// * movetime
// 	search exactly x mseconds
// * infinite
// 	search until the "stop" command. Do not exit the search without being told so in this mode!
type CmdGo struct {
	SearchMoves    []*chess.Move
	Ponder         bool
	WhiteTime      time.Duration
	BlackTime      time.Duration
	WhiteIncrement time.Duration
	BlackIncrement time.Duration
	MovesToGo      int
	Depth          int
	Nodes          int
	Mate           int
	MoveTime       time.Duration
	Infinite       bool
}

func (cmd CmdGo) String() string {
	a := []string{"go"}
	if cmd.Ponder {
		a = append(a, "ponder")
	}
	if cmd.WhiteTime > 0 {
		a = append(a, "wtime", msecStr(cmd.WhiteTime))
	}
	if cmd.BlackTime > 0 {
		a = append(a, "btime", msecStr(cmd.BlackTime))
	}
	if cmd.WhiteIncrement > 0 {
		a = append(a, "winc", msecStr(cmd.WhiteIncrement))
	}
	if cmd.BlackIncrement > 0 {
		a = append(a, "binc", msecStr(cmd.BlackIncrement))
	}
	if cmd.MovesToGo > 0 {
		a = append(a, "movestogo", fmt.Sprint(cmd.MovesToGo))
	}
	if cmd.Depth > 0 {
		a = append(a, "depth", fmt.Sprint(cmd.Depth))
	}
	if cmd.Nodes > 0 {
		a = append(a, "nodes", fmt.Sprint(cmd.Nodes))
	}
	if cmd.Mate > 0 {
		a = append(a, "mate", fmt.Sprint(cmd.Nodes))
	}
	if cmd.MoveTime > 0 {
		a = append(a, "movetime", msecStr(cmd.MoveTime))
	}
	if cmd.Infinite {
		a = append(a, "infinite")
	}
	if len(cmd.SearchMoves) > 0 {
		a = append(a, "searchmoves")
		for _, m := range cmd.SearchMoves {
			mStr := chess.UCINotation{}.Encode(nil, m)
			a = append(a, mStr)
		}
	}
	return strings.Join(a, " ")
}

// ProcessResponse implements the Cmd interface
func (CmdGo) ProcessResponse(e *Engine) error {
	scanner := bufio.NewScanner(e.out)
	results, err := ProcessEngineOutput(scanner, e.getDebugLogger())
	if err != nil {
		return err
	}
	e.results = *results
	return nil
}

func ProcessEngineOutput(scanner *bufio.Scanner, debugLogger *log.Logger) (*SearchResults, error) {
	results := SearchResults{}
	for scanner.Scan() {
		text := scanner.Text()
		if debugLogger != nil {
			debugLogger.Println(text)
		}
		if strings.HasPrefix(text, "bestmove") {
			parts := strings.Split(text, " ")
			if len(parts) <= 1 {
				return nil, errors.New("best move not found " + text)
			}
			bestMove, err := chess.UCINotation{}.Decode(nil, parts[1])
			if err != nil {
				return nil, err
			}
			results.BestMove = bestMove
			if len(parts) >= 4 {
				ponderMove, err := chess.UCINotation{}.Decode(nil, parts[3])
				if err != nil {
					return nil, err
				}
				results.Ponder = ponderMove
			}
			break
		}

		var info Info
		err := info.UnmarshalText([]byte(text))
		if err != nil {
			continue
		}
		switch info.Multipv {
		case 1:
			// We've received the first PV line, so we can clear the multipvInfo
			results.MultiPV = []Info{}
			results.Info = info
		case 0:
			results.Info = info
		}
		if info.Multipv > 0 {
			results.MultiPV = append(results.MultiPV, info)
		}
	}
	return &results, nil
}

func parseIDLine(s string) (string, string, error) {
	if strings.HasPrefix(s, "id") == false {
		return "", "", errors.New("uci: invalid id line")
	}
	parts := strings.Split(s, " ")
	if len(parts) < 3 {
		return "", "", errors.New("uci: invalid id line")
	}
	return parts[1], strings.Join(parts[2:], " "), nil
}

func msecStr(dur time.Duration) string {
	return fmt.Sprint(int(dur / time.Millisecond))
}
