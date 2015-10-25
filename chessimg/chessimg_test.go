package chessimg

import (
	"os"
	"testing"

	"github.com/loganjspears/chess"
)

func TestImg(t *testing.T) {
	game := chess.NewGame()
	f, _ := os.Create("board.svg")
	defer f.Close()
	WriteBoardSVG(f, game.State().Board())
}
