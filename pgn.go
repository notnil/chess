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
		isMoveSeq := strings.HasPrefix(line, "1. ") || strings.HasPrefix(line, "*")
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
	// moveStrs, outcome := moveList(pgn)
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
			s += fmt.Sprintf("%d.%s", (i/2)+1, txt)
		} else {
			s += fmt.Sprintf(" %s ", txt)
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

var (
	moveNumRegex = regexp.MustCompile(`(?:\d+\.+)?(.*)`)
)

type moveWithComment struct {
	MoveStr  string
	Comments []string
}

func moveListWithComments(pgn string) ([]moveWithComment, Outcome) {
	text := stripTagPairs(pgn)
	// remove variations
	text = removeSection(`\(`, `\)`, text)
	text = strings.Replace(text, "\n", " ", -1)
	text = strings.TrimSpace(text)
	tokens := strings.Split(text, " ")
	var outcome Outcome
	moves := []moveWithComment{}
	inComment := false
	commentTokens := []string{}
tokenLoop:
	for _, token := range tokens {
		token = strings.TrimSpace(token)
		switch token {
		case "{":
			inComment = true
			commentTokens = []string{}
		case "}":
			inComment = false
			if len(moves) > 0 {
				moves[len(moves)-1].Comments = append(moves[len(moves)-1].Comments, strings.Join(commentTokens, " "))
			}
		case "":
		case string(NoOutcome), string(WhiteWon), string(BlackWon), string(Draw):
			outcome = Outcome(token)
			break tokenLoop
		default:
			if inComment {
				commentTokens = append(commentTokens, token)
				break
			}
			if strings.HasSuffix(token, ".") {
				break
			}
			moves = append(moves, moveWithComment{MoveStr: token})
		}
	}
	return moves, outcome
}

func moveList(pgn string) ([]string, Outcome) {
	// remove comments
	text := removeSection("{", "}", pgn)
	// remove variations
	text = removeSection(`\(`, `\)`, text)
	// remove tag pairs
	text = removeSection(`\[`, `\]`, text)
	// remove line breaks
	text = strings.Replace(text, "\n", " ", -1)

	list := strings.Split(text, " ")
	filtered := []string{}
	var outcome Outcome
	for _, move := range list {
		move = strings.TrimSpace(move)
		switch move {
		case string(NoOutcome), string(WhiteWon), string(BlackWon), string(Draw):
			outcome = Outcome(move)
		case "":
		default:
			results := moveNumRegex.FindStringSubmatch(move)
			if len(results) == 2 && results[1] != "" {
				filtered = append(filtered, results[1])
			}
		}
	}
	return filtered, outcome
}

func removeSection(leftChar, rightChar, s string) string {
	r := regexp.MustCompile(leftChar + ".*?" + rightChar)
	for {
		i := r.FindStringIndex(s)
		if i == nil {
			return s
		}
		s = s[0:i[0]] + s[i[1]:]
	}
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
