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

const expectedMD5 = "da140af8b83ce7903915ee39973e36dd"

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
