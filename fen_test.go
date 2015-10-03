package chess_test

import (
	"testing"

	. "github.com/loganjspears/chess"
)

var (
	testFENs = []string{
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
)

func TestFEN(t *testing.T) {
	for _, fen := range testFENs {
		state, err := FEN(fen)
		if err != nil {
			t.Fatal("recieved unexpected error", err)
		}
		if fen != state.String() {
			t.Fatalf("fen expected board string %s but got %s", fen, state.String())
		}
	}
}
