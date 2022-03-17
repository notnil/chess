package chess

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// Scanner is modeled on the bufio.Scanner type but
// instead of reading lines, it reads chess games
// from concatenated PGN files.  It is designed to
// replace GamesFromPGN in order to handle very large
// PGN database files such as https://database.lichess.org/.
type Scanner struct {
	scanr *bufio.Scanner
	game  *Game
	err   error
}

// NewScanner returns a new scanner.
func NewScanner(r io.Reader) *Scanner {
	scanr := bufio.NewScanner(r)
	return &Scanner{scanr: scanr}
}

type scanState int

const (
	notInPGN scanState = iota
	inTagPairs
	inMoves
)

// Scan returns false if there was an error parsing
// a game or EOF was reached.  Running scan populates
// data for Next() and Err().
func (s *Scanner) Scan() bool {
	if s.err == io.EOF {
		return false
	}
	s.err = nil
	var sb strings.Builder
	state := notInPGN
	setGame := func() bool {
		game, err := decodePGN(sb.String())
		if err != nil {
			s.err = err
			return false
		}
		s.game = game
		return true
	}
	for {
		scan := s.scanr.Scan()
		if !scan {
			s.err = s.scanr.Err()
			// err is nil if io.EOF
			if s.err == nil {
				s.err = io.EOF
			}
			return setGame()
		}
		line := strings.TrimSpace(s.scanr.Text())
		isTagPair := strings.HasPrefix(line, "[")
		isMoveSeq := strings.HasPrefix(line, "1. ")
		switch state {
		case notInPGN:
			if !isTagPair {
				break
			}
			state = inTagPairs
			sb.WriteString(line + "\n")
		case inTagPairs:
			if isMoveSeq {
				state = inMoves
			}
			sb.WriteString(line + "\n")
		case inMoves:
			if line == "" {
				return setGame()
			}
			sb.WriteString(line + "\n")
		}
	}
}

// Next returns the game from the most recent Scan.
func (s *Scanner) Next() *Game {
	return s.game
}

// Err returns an error encountered during scanning.
// Typically this will be a PGN parsing error or an
// io.EOF.
func (s *Scanner) Err() error {
	return s.err
}

// GamesFromPGN returns all PGN decoding games from the
// reader.  It is designed to be used decoding multiple PGNs
// in the same file.  An error is returned if there is an
// issue parsing the PGNs.
// Deprecated: Use Scanner instead.
func GamesFromPGN(r io.Reader) ([]*Game, error) {
	games := []*Game{}
	current := ""
	count := 0
	totalCount := 0
	br := bufio.NewReader(r)
	eof := false
	for !eof {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			eof = true
		} else if err != nil {
			return nil, err
		}
		if strings.TrimSpace(line) == "" {
			count++
		} else {
			current += line
		}
		if count == 2 || eof {
			game, err := decodePGN(current)
			if err != nil {
				return nil, err
			}
			games = append(games, game)
			count = 0
			current = ""
			totalCount++
		}
	}
	return games, nil
}

type multiDecoder []Decoder

func (a multiDecoder) Decode(pos *Position, s string) (*Move, error) {
	for _, d := range a {
		m, err := d.Decode(pos, s)
		if err == nil {
			return m, nil
		}
	}
	return nil, fmt.Errorf(`chess: failed to decode notation text "%s" for position %s`, s, pos)
}

func decodePGN(pgn string) (*Game, error) {
	tagPairs := getTagPairs(pgn)
	moveComments, outcome := moveListWithComments(pgn)
	gameFuncs := []func(*Game){}
	for _, tp := range tagPairs {
		if strings.ToLower(tp.Key) == "fen" {
			fenFunc, err := FEN(tp.Value)
			if err != nil {
				return nil, fmt.Errorf("chess: pgn decode error %s on tag %s", err.Error(), tp.Key)
			}
			gameFuncs = append(gameFuncs, fenFunc)
			break
		}
	}
	gameFuncs = append(gameFuncs, TagPairs(tagPairs))
	g := NewGame(gameFuncs...)
	g.ignoreAutomaticDraws = true
	decoder := multiDecoder([]Decoder{AlgebraicNotation{}, LongAlgebraicNotation{}, UCINotation{}})
	for _, move := range moveComments {
		m, err := decoder.Decode(g.Position(), move.MoveStr)
		if err != nil {
			return nil, fmt.Errorf("chess: pgn decode error %s on move %d", err.Error(), g.Position().moveCount)
		}
		if err := g.Move(m); err != nil {
			return nil, fmt.Errorf("chess: pgn invalid move error %s on move %d", err.Error(), g.Position().moveCount)
		}
		g.comments = append(g.comments, move.Comments)
	}
	g.outcome = outcome
	return g, nil
}

func encodePGN(g *Game) string {
	s := ""
	for _, tag := range g.tagPairs {
		s += fmt.Sprintf("[%s \"%s\"]\n", tag.Key, tag.Value)
	}
	s += "\n"
	for i, move := range g.moves {
		pos := g.positions[i]
		txt := g.notation.Encode(pos, move)
		if i%2 == 0 {
			s += fmt.Sprintf("%d. %s", (i/2)+1, txt)
		} else {
			s += fmt.Sprintf(" %s ", txt)
		}
		if len(g.comments) > i {
			for _, c := range g.comments[i] {
				s += " { " + c + " } "
			}
		}
	}
	s += " " + string(g.outcome)
	return s
}

var (
	tagPairRegex = regexp.MustCompile(`\[(.*)\s\"(.*)\"\]`)
)

func getTagPairs(pgn string) []*TagPair {
	tagPairs := []*TagPair{}
	matches := tagPairRegex.FindAllString(pgn, -1)
	for _, m := range matches {
		results := tagPairRegex.FindStringSubmatch(m)
		if len(results) == 3 {
			pair := &TagPair{
				Key:   results[1],
				Value: results[2],
			}
			tagPairs = append(tagPairs, pair)
		}
	}
	return tagPairs
}

type moveWithComment struct {
	MoveStr  string
	Comments []string
}

var moveListTokenRe = regexp.MustCompile(`(?:\d+\.)|(O-O(?:-O)?|\w*[abcdefgh][12345678]\w*(?:=[QRBN])?(?:\+|#)?)|(?:\{([^}]*)\})|(?:\([^)]*\))|(\*|0-1|1-0|1\/2-1\/2)`)

func moveListWithComments(pgn string) ([]moveWithComment, Outcome) {
	pgn = stripTagPairs(pgn)
	var outcome Outcome
	moves := []moveWithComment{}

	for _, match := range moveListTokenRe.FindAllStringSubmatch(pgn, -1) {
		move, commentText, outcomeText := match[1], match[2], match[3]
		if len(move+commentText+outcomeText) == 0 {
			continue
		}

		if outcomeText != "" {
			outcome = Outcome(outcomeText)
			break
		}

		if commentText != "" {
			moves[len(moves)-1].Comments = append(moves[len(moves)-1].Comments, strings.TrimSpace(commentText))
		}

		if move != "" {
			moves = append(moves, moveWithComment{MoveStr: move})
		}
	}
	return moves, outcome
}

func stripTagPairs(pgn string) string {
	lines := strings.Split(pgn, "\n")
	cp := []string{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "[") {
			cp = append(cp, line)
		}
	}
	return strings.Join(cp, "\n")
}
