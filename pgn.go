package chess

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Notation interface {
	Encode(*Game) []byte
	Decode(text []byte) (*Game, error)
}

type PGN struct{}

func (n PGN) Encode(g *Game) []byte {
	return []byte{}
}

func (n PGN) Decode(b []byte) (*Game, error) {
	text := string(b)
	// remove comments
	text = removeSection("{", "}", text)
	// remove tag pairs
	text = removeSection(`\[`, `\]`, text)
	// remove line breaks
	text = strings.Replace(text, "\n", "", -1)
	g := NewGame()
	for _, txtMove := range sepMoves(text) {
		// log.Println(g.board)
		// log.Println(txtMove)
		m, err := formMove(g, txtMove)
		if err != nil {
			return nil, err
		} else if m == nil {
			continue
		}
		log.Println("MOVE", m)
		if err := g.Move(m.s1, m.s2, m.promo); err != nil {
			return nil, err
		}
	}
	return g, nil
}

const (
	pgnShortCastle = "O-O"
	pgnLongCastle  = "O-O-O"
	pgnWhiteWin    = "1-0"
	pgnBlackWin    = "0-1"
	pgnDraw        = "1/2-1/2"
)

func formMove(g *Game, txtMove string) (*move, error) {
	// e4 Bb5 Ra6+ Rae8 Nxf7 Rxe1+ O-O O-O-O b8=Q+
	c := g.turn
	switch txtMove {
	case pgnShortCastle:
		s1 := square(fileE, c.backRank())
		s2 := square(fileG, c.backRank())
		return &move{s1: s1, s2: s2}, nil
	case pgnLongCastle:
		s1 := square(fileE, c.backRank())
		s2 := square(fileC, c.backRank())
		return &move{s1: s1, s2: s2}, nil
	case pgnWhiteWin:
		g.status = WhiteWon
		return nil, nil
	case pgnBlackWin:
		g.status = BlackWon
		return nil, nil
	case pgnDraw:
		g.status = Draw
		return nil, nil
	}
	parts, err := calcMoveParts(txtMove)
	if err != nil {
		return nil, err
	}
	s1 := parts.S1(g)
	if s1 == nil {
		return nil, errors.New("chess: pgn decoder couldn't resolve origin square for " + txtMove)
	}
	promo := parts.Promotion(g)
	return &move{s1: s1, s2: parts.s2, promo: promo}, nil
}

type moveParts struct {
	p      pieceType
	s1File *file
	s1Rank *rank
	s2     *Square
	promo  *pieceType
}

const (
	moveRegex = `([K,Q,R,B,N])?([a-h])?([1-8])?x?([a-h][1-8])(?:e\.p\.)?\+?\+?(=[Q,R,B,N])?`
)

func calcMoveParts(txtMove string) (*moveParts, error) {
	// parsed list format: ["Ra1xe8+" "R" "a" "1" "e8" ""]
	// parseErr := errors.New("chess: pgn move text " + txtMove + " could not be parsed")
	r := regexp.MustCompile(moveRegex)
	m := r.FindStringSubmatch(txtMove)
	if m == nil || len(m) != 6 {
		return nil, errors.New("chess: pgn move text " + txtMove + " could not be parsed")
	}
	s2 := squareFromString(m[4])
	if s2 == nil {
		return nil, errors.New("chess: pgn decoder couldn't resolve destination square for " + txtMove)
	}
	p, found := peiceTypeFromChar(m[1])
	if !found {
		return nil, errors.New("chess: pgn decoder couldn't resolve piece type for " + txtMove)
	}
	var promo *pieceType
	promoType, found := peiceTypeFromChar(m[5])
	if found {
		promo = &promoType
	}
	return &moveParts{
		p:      p,
		s1File: fileFromString(m[2]),
		s1Rank: rankFromString(m[3]),
		s2:     s2,
		promo:  promo,
	}, nil
}

func (m *moveParts) S1(g *Game) *Square {
	for _, sq := range g.board.squaresForColor(g.turn) {
		p := g.board.piece(sq)
		if p == nil || p.pieceType() != m.p {
			continue
		}
		if m.s1File != nil && sq.file != *m.s1File {
			continue
		}
		if m.s1Rank != nil && sq.rank != *m.s1Rank {
			continue
		}
		promo := m.Promotion(g)
		if err := g.copy().Move(sq, m.s2, promo); err == nil {
			return sq
		}
	}
	return nil
}

func (m *moveParts) Promotion(g *Game) *Piece {
	var promo *Piece
	if m.promo != nil && *m.promo != pawn {
		promo = piece(g.turn, *m.promo)
	}
	return promo
}

func peiceTypeFromChar(c string) (pieceType, bool) {
	switch c {
	case "K":
		return king, true
	case "Q":
		return queen, true
	case "R":
		return rook, true
	case "B":
		return bishop, true
	case "N":
		return knight, true
	case "":
		return pawn, true
	}
	return pawn, false
}

func sepMoves(n string) []string {
	moves := []string{}
	addMovePair := func(s string) {
		s = strings.TrimSpace(s)
		for _, m := range strings.Split(s, " ") {
			m = strings.TrimSpace(m)
			if m != "" {
				moves = append(moves, m)
			}
		}
	}

	count := 1
	for {
		start := fmt.Sprintf("%d.", count)
		startIndex := strings.Index(n, start)
		count += 1
		end := fmt.Sprintf("%d.", count)
		endIndex := strings.Index(n, end)
		if startIndex == -1 {
			return moves
		} else if endIndex == -1 {
			addMovePair(n[startIndex+len(start) : len(n)])
		} else {
			addMovePair(n[startIndex+len(start) : endIndex])
		}
	}
	return moves
}

func removeSection(leftChar, rightChar, s string) string {
	r := regexp.MustCompile(leftChar + ".*" + rightChar)
	for {
		i := r.FindStringIndex(s)
		if i == nil {
			return s
		}
		s = s[0:i[0]] + s[i[1]:len(s)]
	}
	return s
}
