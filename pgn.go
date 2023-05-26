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
type ScannerOpts struct {
	ExpandVariations bool
}
type Scanner struct {
	scanr *bufio.Scanner
	games []*Game
	err   error
	opts  ScannerOpts
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
	if len(s.games) > 0 {
		return true
	}
	var sb strings.Builder
	state := notInPGN
	setGames := func() bool {
		games, err := decodePGNs(sb.String(), s.opts.ExpandVariations)
		if err != nil {
			s.err = err
			return false
		}
		s.games = games
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
			return setGames()
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
				return setGames()
			}
			sb.WriteString(line + "\n")
		}
	}
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
	gameList, err := decodePGNs(pgn, false)
	if err != nil {
		return nil, err
	}
	if len(gameList) != 1 {
		return nil, fmt.Errorf("chess: pgn decode error unexpected game count %v", len(gameList))
	}

	return gameList[0], nil
}

func decodePGNs(pgn string, expandVariations bool) ([]*Game, error) {
	ret := []*Game{}
	tagPairs := getTagPairs(pgn)
	moveListSet, err := moveListSetWithComments(pgn, expandVariations)
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

	for idx, ml := range moveListSet.moveLists {
		g := NewGame(gameFuncs...)
		g.ignoreAutomaticDraws = true
		decoder := multiDecoder([]Decoder{AlgebraicNotation{}, LongAlgebraicNotation{}, UCINotation{}})
		for _, move := range ml.moves {
			m, err := decoder.Decode(g.Position(), move.MoveStr)
			if err != nil {
				return nil, fmt.Errorf("chess: pgn decode error %s on variation	%d move %d", err.Error(), idx, g.Position().moveCount)
			}
			if err := g.Move(m); err != nil {
				return nil, fmt.Errorf("chess: pgn invalid move error %s on	variation %d move %d", err.Error(), idx, g.Position().moveCount)
			}
			g.comments = append(g.comments, move.Comments)
		}
		g.outcome = ml.outcome

		ret = append(ret, g)
	}

	return ret, nil
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

type moveListAndOutcome struct {
	moves   []moveWithComment
	outcome Outcome
}

type moveListSet struct {
	moveLists []moveListAndOutcome
}

var moveListTokenRe = regexp.MustCompile(`(?:\d+\.)|(O-O(?:-O)?|\w*[abcdefgh][12345678]\w*(?:=[QRBN])?(?:\+|#)?)|(?:\{([^}]*)\})|(?:\([^)]*\))|(\*|0-1|1-0|1\/2-1\/2)`)

func moveListSetWithComments(pgn string, expandVariations bool) (moveListSet, error) {
	ret := moveListSet{
		moveLists: []moveListAndOutcome{},
	}

	if !expandVariations {
		ml, err := moveListWithCommentsNoExpand(pgn)
		if err != nil {
			return ret, err
		}
		ret.moveLists = append(ret.moveLists, ml)
		return ret, nil
	}

	return moveListSetExpanded(pgn)
}

func moveListWithCommentsNoExpand(pgn string) (moveListAndOutcome, error) {
	pgn = stripTagPairs(pgn)
	ret := moveListAndOutcome{
		moves: []moveWithComment{},
	}
	// moveListTokenRe doesn't work w/ nested variations
	pgn, err := stripVariations(pgn)
	if err != nil {
		return ret, err
	}

	for _, match := range moveListTokenRe.FindAllStringSubmatch(pgn, -1) {
		move, commentText, outcomeText := match[1], match[2], match[3]
		if len(move+commentText+outcomeText) == 0 {
			continue
		}

		if outcomeText != "" {
			ret.outcome = Outcome(outcomeText)
			break
		}

		if commentText != "" {
			ret.moves[len(ret.moves)-1].Comments = append(ret.moves[len(ret.moves)-1].Comments, strings.TrimSpace(commentText))
		}

		if move != "" {
			ret.moves = append(ret.moves, moveWithComment{MoveStr: move})
		}
	}
	return ret, nil
}

var moveNumRe = regexp.MustCompile(`(?:\d+\.+)?(.*)`)

func moveListSetExpanded(pgn string) (moveListSet, error) {
	firstGame := moveListAndOutcome{
		moves: []moveWithComment{},
	}
	ret := moveListSet{
		moveLists: []moveListAndOutcome{firstGame},
	}

	pgn = stripTagPairs(pgn)
	// remove comments @todo need to add comments back in
	pgn = removeSection("{", "}", pgn)
	// remove line breaks
	pgn = strings.Replace(pgn, "\n", " ", -1)
	pgn = strings.ReplaceAll(pgn, "(", "( ")
	pgn = strings.ReplaceAll(pgn, ")", " )")

	moveListIdx := 0
	moveListIdxStack := make([]int, 0)
	list := strings.Split(pgn, " ")

	for _, move := range list {
		move = strings.TrimSpace(move)
		switch move {
		case string(NoOutcome), string(WhiteWon), string(BlackWon), string(Draw):
			ret.moveLists[moveListIdx].outcome = Outcome(move)
		case "":
		case "(":
			// begin new variation
			moveListIdxStack = append(moveListIdxStack, moveListIdx)
			newIdx := len(ret.moveLists)
			numMoves := len(ret.moveLists[moveListIdx].moves) - 1
			newGame := moveListAndOutcome{}
			newGame.moves = make([]moveWithComment, numMoves)
			copy(newGame.moves, ret.moveLists[moveListIdx].moves)
			ret.moveLists = append(ret.moveLists, newGame)
			moveListIdx = newIdx

		case ")":
			// end current variation
			stackSize := len(moveListIdxStack)
			if stackSize == 0 {
				return ret, fmt.Errorf("Failed to parse variation")
			}
			moveListIdx = moveListIdxStack[stackSize-1]
			moveListIdxStack = moveListIdxStack[:stackSize-1]
		default:
			results := moveNumRe.FindStringSubmatch(move)
			tmp := moveWithComment{}
			if len(results) == 2 && results[1] != "" {
				tmp.MoveStr = results[1]
				ret.moveLists[moveListIdx].moves = append(ret.moveLists[moveListIdx].moves, tmp)
			}
		}
	}

	return ret, nil
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

func stripVariations(pgn string) (string, error) {
	var ret strings.Builder

	variationDepth := 0
	inCommentSection := false

	for _, c := range pgn {
		if c == '{' {
			if inCommentSection {
				return "", fmt.Errorf("chess: pgn decode mismatched { in variation: %v", pgn)
			}
			inCommentSection = true
		} else if c == '}' {
			if !inCommentSection {
				return "", fmt.Errorf("chess: pgn decode mismatched } in variation: %v", pgn)
			}
			inCommentSection = false
		}
		if !inCommentSection && c == '(' {
			variationDepth++
			continue
		}
		if !inCommentSection && c == ')' {
			if variationDepth <= 0 {
				return "", fmt.Errorf("chess: pgn decode mismatched parenthesis in variation: %v", pgn)
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
