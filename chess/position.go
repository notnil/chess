package chess

import (
	"fmt"
	"strings"
)

// Side represents a side of the board.
type Side int

const (
	// KingSide is the right side of the board from white's perspective.
	KingSide Side = iota + 1
	// QueenSide is the left side of the board from white's perspective.
	QueenSide
)

// CastleRights holds the state of both sides castling abilities.
type CastleRights string

// CanCastle returns true if the given color and side combination
// can castle, otherwise returns false.
func (cr CastleRights) CanCastle(c Color, side Side) bool {
	char := "k"
	if side == QueenSide {
		char = "q"
	}
	if c == White {
		char = strings.ToUpper(char)
	}
	return strings.Contains(string(cr), char)
}

// String implements the fmt.Stringer interface and returns
// a FEN compatible string.  Ex. KQq
func (cr CastleRights) String() string {
	return string(cr)
}

// Position represents the state of the game without regaurd
// to its outcome.  Position is translatable to FEN notation.
type Position struct {
	board           *Board
	turn            Color
	castleRights    CastleRights
	enPassantSquare Square
	halfMoveClock   int
	moveCount       int
}

// Board returns the position's board.
func (pos *Position) Board() *Board {
	return pos.board
}

// Turn returns the color to move next.
func (pos *Position) Turn() Color {
	return pos.turn
}

// CastleRights returns the castling rights of the position.
func (pos *Position) CastleRights() CastleRights {
	return pos.castleRights
}

// String implements the fmt.Stringer interface and returns a
// string with the format: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
func (pos *Position) String() string {
	b := pos.board.String()
	t := pos.turn.String()
	c := pos.castleRights.String()
	sq := "-"
	if pos.enPassantSquare != NoSquare {
		sq = pos.enPassantSquare.String()
	}
	return fmt.Sprintf("%s %s %s %s %d %d", b, t, c, sq, pos.halfMoveClock, pos.moveCount)
}

// MarshalText implements the encoding.TextMarshaler interface and
// encodes the position's FEN.
func (pos *Position) MarshalText() (text []byte, err error) {
	return []byte(pos.String()), nil
}

// UnmarshalText implements the encoding.TextUnarshaler interface and
// assumes the data is in the FEN format.
func (pos *Position) UnmarshalText(text []byte) error {
	cp, err := decodeFEN(string(text))
	if err != nil {
		return err
	}
	pos.board = cp.board
	pos.turn = cp.turn
	pos.castleRights = cp.castleRights
	pos.enPassantSquare = cp.enPassantSquare
	pos.halfMoveClock = cp.halfMoveClock
	pos.moveCount = cp.moveCount
	return nil
}

func (pos *Position) ValidMoves() []*Move {
	moves := []*Move{}
	moves = append(moves, pos.pawnMoves()...)
	moves = append(moves, pos.knightMoves()...)
	return moves
}

type validMoveBB struct {
	bb    bitboard
	shift int
}

func (pos *Position) pawnMoves() []*Move {
	var bbs []*validMoveBB
	var bbEnPassant bitboard
	if pos.enPassantSquare != NoSquare {
		bbEnPassant = newBitboard(map[Square]bool{pos.enPassantSquare: true})
	}
	var p *Piece
	if pos.Turn() == White {
		p = WhitePawn
		bbWhite := pos.board.bbs[WhitePawn]
		bbWhiteCapRight := ((bbWhite & ^bbFileH & ^bbRank8) >> 9) & (pos.board.blackSqs | bbEnPassant)
		bbWhiteCapLeft := ((bbWhite & ^bbFileA & ^bbRank8) >> 7) & (pos.board.blackSqs | bbEnPassant)
		bbWhiteUpOne := ((bbWhite & ^bbRank8) >> 8) & pos.board.emptySqs
		bbWhiteUpTwo := ((bbWhiteUpOne & bbRank3) >> 8) & pos.board.emptySqs
		bbs = []*validMoveBB{
			{bb: bbWhiteCapRight, shift: 9},
			{bb: bbWhiteCapLeft, shift: 7},
			{bb: bbWhiteUpOne, shift: 8},
			{bb: bbWhiteUpTwo, shift: 16},
		}
	} else {
		p = BlackPawn
		bbBlack := pos.board.bbs[BlackPawn]
		bbBlackCapRight := ((bbBlack & ^bbFileH & ^bbRank1) << 7) & (pos.board.whiteSqs | bbEnPassant)
		bbBlackCapLeft := ((bbBlack & ^bbFileA & ^bbRank1) << 9) & (pos.board.whiteSqs | bbEnPassant)
		bbBlackUpOne := ((bbBlack & ^bbRank1) << 8) & pos.board.emptySqs
		bbBlackUpTwo := ((bbBlackUpOne & bbRank6) << 8) & pos.board.emptySqs
		bbs = []*validMoveBB{
			{bb: bbBlackCapRight, shift: -7},
			{bb: bbBlackCapLeft, shift: -9},
			{bb: bbBlackUpOne, shift: -8},
			{bb: bbBlackUpTwo, shift: -16},
		}
	}
	return movesFromValidBBs(p, bbs)
}

func (pos *Position) knightMoves() []*Move {
	p := WhiteKnight
	validBB := ^pos.board.whiteSqs
	if pos.Turn() == Black {
		p = BlackKnight
		validBB = ^pos.board.blackSqs
	}
	bb := pos.board.bbs[p]
	bbUpRight := ((bb & ^(bbRank7 | bbRank8) & ^bbFileH) >> 17) & validBB
	bbUpLeft := ((bb & ^(bbRank7 | bbRank8) & ^bbFileA) >> 15) & validBB
	bbRightUp := (bb & ^bbRank8 & ^(bbFileG | bbFileH) >> 10) & validBB
	bbRightDown := (bb & ^bbRank1 & ^(bbFileG | bbFileH) << 10) & validBB
	bbDownRight := ((bb & ^(bbRank1 | bbRank2) & ^bbFileH) << 15) & validBB
	bbDownLeft := ((bb & ^(bbRank1 | bbRank2) & ^bbFileA) << 17) & validBB
	bbLeftUp := (bb & ^bbRank8 & ^(bbFileA | bbFileB) >> 6) & validBB
	bbLeftDown := (bb & ^bbRank1 & ^(bbFileA | bbFileB) << 6) & validBB
	bbs := []*validMoveBB{
		{bb: bbUpRight, shift: 17},
		{bb: bbUpLeft, shift: 15},
		{bb: bbRightUp, shift: 10},
		{bb: bbRightDown, shift: -10},
		{bb: bbDownRight, shift: -15},
		{bb: bbDownLeft, shift: -17},
		{bb: bbLeftUp, shift: 6},
		{bb: bbLeftDown, shift: -6},
	}
	return movesFromValidBBs(p, bbs)
}

func movesFromValidBBs(p *Piece, bbs []*validMoveBB) []*Move {
	moves := []*Move{}
	for _, validBB := range bbs {
		if validBB.bb == 0 {
			continue
		}
		// TODO reduce number of squares by range of LSB to MSB per bitboard
		for sq := 0; sq < 64; sq++ {
			if validBB.bb.Occupied(Square(sq)) {
				s1 := Square(sq - validBB.shift)
				s2 := Square(sq)
				if (p == WhitePawn && s2.rank() == rank8) || (p == BlackPawn && s2.rank() == rank1) {
					qm := &Move{s1: s1, s2: s2, promo: Queen}
					rm := &Move{s1: s1, s2: s2, promo: Rook}
					bm := &Move{s1: s1, s2: s2, promo: Bishop}
					nm := &Move{s1: s1, s2: s2, promo: Knight}
					moves = append(moves, qm, rm, bm, nm)
				} else {
					moves = append(moves, &Move{s1: s1, s2: s2})
				}
			}
		}
	}
	return moves
}
