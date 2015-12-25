package ai_test

import (
	"strings"
	"testing"

	"github.com/loganjspears/chess"
	"github.com/loganjspears/chess/ai"
)

func TestMinMax(t *testing.T) {
	fenStr := `r1bqkbnr/p2pppp1/1pn5/3N3p/4PB2/5N2/PPP2PPP/R2QKB1R w KQkq - 3 8`
	fen, err := chess.FEN(strings.NewReader(fenStr))
	if err != nil {
		t.Fatal(err)
	}
	g := chess.NewGame(fen)

	a := ai.New(2)
	m := a.Move(g.State())
	if m.S1() == chess.D8 && m.S2() == chess.C7 {
		t.Fatal("athena should not take the knight")
	}
}

func BenchmarkAthena(b *testing.B) {
	g := chess.NewGame()
	a := ai.New(3)
	for n := 0; n < b.N; n++ {
		a.Move(g.State())
		b.Log("POS", a.PositionsEvaluated())
	}
}
