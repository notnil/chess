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
	return moves
}

const (
	bitShiftUpLeft    = 7
	bitShiftUpRight   = 9
	bitShiftUp        = 8
	bitShiftDownLeft  = 9
	bitShiftDownRight = 7
	bitShiftDown      = 8
)

type pawnBB struct {
	bb    bitboard
	shift int
}

func (pos *Position) pawnMoves() []*Move {
	moves := []*Move{}
	bbs := []pawnBB{}
	var bbEnPassant bitboard
	if pos.enPassantSquare != NoSquare {
		bbEnPassant = newBitboard(map[Square]bool{pos.enPassantSquare: true})
	}
	if pos.Turn() == White {
		bbWhite := pos.board.bbs[WhitePawn]
		bbWhiteCapRight := ((bbWhite & ^bbFileH & ^bbRank8) >> bitShiftUpRight) & (pos.board.blackSqs | bbEnPassant)
		bbWhiteCapLeft := ((bbWhite & ^bbFileA & ^bbRank8) >> bitShiftUpLeft) & (pos.board.blackSqs | bbEnPassant)
		bbWhiteUpOne := ((bbWhite & ^bbRank8) >> bitShiftUp) & pos.board.emptySqs
		bbWhiteUpTwo := ((bbWhiteUpOne & bbRank3) >> bitShiftUp) & pos.board.emptySqs
		bbs = append(bbs, pawnBB{bb: bbWhiteCapLeft, shift: bitShiftUpLeft})
		bbs = append(bbs, pawnBB{bb: bbWhiteCapRight, shift: bitShiftUpRight})
		bbs = append(bbs, pawnBB{bb: bbWhiteUpOne, shift: bitShiftUp})
		bbs = append(bbs, pawnBB{bb: bbWhiteUpTwo, shift: bitShiftUp * 2})
	} else {
		bbBlack := pos.board.bbs[BlackPawn]
		bbBlackCapRight := ((bbBlack & ^bbFileH & ^bbRank1) << bitShiftDownRight) & (pos.board.whiteSqs | bbEnPassant)
		bbBlackCapLeft := ((bbBlack & ^bbFileA & ^bbRank1) << bitShiftDownLeft) & (pos.board.whiteSqs | bbEnPassant)
		bbBlackUpOne := ((bbBlack & ^bbRank1) << bitShiftUp) & pos.board.emptySqs
		bbBlackUpTwo := ((bbBlackUpOne & bbRank6) << bitShiftUp) & pos.board.emptySqs
		bbs = append(bbs, pawnBB{bb: bbBlackCapRight, shift: bitShiftDownRight * -1})
		bbs = append(bbs, pawnBB{bb: bbBlackCapLeft, shift: bitShiftDownLeft * -1})
		bbs = append(bbs, pawnBB{bb: bbBlackUpOne, shift: bitShiftDown * -1})
		bbs = append(bbs, pawnBB{bb: bbBlackUpTwo, shift: bitShiftDown * -2})
	}

	for _, pawnBB := range bbs {
		if pawnBB.bb == 0 {
			continue
		}
		// TODO reduce number of squares by range of LSB to MSB per bitboard
		for sq := 0; sq < 64; sq++ {
			if pawnBB.bb.Occupied(Square(sq)) {
				s1 := Square(sq - pawnBB.shift)
				s2 := Square(sq)
				if s2.rank() == rank8 {
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
