package chess_test

import (
	"reflect"
	"testing"

	. "github.com/loganjspears/chess"
)

var (
	validFENs = map[string]*GameState{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1": &GameState{
			Board: map[*Square]*Piece{
				A1: WhiteRook, B1: WhiteKnight, C1: WhiteBishop, D1: WhiteQueen, E1: WhiteKing, F1: WhiteBishop, G1: WhiteKnight, H1: WhiteRook, // white back row
				A2: WhitePawn, B2: WhitePawn, C2: WhitePawn, D2: WhitePawn, E2: WhitePawn, F2: WhitePawn, G2: WhitePawn, H2: WhitePawn, // white front row
				A7: BlackPawn, B7: BlackPawn, C7: BlackPawn, D7: BlackPawn, E7: BlackPawn, F7: BlackPawn, G7: BlackPawn, H7: BlackPawn, // black front row
				A8: BlackRook, B8: BlackKnight, C8: BlackBishop, D8: BlackQueen, E8: BlackKing, F8: BlackBishop, G8: BlackKnight, H8: BlackRook, // black back row
			},
			Turn:            White,
			CastleRights:    &CastleRights{true, true, true, true},
			EnPassantSquare: nil,
			HalfMoveClock:   0,
			MoveCount:       1,
		},
	}
)

func TestValidFENS(t *testing.T) {
	for fen, expectedState := range validFENs {
		state, err := FEN(fen)
		if err != nil {
			t.Fatal("recieved unexpected error", err)
		}
		if !reflect.DeepEqual(expectedState, state) {
			t.Fatalf("fen expected game state %+v but got %+v", expectedState, state)
		}
	}
}

func TestFEN(t *testing.T) {
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	state, err := FEN(fen)
	if err != nil {
		t.Fatal("recieved unexpected error", err)
	}
	if fen != state.String() {
		t.Fatalf("fen expected board string %s but got %s", fen, state.String())
	}
}
