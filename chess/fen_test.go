package chess

import "testing"

var (
	validFENs = []string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
		"rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2",
		"rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2",
		"7k/8/8/8/8/8/8/R6K w - - 0 1",
		"7k/8/8/8/8/8/8/2B1KB2 w - - 0 1",
		"8/8/8/4k3/8/8/8/R3K2R w KQ - 0 1",
		"8/8/8/8/4k3/8/3KP3/8 w - - 0 1",
		"8/8/5k2/8/5K2/8/4P3/8 w - - 0 1",
		"r4rk1/1b2bppp/ppq1p3/2pp3n/5P2/1P1BP3/PBPPQ1PP/R4RK1 w - - 0 1",
		"3r1rk1/p3qppp/2bb4/2p5/3p4/1P2P3/PBQN1PPP/2R2RK1 w - - 0 1",
		"4r1k1/1b3p1p/ppq3p1/2p5/8/1P3R1Q/PBP3PP/7K w - - 0 1",
		"5k2/ppp5/4P3/3R3p/6P1/1K2Nr2/PP3P2/8 b - - 1 32",
	}

	invalidFENs = []string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPP/RNBQKBNR w KQkq - 0 1",
		"rnbqkbnr/pppppppp/8/8/4P2/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
		"rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KKkq c6 0 2",
		"rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq c12 1 2",
		"7k/8/8/8/8/8/8/R6K w - - 0 -1",
		"7k/8/8/8/8/8/8/2B1KB2 w - - -1 1",
		"8/8/8/4k3/8/8/8/R3K2R w KQ - 0 0",
		"8/8/8/8/4k3/8/3KP3/8 c - - 0 1",
		"8/8/5k2/8/5K2/8/4P3P/8 w - - 0 1",
		"r4rk1/1b2bppp/ppq1p3/2pp3n/5P2/1P1BP3/PBPPQ1PP/R4RK1 w e4 - 0 1",
	}
)

func TestValidFENs(t *testing.T) {
	for _, f := range validFENs {
		state, err := decodeFEN(f)
		if err != nil {
			t.Fatal("recieved unexpected error", err)
		}
		if f != state.String() {
			t.Fatalf("fen expected board string %s but got %s", f, state.String())
		}
	}
}

func TestInvalidFENs(t *testing.T) {
	for _, f := range invalidFENs {
		if _, err := decodeFEN(f); err == nil {
			t.Fatal("fen expected error from ", f)
		}
	}
}
