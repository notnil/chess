package chess

import (
	"fmt"
	"strings"
)

type GameState struct {
	board           Board
	turn            Color
	castleRights    CastleRights
	enPassantSquare *Square
	halfMoveClock   int
	moveCount       int
}

// String implements the fmt.Stringer interface and returns a
// string with the format: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
func (gs *GameState) String() string {
	b := gs.board.String()
	t := gs.turn.String()
	c := gs.castleRights.String()
	sq := "-"
	if gs.enPassantSquare != nil {
		sq = gs.enPassantSquare.String()
	}
	return fmt.Sprintf("%s %s %s %s %d %d", b, t, c, sq, gs.halfMoveClock, gs.moveCount)
}

type Side int

const (
	KingSide Side = iota + 1
	QueenSide
)

type CastleRights string

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

func (cr CastleRights) String() string {
	return string(cr)
}

//
// func (b Board) inCheckmate(c color) bool {
// 	// should only happen in unit tests
// 	if b.kingSquare(c) == nil {
// 		return false
// 	}
// 	return b.inCheck(c) && !b.hasValidMoves(c)
// }
