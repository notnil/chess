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
		&Move{s1: A4, s2: B3, state: unsafeFEN("2r3k1/1q1nbppp/r3p3/3pP3/pPpP4/P1Q2N2/2RN1PPP/2R4K b - b3 0 23")},
		&Move{s1: A2, s2: A1, promo: Queen.Ptr(), state: unsafeFEN("8/8/8/8/8/8/p7/8 b - - 0 1")},
		// knight moves
		&Move{s1: E4, s2: F6, state: unsafeFEN("8/8/8/3pp3/4N3/8/5B2/8 w - - 0 1")},
		&Move{s1: E4, s2: D6, state: unsafeFEN("8/8/8/3pp3/4N3/8/5B2/8 w - - 0 1")},
		&Move{s1: E4, s2: C3, state: unsafeFEN("8/8/8/3pp3/4N3/8/5B2/8 w - - 0 1")},
		// bishop moves
		&Move{s1: E4, s2: H7, state: unsafeFEN("8/8/8/3pp3/4B3/5N2/8/8 w - - 0 1")},
		&Move{s1: E4, s2: D5, state: unsafeFEN("8/8/8/3pp3/4B3/5N2/8/8 w - - 0 1")},
		&Move{s1: E4, s2: B1, state: unsafeFEN("8/8/8/3pp3/4B3/5N2/8/8 w - - 0 1")},
		// rook moves
		&Move{s1: B2, s2: B4, state: unsafeFEN("8/1p5b/4N3/4p3/8/8/1R6/1B6 w - - 0 1")},
		&Move{s1: B2, s2: B7, state: unsafeFEN("8/1p5b/4N3/4p3/8/8/1R6/1B6 w - - 0 1")},
		&Move{s1: B2, s2: A2, state: unsafeFEN("8/1p5b/4N3/4p3/8/8/1R6/1B6 w - - 0 1")},
		&Move{s1: B2, s2: H2, state: unsafeFEN("8/1p5b/4N3/4p3/8/8/1R6/1B6 w - - 0 1")},
		// queen moves
		&Move{s1: B2, s2: E5, state: unsafeFEN("8/1p5b/4N3/4p3/8/8/1Q6/1B6 w - - 0 1")},
		&Move{s1: B2, s2: A1, state: unsafeFEN("8/1p5b/4N3/4p3/8/8/1Q6/1B6 w - - 0 1")},
		&Move{s1: B2, s2: A2, state: unsafeFEN("8/1p5b/4N3/4p3/8/8/1Q6/1B6 w - - 0 1")},
		&Move{s1: B2, s2: H2, state: unsafeFEN("8/1p5b/4N3/4p3/8/8/1Q6/1B6 w - - 0 1")},
	}

	invalidMoves = []*Move{
		// out of turn moves
		&Move{s1: E7, s2: E5, state: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		&Move{s1: E2, s2: E4, state: unsafeFEN("rnbqkbnr/1ppppppp/p7/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1")},
		// pawn moves
		&Move{s1: E2, s2: D3, state: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		&Move{s1: E2, s2: F3, state: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		&Move{s1: E2, s2: E5, state: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		&Move{s1: A2, s2: A1, promo: nil, state: unsafeFEN("8/8/8/8/8/8/p7/8 b - - 0 1")},
		// knight moves
		&Move{s1: E4, s2: F2, state: unsafeFEN("8/8/8/3pp3/4N3/8/5B2/8 w - - 0 1")},
		&Move{s1: E4, s2: F3, state: unsafeFEN("8/8/8/3pp3/4N3/8/5B2/8 w - - 0 1")},
		// bishop moves
		&Move{s1: E4, s2: C6, state: unsafeFEN("8/8/8/3pp3/4B3/5N2/8/8 w - - 0 1")},
		&Move{s1: E4, s2: E5, state: unsafeFEN("8/8/8/3pp3/4B3/5N2/8/8 w - - 0 1")},
		&Move{s1: E4, s2: E4, state: unsafeFEN("8/8/8/3pp3/4B3/5N2/8/8 w - - 0 1")},
		&Move{s1: E4, s2: F3, state: unsafeFEN("8/8/8/3pp3/4B3/5N2/8/8 w - - 0 1")},
		// rook moves
		&Move{s1: B2, s2: B1, state: unsafeFEN("8/1p5b/4N3/4p3/8/8/1R6/1B6 w - - 0 1")},
		&Move{s1: B2, s2: C3, state: unsafeFEN("8/1p5b/4N3/4p3/8/8/1R6/1B6 w - - 0 1")},
		&Move{s1: B2, s2: B8, state: unsafeFEN("8/1p5b/4N3/4p3/8/8/1R6/1B6 w - - 0 1")},
		&Move{s1: B2, s2: G7, state: unsafeFEN("8/1p5b/4N3/4p3/8/8/1R6/1B6 w - - 0 1")},
		// queen moves
		&Move{s1: B2, s2: B1, state: unsafeFEN("8/1p5b/4N3/4p3/8/8/1Q6/1B6 w - - 0 1")},
		&Move{s1: B2, s2: C4, state: unsafeFEN("8/1p5b/4N3/4p3/8/8/1Q6/1B6 w - - 0 1")},
		&Move{s1: B2, s2: B8, state: unsafeFEN("8/1p5b/4N3/4p3/8/8/1Q6/1B6 w - - 0 1")},
		&Move{s1: B2, s2: G7, state: unsafeFEN("8/1p5b/4N3/4p3/8/8/1Q6/1B6 w - - 0 1")},
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

func TestInvalidMoves(t *testing.T) {
	for _, m := range invalidMoves {
		if m.isValid() {
			log.Println(m.state.Board.Draw())
			t.Fatalf("expected move %s to be invalid", m)
		}
	}
}
