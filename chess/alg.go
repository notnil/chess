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
	if move.HasTag(Capture) {
		capChar = "x"
		if p.Type() == Pawn && s1Str == "" {
			capChar = move.s1.file().String() + "x"
		}
	}
	epText := ""
	if move.HasTag(EnPassant) {
		epText = "e.p."
	}
	promoText := charFromPromo(move.promo)
	return fmt.Sprint(pChar, s1Str, capChar, move.s2, epText, promoText, checkChar)
}

func decodeMove(pos *Position, s string) (*Move, error) {
	s = strings.Replace(s, "?", "", -1)
	moves := pos.ValidMoves()
	for _, move := range moves {
		str := encodeMove(pos, move)
		if str == s {
			return move, nil
		}
	}
	return nil, fmt.Errorf("chess: could not decode algebraic notation %s for position %s", s, pos.String())
}

func getCheckChar(pos *Position, move *Move) string {
	if !move.HasTag(Check) {
		return ""
	}
	nextPos := pos.Update(move)
	status := nextPos.status()
	if status == Checkmate {
		return "#"
	}
	return "+"
}

func formS1(pos *Position, m *Move) string {
	moves := pos.ValidMoves()
	// find moves for piece type
	pMoves := []*Move{}
	files := map[file]int{}
	ranks := map[rank]int{}
	p := pos.board.piece(m.s1)
	for _, mv := range moves {
		if mv.s2 == m.s2 && p == pos.board.piece(mv.s1) {
			pMoves = append(pMoves, mv)
			files[mv.s1.file()] = files[mv.s1.file()] + 1
			ranks[mv.s1.rank()] = ranks[mv.s1.rank()] + 1
		}
	}
	if len(pMoves) == 1 {
		return ""
	} else if len(files) == len(pMoves) {
		return m.s1.file().String()
	} else if len(ranks) == len(pMoves) {
		return m.s1.rank().String()
	}
	return m.s1.String()
}

var (
	algPieceMap = map[string]PieceType{
		"K": King,
		"Q": Queen,
		"R": Rook,
		"B": Bishop,
		"N": Knight,
		"":  Pawn,
	}
	algPromoMap = map[string]PieceType{
		"Q": Queen,
		"R": Rook,
		"B": Bishop,
		"N": Knight,
	}
)

func peiceTypeFromChar(c string) PieceType {
	return algPieceMap[c]
}

func promoFromChar(c string) PieceType {
	return algPromoMap[c]
}

func charFromPieceType(p PieceType) string {
	for c, pt := range algPieceMap {
		if pt == p {
			return c
		}
	}
	return ""
}

func charFromPromo(p PieceType) string {
	for c, pt := range algPromoMap {
		if pt == p {
			return c
		}
	}
	return ""
}
