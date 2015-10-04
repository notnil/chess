package chess

import "fmt"

type CastleRights struct {
	WhiteKingSide  bool
	WhiteQueenSide bool
	BlackKingSide  bool
	BlackQueenSide bool
}

func (c CastleRights) String() string {
	s := ""
	if c.WhiteKingSide {
		s += "K"
	}
	if c.WhiteQueenSide {
		s += "Q"
	}
	if c.BlackKingSide {
		s += "k"
	}
	if c.BlackQueenSide {
		s += "q"
	}
	if len(s) == 0 {
		s = "-"
	}
	return s
}

type GameState struct {
	Board           Board
	Turn            Color
	CastleRights    CastleRights
	EnPassantSquare *Square
	HalfMoveClock   int
	MoveCount       int
}

// String implements the fmt.Stringer interface and returns a
// string with the format: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
func (gs *GameState) String() string {
	b := gs.Board.String()
	t := gs.Turn.String()
	c := gs.CastleRights.String()
	sq := "-"
	if gs.EnPassantSquare != nil {
		sq = gs.EnPassantSquare.String()
	}
	return fmt.Sprintf("%s %s %s %s %d %d", b, t, c, sq, gs.HalfMoveClock, gs.MoveCount)
}

//
// func (b Board) inCheckmate(c color) bool {
// 	// should only happen in unit tests
// 	if b.kingSquare(c) == nil {
// 		return false
// 	}
// 	return b.inCheck(c) && !b.hasValidMoves(c)
// }
