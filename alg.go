package chess

import (
	"fmt"
	"regexp"
)

const (
	algKingSideCastle  = "O-O"
	algQueenSideCastle = "O-O-O"
)

var (
	// parsed list format: ["Ra1xe8+" "R" "a" "1" "e8" ""]
	moveReg     = regexp.MustCompile(`([KQRBN])?([a-h])?([1-8])?x?([a-h][1-8])(?:e\.p\.)?\+?\+?=?([QRBN])?`)
	ksCastleReg = regexp.MustCompile(`^O-O[\+#\?!]*$`)
	qsCastleReg = regexp.MustCompile(`^O-O-O[\+#\?!]*$`)
)

const (
	algPieceIdx = iota + 1
	algFileIdx
	algRankIdx
	algS2Idx
	algPromoIdx
)

func encodeMove(move *Move) string {
	checkChar := getCheckChar(move)
	if move.isCastling() {
		switch move.s2.file {
		case G:
			return algKingSideCastle + checkChar
		}
		return algQueenSideCastle + checkChar
	}
	pChar := charFromPieceType(move.piece().Type())
	s1Str := formS1(move)
	capChar := ""
	if move.isCapture() {
		capChar = "x"
		if move.piece().Type() == Pawn && s1Str == "" {
			capChar = move.s1.file.String() + "x"
		}
	}

	epText := ""
	if pawnEnPassantFilter(move) {
		epText = "e.p."
	}
	promoText := charFromPromo(move.promo)
	return fmt.Sprint(pChar, s1Str, capChar, move.s2, epText, promoText, checkChar)
}

func getCheckChar(move *Move) string {
	checkChar := ""
	postState := move.postMoveState()
	if _, method := postState.getOutcome(); method == Checkmate {
		checkChar = "#"
	} else if postState.board.inCheck(move.state.turn.Other()) {
		checkChar = "+"
	}
	return checkChar
}

func decodeMove(state *GameState, s string) (*Move, error) {
	ksCastle := ksCastleReg.MatchString(s)
	qsCastle := qsCastleReg.MatchString(s)
	if ksCastle || qsCastle {
		return formCastle(state, s), nil
	}
	match := moveReg.FindStringSubmatch(s)
	if match == nil || len(match) != 6 {
		return nil, fmt.Errorf("chess: algebraic notation move %s could not be parsed", s)
	}
	s2 := squareFromStr(match[algS2Idx])
	if s2 == nil {
		return nil, fmt.Errorf("chess: algebraic notation move %s could not resolve destination square", s)
	}
	pieceType := peiceTypeFromChar(match[algPieceIdx])
	if pieceType == NoPiece {
		return nil, fmt.Errorf("chess: algebraic notation move %s could not resolve piece type", s)
	}
	promo := promoFromChar(match[algPromoIdx])
	file := fileFromStr(match[algFileIdx])
	rank := rankFromStr(match[algRankIdx])
	move := findMove(state, s2, pieceType, file, rank, promo)

	if move == nil {
		return nil, fmt.Errorf("chess: algebraic notation move %s could not resolve origin square", s)
	}
	return move, nil
}

func findMove(state *GameState, s2 *Square, pieceType PieceType, file File, rank Rank, promo PieceType) *Move {
	moves := []*Move{}
	for _, sq := range state.board.squaresForColor(state.turn) {
		p := state.board.piece(sq)
		if p == nil || p.Type() != pieceType {
			continue
		}
		if file != NoFile && sq.file != file {
			continue
		}
		if rank != NoRank && sq.rank != rank {
			continue
		}
		move := &Move{s1: sq, s2: s2, promo: promo, state: state}
		if move.isValid() {
			moves = append(moves, move)
		}
	}
	if len(moves) == 1 {
		return moves[0]
	}
	return nil
}

func formS1(m *Move) string {
	possibleSqs := []*Square{}
	for _, sq := range m.state.board.squaresForColor(m.state.turn) {
		if m.piece().Type() != m.state.board.piece(sq).Type() {
			continue
		}
		move := &Move{s1: sq, s2: m.s2, promo: m.promo, state: m.state}
		if move.isValid() {
			possibleSqs = append(possibleSqs, sq)
		}
	}
	if len(possibleSqs) == 1 {
		return ""
	}
	files := map[File]int{}
	ranks := map[Rank]int{}
	for _, sq := range possibleSqs {
		files[sq.file] = files[sq.file] + 1
		ranks[sq.rank] = ranks[sq.rank] + 1
	}
	if len(files) == len(possibleSqs) {
		return m.s1.file.String()
	} else if len(ranks) == len(possibleSqs) {
		return m.s1.rank.String()
	}
	return m.s1.String()
}

func formCastle(state *GameState, s string) *Move {
	backRow := [8]*Square{A1, B1, C1, D1, E1, F1, G1, H1}
	if state.turn == Black {
		backRow = [8]*Square{A8, B8, C8, D8, E8, F8, G8, H8}
	}
	ksCastle := ksCastleReg.MatchString(s)
	qsCastle := qsCastleReg.MatchString(s)
	if qsCastle {
		return &Move{
			s1:    backRow[4],
			s2:    backRow[2],
			state: state,
		}
	} else if ksCastle {
		return &Move{
			s1:    backRow[4],
			s2:    backRow[6],
			state: state,
		}
	}
	return nil
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
