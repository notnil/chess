package chess

import (
	"fmt"
	"strings"
)

func encodeMove(pos *Position, move *Move) string {
	checkChar := getCheckChar(pos, move)
	if move.HasTag(KingSideCastle) {
		return "O-O" + checkChar
	} else if move.HasTag(QueenSideCastle) {
		return "O-O-O" + checkChar
	}
	p := pos.Board().piece(move.S1())
	pChar := charFromPieceType(p.Type())
	s1Str := formS1(pos, move)
	capChar := ""
	if move.HasTag(Capture) || move.HasTag(EnPassant) {
		capChar = "x"
		if p.Type() == Pawn && s1Str == "" {
			capChar = move.s1.File().String() + "x"
		}
	}
	epText := ""
	if move.HasTag(EnPassant) {
		epText = "e.p."
	}
	promoText := charForPromo(move.promo)
	return fmt.Sprint(pChar, s1Str, capChar, move.s2, epText, promoText, checkChar)
}

func decodeMove(pos *Position, s string) (*Move, error) {
	s = strings.Replace(s, "?", "", -1)
	s = strings.Replace(s, "!", "", -1)
	s = strings.Replace(s, "+", "", -1)
	s = strings.Replace(s, "#", "", -1)
	s = strings.Replace(s, "e.p.", "", -1)
	moves := pos.ValidMoves()
	for _, move := range moves {
		str := encodeMove(pos, move)
		str = strings.Replace(str, "+", "", -1)
		str = strings.Replace(str, "#", "", -1)
		str = strings.Replace(str, "e.p.", "", -1)
		if str == s {
			return move, nil
		}
	}
	return nil, fmt.Errorf("chess: could not decode algebraic notation %s for position %s %s", s, pos.String(), moves)
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
