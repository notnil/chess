package chess

import (
	"log"
	"testing"
)

type moveState struct {
	Move      *Move
	PostState *GameState
}

var (
	validMoves = []*Move{
		// pawn moves
		&Move{s1: E2, s2: E4, state: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		&Move{s1: A2, s2: A3, state: unsafeFEN("8/8/8/8/8/8/P7/8 w - - 0 1")},
		&Move{s1: A7, s2: A6, state: unsafeFEN("8/p7/8/8/8/8/8/8 b - - 0 1")},
		&Move{s1: A7, s2: A5, state: unsafeFEN("8/p7/8/8/8/8/8/8 b - - 0 1")},
		&Move{s1: C4, s2: B5, state: unsafeFEN("8/8/8/1p1p4/2P5/8/8/8 w - - 0 1")},
		&Move{s1: C4, s2: D5, state: unsafeFEN("8/8/8/1p1p4/2P5/8/8/8 w - - 0 1")},
		&Move{s1: C4, s2: C5, state: unsafeFEN("8/8/8/1p1p4/2P5/8/8/8 w - - 0 1")},
		&Move{s1: C5, s2: B4, state: unsafeFEN("8/8/8/2p5/1P1P4/8/8/8 b - - 0 1")},
		&Move{s1: C5, s2: D4, state: unsafeFEN("8/8/8/2p5/1P1P4/8/8/8 b - - 0 1")},
		&Move{s1: C5, s2: C4, state: unsafeFEN("8/8/8/2p5/1P1P4/8/8/8 b - - 0 1")},
		&Move{s1: A4, s2: B3, state: unsafeFEN("2r3k1/1q1nbppp/r3p3/3pP3/pPpP4/P1Q2N2/2RN1PPP/2R4K b - b3 0 23")},
		&Move{s1: A2, s2: A1, promo: Queen, state: unsafeFEN("8/8/8/8/8/8/p7/8 b - - 0 1")},
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
		// king moves
		&Move{s1: E4, s2: E5, state: unsafeFEN("5r2/8/8/8/4K3/8/8/8 w - - 0 1")},
		&Move{s1: E4, s2: E3, state: unsafeFEN("5r2/8/8/8/4K3/8/8/8 w - - 0 1")},
		&Move{s1: E4, s2: D3, state: unsafeFEN("5r2/8/8/8/4K3/8/8/8 w - - 0 1")},
		&Move{s1: E4, s2: D4, state: unsafeFEN("5r2/8/8/8/4K3/8/8/8 w - - 0 1")},
		&Move{s1: E4, s2: D5, state: unsafeFEN("5r2/8/8/8/4K3/8/8/8 w - - 0 1")},
		// castleing
		&Move{s1: E1, s2: G1, state: unsafeFEN("r3k2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1")},
		&Move{s1: E1, s2: C1, state: unsafeFEN("r3k2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1")},
		&Move{s1: E8, s2: G8, state: unsafeFEN("r3k2r/8/8/8/8/8/8/R3K2R b KQkq - 0 1")},
		&Move{s1: E8, s2: C8, state: unsafeFEN("r3k2r/8/8/8/8/8/8/R3K2R b KQkq - 0 1")},
	}

	invalidMoves = []*Move{
		// out of turn moves
		&Move{s1: E7, s2: E5, state: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		&Move{s1: E2, s2: E4, state: unsafeFEN("rnbqkbnr/1ppppppp/p7/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1")},
		// pawn moves
		&Move{s1: E2, s2: D3, state: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		&Move{s1: E2, s2: F3, state: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		&Move{s1: E2, s2: E5, state: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		&Move{s1: A2, s2: A1, state: unsafeFEN("8/8/8/8/8/8/p7/8 b - - 0 1")},
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
		// king moves
		&Move{s1: E4, s2: F3, state: unsafeFEN("5r2/8/8/8/4K3/8/8/8 w - - 0 1")},
		&Move{s1: E4, s2: F4, state: unsafeFEN("5r2/8/8/8/4K3/8/8/8 w - - 0 1")},
		&Move{s1: E4, s2: F5, state: unsafeFEN("5r2/8/8/8/4K3/8/8/8 w - - 0 1")},
		// castleing
		&Move{s1: E1, s2: B1, state: unsafeFEN("r3k2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1")},
		&Move{s1: E8, s2: B8, state: unsafeFEN("r3k2r/8/8/8/8/8/8/R3K2R b KQkq - 0 1")},
		&Move{s1: E1, s2: C1, state: unsafeFEN("r3k2r/8/8/8/8/8/8/R2QK2R w KQkq - 0 1")},
		&Move{s1: E1, s2: C1, state: unsafeFEN("2r1k2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1")},
		&Move{s1: E1, s2: C1, state: unsafeFEN("3rk2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1")},
		&Move{s1: E1, s2: G1, state: unsafeFEN("r3k2r/8/8/8/8/8/8/R3K2R w Qkq - 0 1")},
		&Move{s1: E1, s2: C1, state: unsafeFEN("r3k2r/8/8/8/8/8/8/R3K2R w Kkq - 0 1")},
		// invalid promotion for non-pawn move
		&Move{s1: B8, s2: D7, promo: Pawn, state: unsafeFEN("rn1qkb1r/pp3ppp/2p1pn2/3p4/2PP4/2NQPN2/PP3PPP/R1B1K2R b KQkq - 0 7")},
	}

	validMoveState = []moveState{
		{
			Move:      &Move{s1: E2, s2: E4, state: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
			PostState: unsafeFEN("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"),
		},
		{
			Move:      &Move{s1: E1, s2: G1, state: unsafeFEN("r3k2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1")},
			PostState: unsafeFEN("r3k2r/8/8/8/8/8/8/R4RK1 b kq - 0 1"),
		},
		{
			Move:      &Move{s1: A4, s2: B3, state: unsafeFEN("2r3k1/1q1nbppp/r3p3/3pP3/pPpP4/P1Q2N2/2RN1PPP/2R4K b - b3 0 23")},
			PostState: unsafeFEN("2r3k1/1q1nbppp/r3p3/3pP3/11pP4/PpQ2N2/2RN1PPP/2R4K w - - 0 24"),
		},
		// castling should reset the half move clock
		{
			Move:      &Move{s1: E1, s2: G1, state: unsafeFEN("r2qk2r/pp1n1ppp/2pbpn2/3p4/2PP4/1PNQPN2/P4PPP/R1B1K2R w KQkq - 1 9")},
			PostState: unsafeFEN("r2qk2r/pp1n1ppp/2pbpn2/3p4/2PP4/1PNQPN2/P4PPP/R1B2RK1 b kq - 0 9"),
		},
	}
)

func unsafeFEN(s string) *GameState {
	g, err := decodeFEN(s)
	if err != nil {
		log.Fatal(err)
	}
	return g
}

func TestValidMoves(t *testing.T) {
	for _, m := range validMoves {
		if !m.isValid() {
			log.Println(m.state.board.Draw())
			t.Fatalf("expected move %s to be valid", m)
		}
	}
}

func TestInvalidMoves(t *testing.T) {
	for _, m := range invalidMoves {
		if m.isValid() {
			log.Println(m.state.board.Draw())
			t.Fatalf("expected move %s to be invalid", m)
		}
	}
}

func TestValidMoveStates(t *testing.T) {
	for _, ms := range validMoveState {
		if !ms.Move.isValid() {
			log.Println(ms.Move.state.board.Draw())
			t.Fatalf("expected move %s to be valid", ms.Move)
		}
		postState := ms.Move.postMoveState()
		if postState.String() != ms.PostState.String() {
			t.Fatalf("starting from board \n%s\n after move %s\n expected board to be %s\n%s\n but was %s\n%s\n",
				ms.Move.state.board.Draw(), ms.Move.String(), ms.PostState.String(),
				ms.PostState.board.Draw(), postState.String(), postState.board.Draw())
		}
	}
}
