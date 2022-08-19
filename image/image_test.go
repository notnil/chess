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

	"github.com/torresomarmx/chess-lib"
	"github.com/torresomarmx/chess-lib/image"
)

const expectedMD5 = "08aaa6fcfde3bb900fc54bdfef3d5c81"
const expectedMD5Black = "badac5ca5cfbdea9b98a1f9988ba54bc"

func TestSVG(t *testing.T) {
	// create buffer of actual svg
	buf := bytes.NewBuffer([]byte{})
	fenStr := "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq - 0 1"
	pos := &chess.Position{}
	if err := pos.UnmarshalText([]byte(fenStr)); err != nil {
		t.Error(err)
	}
	mark := image.MarkSquares(color.RGBA{255, 255, 0, 1}, chess.D2, chess.D4)
	if err := image.SVG(buf, pos.Board(), mark); err != nil {
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
	per := image.Perspective(chess.Black)
	if err := image.SVG(buf, pos.Board(), mark, per); err != nil {
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
