package chess

import (
	"log"
	"testing"
)

var (
	validMoves = []*Move{
		// pawn moves
		&Move{s1: E2, s2: E4, state: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		&Move{s1: A2, s2: A3, state: unsafeFEN("8/8/8/8/8/8/P7/8 w - - 0 1")},
		&Move{s1: A2, s2: A4, state: unsafeFEN("8/8/8/8/8/8/P7/8 w - - 0 1")},
		&Move{s1: A7, s2: A6, state: unsafeFEN("8/p7/8/8/8/8/8/8 b - - 0 1")},
		&Move{s1: A7, s2: A5, state: unsafeFEN("8/p7/8/8/8/8/8/8 b - - 0 1")},
		&Move{s1: C4, s2: B5, state: unsafeFEN("8/8/8/1p1p4/2P5/8/8/8 w - - 0 1")},
		&Move{s1: C4, s2: D5, state: unsafeFEN("8/8/8/1p1p4/2P5/8/8/8 w - - 0 1")},
		&Move{s1: C4, s2: C5, state: unsafeFEN("8/8/8/1p1p4/2P5/8/8/8 w - - 0 1")},
		&Move{s1: C5, s2: B4, state: unsafeFEN("8/8/8/2p5/1P1P4/8/8/8 b - - 0 1")},
		&Move{s1: C5, s2: D4, state: unsafeFEN("8/8/8/2p5/1P1P4/8/8/8 b - - 0 1")},
		&Move{s1: C5, s2: C4, state: unsafeFEN("8/8/8/2p5/1P1P4/8/8/8 b - - 0 1")},
	}

	invalidMoves = []*Move{
		// out of turn moves
		&Move{s1: E7, s2: E5, state: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		&Move{s1: E2, s2: E4, state: unsafeFEN("rnbqkbnr/1ppppppp/p7/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1")},
		// pawn moves
		&Move{s1: E2, s2: D3, state: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		&Move{s1: E2, s2: F3, state: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		&Move{s1: E2, s2: E5, state: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
	}
)

func unsafeFEN(s string) *GameState {
	g, err := FEN(s)
	if err != nil {
		log.Fatal(err)
	}
	return g
}

func TestValidMoves(t *testing.T) {
	for _, m := range validMoves {
		if !m.isValid() {
			log.Println(m.state.Board.Draw())
			t.Fatalf("expected move %s to be valid", m)
		}
	}
}

// func TestInvalidMoves(t *testing.T) {
// 	for _, m := range invalidMoves {
// 		if m.isValid() {
// 			t.Fatalf("expected move %s to be invalid", m)
// 		}
// 	}
// }
