// Package image is a go library that creates images from board positions
package image

import (
	"fmt"
	"image/color"
	"io"
	"strings"

	svg "github.com/ajstarks/svgo"
	"github.com/notnil/chess"
	"github.com/notnil/chess/image/internal"
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

// Arrow allocates and returns a new arrow with the default color
func Arrow(from, to chess.Square) arrow {
	return arrow{
		from:  from,
		to:    to,
		color: color.RGBA{247, 181, 75, 1}, // orange
	}
}

// WithColor alters the color of the arrow (the default is a light orange)
func (arrow arrow) WithColor(col color.Color) arrow {
	arrow.color = col
	return arrow
}

// MarkArrows is designed to be used as an optional argument
// to the SVG function.  It marks an arrow between the given
// squares with the an Arrow of the given color.
func MarkArrows(arrows ...arrow) func(*encoder) {
	return func(e *encoder) {
		e.arrows = append(e.arrows, arrows...)
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
	arrows      []arrow
}

// An arrow represents the visualization of a move
type arrow struct {
	from  chess.Square
	to    chess.Square
	color color.Color
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
		arrows:      []arrow{},
	}
	for _, op := range options {
		op(e)
	}
	return e
}

const (
	sqWidth     = 45
	sqHeight    = 45
	sqCenter    = 22.5
	boardWidth  = 8 * sqWidth
	boardHeight = 8 * sqHeight

	arrowWidth   = 8
	arrowOpacity = "75%"
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
	for i, arrow := range e.arrows {
		var xml string
		if isKnightMove(arrow.from, arrow.to) {
			xml = knightArrowXML(i, arrow, e.perspective)
		} else {
			xml = straightArrowXML(i, arrow, e.perspective)
		}
		if _, err := io.WriteString(canvas.Writer, xml); err != nil {
			return err
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

func straightArrowXML(id int, a arrow, perspective chess.Color) string {
	c := colorToHex(a.color)
	x1, y1 := squareCenter(a.from, perspective)
	x2, y2 := squareCenter(a.to, perspective)

	// move the start of the arrow away from the center of the square
	// as well as head of the arrow to the center of its square
	offset := float32(sqCenter / 1.5)
	if x1 == x2 {
		// horizontal line
		if y1 < y2 {
			y1 += offset
			y2 -= offset
		} else {
			y1 -= offset
			y2 += offset
		}
	} else if y1 == y2 {
		// vertical line
		if x1 < x2 {
			x1 += offset
			x2 -= offset
		} else {
			x1 -= offset
			x2 += offset
		}
	} else {
		// diagonal line
		offset /= 2
		if x1 < x2 {
			x1 += offset
			x2 -= offset
		} else {
			x1 -= offset
			x2 += offset
		}

		if y1 < y2 {
			y1 += offset
			y2 -= offset
		} else {
			y1 -= offset
			y2 += offset
		}
	}

	return fmt.Sprintf(`<svg>
  <defs>
    <marker 
      id="head-%d" 
      orient="auto"
      markerWidth="3"
      markerHeight="4"
      refX="0"
      refY="1.5">
      <path d="M0,0 V3 L2,1.5 Z" fill="%s" fill-opacity="%s" />
    </marker>
  </defs>
  <path
    marker-end="url(#head-%d)"
    stroke-width="%dpx"
    d="M%f,%f %f,%f"
    stroke="%s"
    stroke-opacity="%s" />
</svg>`,
		// arrow head
		id,
		c,
		arrowOpacity,

		// line
		id,
		arrowWidth,
		x1,
		y1,
		x2,
		y2,
		c,
		arrowOpacity,
	)
}

func knightArrowXML(id int, a arrow, perspective chess.Color) string {
	c := colorToHex(a.color)

	horizontal := horizontalMoves(a.from, a.to)
	vertical := verticalMoves(a.from, a.to)

	pivot := a.from
	if abs(vertical) == 1 {
		pivot = squareOffset(pivot, horizontal)
	} else {
		if vertical > 0 {
			pivot = squareOffset(pivot, 16)
		} else {
			pivot = squareOffset(pivot, -16)
		}
	}

	x11, y11 := squareCenter(a.from, perspective)
	x12, y12 := squareCenter(pivot, perspective)

	x21, y21 := squareCenter(pivot, perspective)
	x22, y22 := squareCenter(a.to, perspective)

	// move the start and end of the first line further away from the center of the square
	offset := float32(sqCenter / 1.5)
	halfWidth := float32(arrowWidth / 2)
	if abs(horizontal) > abs(vertical) {
		if x11 < x12 {
			x11 += offset
			x12 += halfWidth
		} else {
			x11 -= offset
			x12 -= halfWidth
		}
	} else {
		if y11 < y12 {
			y11 += offset
			y12 += halfWidth
		} else {
			y11 -= offset
			y12 -= halfWidth
		}
	}

	// move the arrow head to the center of the square
	if perspective == chess.Black {
		offset *= -1
	}
	if abs(horizontal) > abs(vertical) {
		if vertical > 0 {
			y21 += halfWidth
			y22 += offset
		} else {
			y21 -= halfWidth
			y22 -= offset
		}
	} else {
		if horizontal > 0 {
			x21 -= halfWidth
			x22 -= offset
		} else {
			x21 += halfWidth
			x22 += offset
		}
	}

	template := `<svg>
  <defs>
    <marker 
      id="head-%d" 
      orient="auto"
      markerWidth="3"
      markerHeight="4"
      refX="0"
      refY="1.5">
      <path d="M0,0 V3 L2,1.5 Z" fill="%s" fill-opacity="%s" />
    </marker>
  </defs>
  <path
    stroke-width="%dpx"
    d="M%f,%f %f,%f"
    stroke="%s"
    stroke-opacity="%s" />
  <path
    marker-end="url(#head-%d)"
    stroke-width="%dpx"
    d="M%f,%f %f,%f"
    stroke="%s"
    stroke-opacity="%s" />
</svg>`

	return fmt.Sprintf(
		template,

		// arrow head
		id,
		c,
		arrowOpacity,

		// first line
		arrowWidth,
		x11,
		y11,
		x12,
		y12,
		c,
		arrowOpacity,

		// second line
		id,
		arrowWidth,
		x21,
		y21,
		x22,
		y22,
		c,
		arrowOpacity,
	)
}

func isKnightMove(from, to chess.Square) bool {
	vertical := abs(verticalMoves(from, to))
	horizontal := abs(horizontalMoves(from, to))
	return (vertical == 1 && horizontal == 2) || (vertical == 2 && horizontal == 1)
}

func verticalMoves(from, to chess.Square) int {
	fromRank := int(from) / 8
	toRank := int(to) / 8
	return toRank - fromRank
}

func horizontalMoves(from, to chess.Square) int {
	fromMod := int(from) % 8
	toMod := int(to) % 8
	return toMod - fromMod
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func squareCenter(sq chess.Square, perspective chess.Color) (float32, float32) {
	if perspective == chess.White {
		x := float32(sq%8)*sqWidth + sqCenter
		y := boardHeight - float32(sq/8)*sqWidth - sqCenter
		return x, y
	} else {
		x := boardWidth - float32(sq%8)*sqWidth - sqCenter
		y := float32(sq/8)*sqWidth + sqCenter
		return x, y
	}
}

func squareOffset(sq chess.Square, offset int) chess.Square {
	return chess.Square(int(sq) + offset)
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
