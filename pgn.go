package chess

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
)

// Scanner is modeled on the bufio.Scanner type but
// instead of reading lines, it reads chess games
// from concatenated PGN files.  It is designed to
// replace GamesFromPGN in order to handle very large
// PGN database files such as https://database.lichess.org/.
type ScannerOpts struct {
	ExpandVariations bool
}
type Scanner struct {
	scanr *bufio.Scanner
	games []*Game
	err   error
	opts  ScannerOpts
}

type moveList struct {
	moveStrs []string
	outcome  Outcome
}

// NewScanner returns a new scanner with default options
func NewScanner(r io.Reader) *Scanner {
	defaultOpts := ScannerOpts{ExpandVariations: false}

	return NewScannerWithOptions(r, defaultOpts)
}

// NewScanner returns a new scanner with explicit options
func NewScannerWithOptions(r io.Reader, o ScannerOpts) *Scanner {
	scanr := bufio.NewScanner(r)
	g := make([]*Game, 0)
	return &Scanner{scanr: scanr, opts: o, games: g}
}

// Scan returns false if there was an error parsing
// a game or EOF was reached.  Running scan populates
// data for Next() and Err().
func (s *Scanner) Scan() bool {
	s.err = nil

	if len(s.games) > 0 {
		return true
	}

	var sb strings.Builder
	var parsedOne bool

	for {
		scan := s.scanr.Scan()
		if scan == false {
			s.err = s.scanr.Err()
			return false
		}
		line := s.scanr.Text() + "\n"
		if strings.TrimSpace(line) == "" {
			continue
		} else {
			sb.WriteString(line)
		}

		parsedOne = strings.HasPrefix(line, "1.")
		if parsedOne {
			parsedOne = false
			games, err := decodePGN(sb.String(), s.opts.ExpandVariations)
			if err != nil {
				s.err = err
				return false
			}
			s.games = games
			break
		}
	}
	return true
}

// Next returns the game from the most recent Scan.
func (s *Scanner) Next() *Game {
	if len(s.games) == 0 {
		return nil
	}

	g := s.games[0]
	s.games = s.games[1:]

	return g
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
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if strings.TrimSpace(line) == "" {
			count++
		} else {
			current += line
		}
		if count == 2 {
			newGames, err := decodePGN(current, false)
			if err != nil {
				return nil, err
			}
			games = append(games, newGames[0])
			count = 0
			current = ""
			totalCount++
			log.Println("Processed game", totalCount)
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

func decodePGN(pgn string, expandVariations bool) ([]*Game, error) {
	tagPairs := getTagPairs(pgn)
	moveLists, err := getMoveLists(pgn, expandVariations)
	if err != nil {
		return nil, err
	}
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
	games := make([]*Game, 0)
	varName := "main line"

	for idx, moveList := range moveLists {
		if idx > 0 {
			varName = fmt.Sprintf("variation %v", idx)
		}
		g := NewGame(gameFuncs...)
		g.ignoreAutomaticDraws = true
		decoder := multiDecoder([]Decoder{AlgebraicNotation{}, LongAlgebraicNotation{}, UCINotation{}})
		for _, alg := range moveList.moveStrs {
			m, err := decoder.Decode(g.Position(), alg)
			if err != nil {
				return nil, fmt.Errorf("chess: pgn decode error %s on move %d of %v", err.Error(), g.Position().moveCount, varName)
			}
			if err := g.Move(m); err != nil {
				return nil, fmt.Errorf("chess: pgn invalid move error %s on move %d of %v", err.Error(), g.Position().moveCount, varName)
			}
		}
		g.outcome = moveList.outcome
		games = append(games, g)
	}

	return games, nil
}

func encodePGN(g *Game) string {
	s := ""
	for _, tag := range g.tagPairs {
		s += fmt.Sprintf("[%s \"%s\"]\n", tag.Key, tag.Value)
	}
	if s != "" {
		s += "\n"
	}
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

func getMoveLists(pgn string, expandVariations bool) ([]moveList, error) {
	// remove comments
	text := removeSection("{", "}", pgn)
	// remove tag pairs
	text = removeSection(`\[`, `\]`, text)
	// remove line breaks
	text = strings.Replace(text, "\n", " ", -1)

	var err error
	if !expandVariations {
		text, err = removeVariations(text)
		if err != nil {
			return nil, err
		}
	} else {
		text = strings.ReplaceAll(text, "(", "( ")
		text = strings.ReplaceAll(text, ")", " )")
	}

	list := strings.Split(text, " ")

	moveListIdx := 0
	moveListIdxStack := make([]int, 0)
	moveLists := make([]moveList, 0)
	moveLists = append(moveLists, moveList{moveStrs: make([]string, 0)})

	for _, move := range list {
		move = strings.TrimSpace(move)
		switch move {
		case string(NoOutcome), string(WhiteWon), string(BlackWon), string(Draw):
			moveLists[moveListIdx].outcome = Outcome(move)
		case "":
		case "(":
			// begin new variation
			moveListIdxStack = append(moveListIdxStack, moveListIdx)
			newIdx := len(moveLists)
			numMoves := len(moveLists[moveListIdx].moveStrs) - 1
			moveLists = append(moveLists, moveList{moveStrs: make([]string, numMoves)})
			copy(moveLists[newIdx].moveStrs, moveLists[moveListIdx].moveStrs)
			moveListIdx = newIdx

		case ")":
			// end current variation
			stackSize := len(moveListIdxStack)
			if stackSize == 0 {
				return nil, fmt.Errorf("Failed to parse variation")
			}
			moveListIdx = moveListIdxStack[stackSize-1]
			moveListIdxStack = moveListIdxStack[:stackSize-1]
		default:
			results := moveNumRegex.FindStringSubmatch(move)
			if len(results) == 2 && results[1] != "" {
				moveLists[moveListIdx].moveStrs = append(moveLists[moveListIdx].moveStrs, results[1])
			}
		}
	}
	return moveLists, nil
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

func removeVariations(s string) (string, error) {
	var ret strings.Builder
	variationDepth := 0

	// cannot just use regexp since variations may be nested
	for _, c := range s {
		if c == '(' {
			variationDepth++
			continue
		}
		if c == ')' {
			if variationDepth <= 0 {
				return "", fmt.Errorf("Mismatched parenthesis in variation: %v", s)
			}
			variationDepth--
			continue
		}
		if variationDepth == 0 {
			_, err := ret.WriteRune(c)
			if err != nil {
				return "", err
			}
		}
	}

	return ret.String(), nil
}
