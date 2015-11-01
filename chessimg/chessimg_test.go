package chessimg

import (
	"bytes"
	"os"
	"testing"

	"github.com/loganjspears/chess"
)

func TestSVG(t *testing.T) {
	game := chess.NewGame()
	f, _ := os.Create("board.svg")
	defer f.Close()

	if err := New(f).EncodeSVG(game.State().Board()); err != nil {
		t.Fatal(err)
	}
}

func TestPNG(t *testing.T) {
	game := chess.NewGame()
	f, _ := os.Create("board.png")
	defer f.Close()

	if err := New(f).EncodePNG(game.State().Board(), 500); err != nil {
		t.Fatal(err)
	}
}

func TestPNGBytes(t *testing.T) {
	game := chess.NewGame()
	b := bytes.NewBuffer([]byte{})
	if err := New(b).EncodePNG(game.State().Board(), 500); err != nil {
		t.Fatal(err)
	}
	if len(b.Bytes()) == 0 {
		t.Fatal("No bytes")
	}
}

//
// func TestGIF(t *testing.T) {
// 	game := chess.NewGame()
// 	f, _ := os.Create("game.gif")
// 	defer f.Close()
//
// 	if err := New(f).EncodeGIF(game, 500, time.Second); err != nil {
// 		t.Fatal(err)
// 	}
// }
