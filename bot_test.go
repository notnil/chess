package chess_test

import (
	"log"
	"testing"

	"github.com/notnil/chess"
)

func TestOpening(t *testing.T) {
	bot := &chess.Bot{UseBook: true}
	m, _ := bot.Move(chess.NewGame())
	if m.S2() != chess.E4 {
		t.Fatal("should start with e4")
	}
}

func TestOpening2(t *testing.T) {
	// r1bqkb1r/pppp1ppp/2n2n2/4p3/2B1P3/2N2N2/PPPP1PPP/R1BQK2R b KQkq - 5 4
}

func TestTactics(t *testing.T) {
	type tacticTest struct {
		preFEN  string
		postFEN string
	}
	tacticTests := []tacticTest{
		// {`r6Q/1b2k3/3p4/2q5/1p2PP2/2n5/PKPN4/7R w - - 0 1`, "r7/1b2k1Q1/3p4/2q5/1p2PP2/2n5/PKPN4/7R b - - 1 1"},
		// {`8/1p2r3/1R4pk/p1p5/PnP5/1P2NK1p/8/8 w - - 0 44`, "8/1p2r3/1R4pk/p1p2N2/PnP5/1P3K1p/8/8 b - - 1 44"},
		// {`8/6bp/6p1/3b3k/5Qp1/5P1P/5q2/2R4K w - - 0 38`, "8/6bp/6p1/3b3k/5QP1/5P2/5q2/2R4K b - - 0 38"},
		{`2k1N3/pppr4/1b6/5P2/6n1/2P3q1/PP4B1/1K3Q1R b - - 2 30`, ""},
	}
	for _, tactic := range tacticTests {
		bot := &chess.Bot{}
		fen, err := chess.FEN(tactic.preFEN)
		if err != nil {
			t.Fatal(err)
		}
		g := chess.NewGame(fen)
		m, score := bot.Move(g)
		log.Println(m, score)
		actual := g.Position().Update(m).String()
		if tactic.postFEN != actual {
			t.Fatalf("from %s expected bot to move to %s but moved %s to %s", tactic.preFEN, tactic.postFEN, m.String(), actual)
		}
	}
}

func BenchmarkTactics(b *testing.B) {
	bot := &chessbot.Bot{}
	fen, err := chess.FEN(`2k1N3/pppr4/1b6/5P2/6n1/2P3q1/PP4B1/1K3Q1R b - - 2 30`)
	if err != nil {
		b.Fatal(err)
	}
	g := chess.NewGame(fen)
	b.ResetTimer()
	bot.Move(g)
}
