package image_test

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"image/color"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/notnil/chess"
	"github.com/notnil/chess/image"
)

const expectedMD5 = "a2ee66ca19e4c347aec41371c1ca07f8"
const expectedMD5Black = "ce4d4e033a50678898c62928b8e0a15c"
const expectedMD5KnightsAndDiagonalArrows = "9c95aa56cec67be2ceee141f259753f7"

func TestSVG(t *testing.T) {
	// create buffer of actual svg
	buf := bytes.NewBuffer([]byte{})
	fenStr := "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq - 0 1"
	pos := &chess.Position{}
	if err := pos.UnmarshalText([]byte(fenStr)); err != nil {
		t.Error(err)
	}
	mark := image.MarkSquares(color.RGBA{255, 255, 0, 1}, chess.D2, chess.D4)
	arrows := image.MarkArrows(image.Arrow(chess.D2, chess.D4))
	if err := image.SVG(buf, pos.Board(), mark, arrows); err != nil {
		t.Error(err)
	}

	// compare to expected svg
	actualSVG := strings.TrimSpace(buf.String())
	actualMD5 := fmt.Sprintf("%x", md5.Sum([]byte(actualSVG)))
	if actualMD5 != expectedMD5 {
		t.Errorf("expected actual md5 hash to be %s but got %s", expectedMD5, actualMD5)
	}

	// create actual svg file for visualization
	f, err := os.Create("example.svg")
	defer f.Close()
	if err != nil {
		t.Error(err)
	}
	if _, err := io.Copy(f, bytes.NewBufferString(actualSVG)); err != nil {
		t.Error(err)
	}
}

func TestSVGFromBlack(t *testing.T) {
	// create buffer of actual svg
	buf := bytes.NewBuffer([]byte{})
	fenStr := "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq - 0 1"
	pos := &chess.Position{}
	if err := pos.UnmarshalText([]byte(fenStr)); err != nil {
		t.Error(err)
	}
	mark := image.MarkSquares(color.RGBA{255, 255, 0, 1}, chess.D2, chess.D4)
	arrows := image.MarkArrows(image.Arrow(chess.D2, chess.D4).WithColor(color.Black))
	per := image.Perspective(chess.Black)
	if err := image.SVG(buf, pos.Board(), mark, arrows, per); err != nil {
		t.Error(err)
	}

	// compare to expected svg
	actualSVG := strings.TrimSpace(buf.String())
	actualMD5 := fmt.Sprintf("%x", md5.Sum([]byte(actualSVG)))
	if actualMD5 != expectedMD5Black {
		t.Errorf("expected actual md5 hash to be %s but got %s", expectedMD5Black, actualMD5)
	}

	// create actual svg file for visualization
	f, err := os.Create("black_example.svg")
	defer f.Close()
	if err != nil {
		t.Error(err)
	}
	if _, err := io.Copy(f, bytes.NewBufferString(actualSVG)); err != nil {
		t.Error(err)
	}
}

func TestSVGKnightsAndDiagonals(t *testing.T) {
	// create buffer of actual svg
	buf := bytes.NewBuffer([]byte{})
	fenStr := "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq - 0 1"
	pos := &chess.Position{}
	if err := pos.UnmarshalText([]byte(fenStr)); err != nil {
		t.Error(err)
	}
	arrows := image.MarkArrows(
		// all possible knight directions
		image.Arrow(chess.F6, chess.E4),
		image.Arrow(chess.F6, chess.D5),
		image.Arrow(chess.F6, chess.D7),
		image.Arrow(chess.F6, chess.E8),
		image.Arrow(chess.F6, chess.G4),
		image.Arrow(chess.F6, chess.H5),
		image.Arrow(chess.F6, chess.H7),
		image.Arrow(chess.F6, chess.G8),

		// a couple knight moves with no overlapping arrows
		image.Arrow(chess.B1, chess.D2),
		image.Arrow(chess.B8, chess.C6),

		// diagonal arrows
		image.Arrow(chess.C4, chess.A6),
		image.Arrow(chess.C4, chess.D5),

		// anti-diagonal arrows
		image.Arrow(chess.C4, chess.A2),
		image.Arrow(chess.C4, chess.D3),
	)
	per := image.Perspective(chess.Black)
	if err := image.SVG(buf, pos.Board(), arrows, per); err != nil {
		t.Error(err)
	}

	// compare to expected svg
	actualSVG := strings.TrimSpace(buf.String())
	actualMD5 := fmt.Sprintf("%x", md5.Sum([]byte(actualSVG)))
	if actualMD5 != expectedMD5KnightsAndDiagonalArrows {
		t.Errorf("expected actual md5 hash to be %s but got %s", expectedMD5KnightsAndDiagonalArrows, actualMD5)
	}

	// create actual svg file for visualization
	f, err := os.Create("knight_arrows_example.svg")
	defer f.Close()
	if err != nil {
		t.Error(err)
	}
	if _, err := io.Copy(f, bytes.NewBufferString(actualSVG)); err != nil {
		t.Error(err)
	}
}
