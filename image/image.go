// Package image is a go library that creates images from board positions
package image

import (
	"fmt"
	"image/color"
	"io"
	"strings"

	svg "github.com/ajstarks/svgo"
	"github.com/torresomarmx/chess-lib"
	"github.com/torresomarmx/chess-lib/image/internal"
)

// SVG writes the board SVG representation into the writer.
// An error is returned if there is there is an error writing data.
// SVG also takes options which can customize the image output.
func SVG(w io.Writer, b *chess.Board, opts ...func(*encoder)) error {
	e := new(w, opts)
	return e.EncodeSVG(b)
}

// SquareColors is designed to be used as an optional argument
// to the SVG function.  It changes the default light and
// dark square colors to the colors given.
func SquareColors(light, dark color.Color) func(*encoder) {
	return func(e *encoder) {
		e.light = light
		e.dark = dark
	}
}

// MarkSquares is designed to be used as an optional argument
// to the SVG function.  It marks the given squares with the
// color.  A possible usage includes marking squares of the
// previous move.
func MarkSquares(c color.Color, sqs ...chess.Square) func(*encoder) {
	return func(e *encoder) {
		for _, sq := range sqs {
			e.marks[sq] = c
		}
	}
}

// Perspective is designed to be used as an optional argument
// to the SVG function.  It draws the board from the perspective
// of the given color.  White is the default.
func Perspective(c chess.Color) func(*encoder) {
	return func(e *encoder) {
		e.perspective = c
	}
}

// A Encoder encodes chess boards into images.
type encoder struct {
	w           io.Writer
	light       color.Color
	dark        color.Color
	perspective chess.Color
	marks       map[chess.Square]color.Color
}

// New returns an encoder that writes to the given writer.
// New also takes options which can customize the image
// output.
func new(w io.Writer, options []func(*encoder)) *encoder {
	e := &encoder{
		w:           w,
		light:       color.RGBA{235, 209, 166, 1},
		dark:        color.RGBA{165, 117, 81, 1},
		perspective: chess.White,
		marks:       map[chess.Square]color.Color{},
	}
	for _, op := range options {
		op(e)
	}
	return e
}

const (
	sqWidth     = 45
	sqHeight    = 45
	boardWidth  = 8 * sqWidth
	boardHeight = 8 * sqHeight
)

var (
	orderOfRanks      = []chess.Rank{chess.Rank8, chess.Rank7, chess.Rank6, chess.Rank5, chess.Rank4, chess.Rank3, chess.Rank2, chess.Rank1}
	orderOfRanksBlack = []chess.Rank{chess.Rank1, chess.Rank2, chess.Rank3, chess.Rank4, chess.Rank5, chess.Rank6, chess.Rank7, chess.Rank8}
	orderOfFiles      = []chess.File{chess.FileA, chess.FileB, chess.FileC, chess.FileD, chess.FileE, chess.FileF, chess.FileG, chess.FileH}
	orderOfFilesBlack = []chess.File{chess.FileH, chess.FileG, chess.FileF, chess.FileE, chess.FileD, chess.FileC, chess.FileB, chess.FileA}
)

// EncodeSVG writes the board SVG representation into
// the Encoder's writer.  An error is returned if there
// is there is an error writing data.
func (e *encoder) EncodeSVG(b *chess.Board) error {
	boardMap := b.SquareMap()
	canvas := svg.New(e.w)
	canvas.Start(boardWidth, boardHeight)
	canvas.Rect(0, 0, boardWidth, boardHeight)

	ranks := orderOfRanks
	files := orderOfFiles
	if e.perspective == chess.Black {
		ranks = orderOfRanksBlack
		files = orderOfFilesBlack
	}
	for i, rank := range ranks {
		for j, file := range files {
			x := j * sqHeight
			y := i * sqHeight
			sq := chess.NewSquare(file, rank)
			c := e.colorForSquare(sq)
			canvas.Rect(x, y, sqWidth, sqHeight, "fill: "+colorToHex(c))
			markColor, ok := e.marks[sq]
			if ok {
				canvas.Rect(x, y, sqWidth, sqHeight, "fill-opacity:0.2;fill: "+colorToHex(markColor))
			}
			// draw piece
			p := boardMap[sq]
			if p != chess.NoPiece {
				xml := pieceXML(x, y, p)
				if _, err := io.WriteString(canvas.Writer, xml); err != nil {
					return err
				}
			}
			// draw rank text on file A
			txtColor := e.colorForText(sq)
			if j == 0 {
				style := "font-size:11px;fill: " + colorToHex(txtColor)
				canvas.Text(x+(sqWidth*1/20), y+(sqHeight*5/20), sq.Rank().String(), style)
			}
			// draw file text on rank 1
			if i == 7 {
				style := "text-anchor:end;font-size:11px;fill: " + colorToHex(txtColor)
				canvas.Text(x+(sqWidth*19/20), y+sqHeight-(sqHeight*1/15), sq.File().String(), style)
			}
		}
	}
	canvas.End()
	return nil
}

func (e *encoder) colorForSquare(sq chess.Square) color.Color {
	sqSum := int(sq.File()) + int(sq.Rank())
	if sqSum%2 == 0 {
		return e.dark
	}
	return e.light
}

func (e *encoder) colorForText(sq chess.Square) color.Color {
	sqSum := int(sq.File()) + int(sq.Rank())
	if sqSum%2 == 0 {
		return e.light
	}
	return e.dark
}

func colorToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", uint8(float64(r)+0.5), uint8(float64(g)*1.0+0.5), uint8(float64(b)*1.0+0.5))
}

func pieceXML(x, y int, p chess.Piece) string {
	fileName := fmt.Sprintf("pieces/%s%s.svg", p.Color().String(), pieceTypeMap[p.Type()])
	svgStr := string(internal.MustAsset(fileName))
	old := `<svg xmlns="http://www.w3.org/2000/svg" version="1.1" width="45" height="45">`
	new := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" version="1.1" width="360" height="360" viewBox="%d %d 360 360">`, (-1 * x), (-1 * y))
	return strings.Replace(svgStr, old, new, 1)
}

var (
	pieceTypeMap = map[chess.PieceType]string{
		chess.King:   "K",
		chess.Queen:  "Q",
		chess.Rook:   "R",
		chess.Bishop: "B",
		chess.Knight: "N",
		chess.Pawn:   "P",
	}
)
