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
	// cp, err := decodeFEN(string(text))
	// if err != nil {
	// 	return err
	// }
	// pos.board = cp.board
	// pos.turn = cp.turn
	// pos.castleRights = cp.castleRights
	// pos.enPassantSquare = cp.enPassantSquare
	// pos.halfMoveClock = cp.halfMoveClock
	// pos.moveCount = cp.moveCount
	return nil
}
