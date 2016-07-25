package chess

import (
	"fmt"
	"strings"
)

type Encoder interface {
	Encode(pos *Position, m *Move) string
}

type Decoder interface {
	Decode(pos *Position, s string) (*Move, error)
}

type Notation interface {
	Encoder
	Decoder
	Name() string
}

/*
Move format:
------------
The move format is in long algebraic notation.
A nullmove from the Engine to the GUI should be sent as 0000.
Examples:  e2e4, e7e5, e1g1 (white short castling), e7e8q (for promotion)
*/
type LongAlgebraicNotation struct{}

func (_ LongAlgebraicNotation) Name() string {
	return "Long Algebraic Notation"
}

func (_ LongAlgebraicNotation) Encode(pos *Position, m *Move) string {
	return m.S1().String() + m.S2().String() + m.Promo().String()
}

func (_ LongAlgebraicNotation) Decode(pos *Position, s string) (*Move, error) {
	l := len(s)
	err := fmt.Errorf(`chess: failed to decode long algebraic notation text "%s" for position %s`, s, pos.String())
	if l < 4 || l > 5 {
		return nil, err
	}
	s1, ok := strToSquareMap[s[0:2]]
	if !ok {
		return nil, err
	}
	s2, ok := strToSquareMap[s[2:4]]
	if !ok {
		return nil, err
	}
	promo := NoPieceType
	if l == 5 {
		promo = pieceTypeFromChar(s[4:5])
		if promo == NoPieceType {
			return nil, err
		}
	}
	m := &Move{s1: s1, s2: s2, promo: promo}
	p := pos.Board().piece(s1)
	if p.Type() == King {
		if (s1 == E1 && s2 == G1) || (s1 == E8 && s2 == G8) {
			m.addTag(KingSideCastle)
		} else if (s1 == E1 && s2 == C1) || (s1 == E8 && s2 == C8) {
			m.addTag(QueenSideCastle)
		}
	} else if p.Type() == Pawn && s2 == pos.enPassantSquare {
		m.addTag(EnPassant)
		m.addTag(Capture)
	}
	c1 := p.Color()
	c2 := pos.Board().piece(s2).Color()
	if c2 != NoColor && c1 != c2 {
		m.addTag(Capture)
	}
	return m, nil
}

type AlgebraicNotation struct{}

func (_ AlgebraicNotation) Name() string {
	return "Algebraic Notation"
}

func (_ AlgebraicNotation) Encode(pos *Position, m *Move) string {
	checkChar := getCheckChar(pos, m)
	if m.HasTag(KingSideCastle) {
		return "O-O" + checkChar
	} else if m.HasTag(QueenSideCastle) {
		return "O-O-O" + checkChar
	}
	p := pos.Board().piece(m.S1())
	pChar := charFromPieceType(p.Type())
	s1Str := formS1(pos, m)
	capChar := ""
	if m.HasTag(Capture) || m.HasTag(EnPassant) {
		capChar = "x"
		if p.Type() == Pawn && s1Str == "" {
			capChar = m.s1.File().String() + "x"
		}
	}
	epText := ""
	if m.HasTag(EnPassant) {
		epText = "e.p."
	}
	promoText := charForPromo(m.promo)
	return pChar + s1Str + capChar + m.s2.String() + epText + promoText + checkChar
}

func (_ AlgebraicNotation) Decode(pos *Position, s string) (*Move, error) {
	s = removeSubstrings(s, "?", "!", "+", "#", "e.p.")
	for _, m := range pos.ValidMoves() {
		str := AlgebraicNotation{}.Encode(pos, m)
		str = removeSubstrings(str, "?", "!", "+", "#", "e.p.")
		if str == s {
			return m, nil
		}
	}
	return nil, fmt.Errorf("chess: could not decode algebraic notation %s for position %s", s, pos.String())
}

func getCheckChar(pos *Position, move *Move) string {
	if !move.HasTag(Check) {
		return ""
	}
	nextPos := pos.Update(move)
	if nextPos.Status() == Checkmate {
		return "#"
	}
	return "+"
}

func formS1(pos *Position, m *Move) string {
	moves := pos.ValidMoves()
	// find moves for piece type
	pMoves := []*Move{}
	files := map[File]int{}
	ranks := map[Rank]int{}
	p := pos.board.piece(m.s1)
	if p.Type() == Pawn {
		return ""
	}
	for _, mv := range moves {
		if mv.s2 == m.s2 && p == pos.board.piece(mv.s1) {
			pMoves = append(pMoves, mv)
			files[mv.s1.File()] = files[mv.s1.File()] + 1
			ranks[mv.s1.Rank()] = ranks[mv.s1.Rank()] + 1
		}
	}
	if len(pMoves) == 1 {
		return ""
	} else if len(files) == len(pMoves) {
		return m.s1.File().String()
	} else if len(ranks) == len(pMoves) {
		return m.s1.Rank().String()
	}
	return m.s1.String()
}

func charForPromo(p PieceType) string {
	c := charFromPieceType(p)
	if c != "" {
		c = "=" + c
	}
	return c
}

func charFromPieceType(p PieceType) string {
	switch p {
	case King:
		return "K"
	case Queen:
		return "Q"
	case Rook:
		return "R"
	case Bishop:
		return "B"
	case Knight:
		return "N"
	}
	return ""
}

func moveFuncForPieceType(p PieceType) moveFunc {
	switch p {
	case King:
		return kingMoves
	case Queen:
		return queenMoves
	case Rook:
		return rookMoves
	case Bishop:
		return bishopMoves
	case Knight:
		return knightMoves
	}
	return pawnMoves
}

func pieceTypeFromChar(c string) PieceType {
	switch c {
	case "q":
		return Queen
	case "r":
		return Rook
	case "b":
		return Bishop
	case "n":
		return Knight
	}
	return NoPieceType
}

func removeSubstrings(s string, subs ...string) string {
	for _, sub := range subs {
		s = strings.Replace(s, sub, "", -1)
	}
	return s
}
