package chessimg

import (
	"os"
	"testing"

	"github.com/loganjspears/chess"
)

func TestSVG(t *testing.T) {
	game := chess.NewGame()
	f, _ := os.Create("test.svg")
	defer f.Close()

	if err := New(f).EncodeSVG(game.State().Board()); err != nil {
		t.Fatal(err)
	}
}
