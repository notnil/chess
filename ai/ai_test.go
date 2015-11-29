package ai_test

import (
	"testing"

	"github.com/loganjspears/chess"
	"github.com/loganjspears/chess/ai"
)

func BenchmarkAthena(b *testing.B) {
	g := chess.NewGame()
	a := ai.Athena{2}
	for n := 0; n < b.N; n++ {
		a.Move(g.State())
	}
}
