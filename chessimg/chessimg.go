// Package chessimg provides chess image utilities
package chessimg

import (
	"fmt"
	"image/color"
	"io"
	"strings"

	"github.com/loganjspears/chess"
	"github.com/loganjspears/chess/Godeps/_workspace/src/github.com/ajstarks/svgo"
	"github.com/loganjspears/chess/Godeps/_workspace/src/github.com/lucasb-eyer/go-colorful"
	"github.com/loganjspears/chess/chessimg/internal"
)

// A Encoder encodes chess boards into images.
type Encoder struct {
	w     io.Writer
	light color.Color
	dark  color.Color
}

// SquareColors is designed to be used as an optional argument
// to the New function.  It changes the default light and
// dark square colors to the colors given.
func SquareColors(light, dark color.Color) func(*Encoder) {
	return func(e *Encoder) {
		e.light = light
		e.dark = dark
	}
}

// New returns an encoder that writes to the given writer.
// New also takes options which can customize the image
// output.
func New(w io.Writer, options ...func(*Encoder)) *Encoder {
	e := &Encoder{
		w:     w,
		light: color.RGBA{235, 209, 166, 1},
		dark:  color.RGBA{165, 117, 81, 1},
	}
	for _, op := range options {
		op(e)
	}
	return e
}

// EncodeSVG writes the board SVG representation into
// the Encoder's writer.  An error is returned if there
// is there is an error writing data.
func (e *Encoder) EncodeSVG(board chess.Board) error {
	sqSize := 45
	canvas := svg.New(e.w)
	width := 8 * sqSize
	height := 8 * sqSize
	ranks := []string{"8", "7", "6", "5", "4", "3", "2", "1"}
	files := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	lightHex := colorToHex(e.light)
	darkHex := colorToHex(e.dark)

	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height)
	for file := 1; file <= 8; file++ {
		for rank := 1; rank <= 8; rank++ {
			sqHex := lightHex
			txtHex := darkHex
			if ((file + rank) % 2) == 0 {
				sqHex = darkHex
				txtHex = lightHex
			}
			x1 := sqSize * (file - 1)
			y1 := sqSize * (rank - 1)
			canvas.Rect(x1, y1, sqSize, sqSize, "fill: "+sqHex)
			for sq, p := range board {
				hasFile := int(sq.File()) == file
				hasRank := int(sq.Rank()) == 9-rank
				if hasFile && hasRank && p != nil {
					xml := pieceXML(x1, y1, p)
					if _, err := io.WriteString(canvas.Writer, xml); err != nil {
						return err
					}
				}
			}
			if file == 1 {
				style := "font-size:11px;fill: " + txtHex
				canvas.Text(x1+(sqSize*1/20), y1+(sqSize*5/20), ranks[rank-1], style)
			}
			if rank == 8 {
				style := "text-anchor:end;font-size:11px;fill: " + txtHex
				canvas.Text(x1+(sqSize*19/20), y1+sqSize-(sqSize*1/15), files[file-1], style)
			}
		}
	}
	canvas.End()
	return nil
}

func colorToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return colorful.Color{float64(r) / 255.0, float64(g) / 255.0, float64(b) / 255.0}.Hex()
}

func pieceXML(x, y int, p *chess.Piece) string {
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
