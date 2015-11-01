// Package chessimg provides chess image utilities
package chessimg

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/ajstarks/svgo"
	"github.com/loganjspears/chess"
	"github.com/lucasb-eyer/go-colorful"
)

// SquareColors is designed to be used as an optional argument
// to the WriteBoardSVG function.  It changes the default light and
// dark square colors to the ones given.
func SquareColors(light, dark color.Color) func(*Encoder) {
	return func(e *Encoder) {
		e.light = light
		e.dark = dark
	}
}

type Encoder struct {
	w     io.Writer
	light color.Color
	dark  color.Color
}

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

func (e *Encoder) EncodeGIF(game *chess.Game, sideLength int, delay time.Duration) error {
	anim := gif.GIF{LoopCount: 1}
	palette := []color.Color{color.White, color.Black, e.light, e.dark}
	for _, state := range game.States() {
		b := &bytes.Buffer{}
		if err := e.copy(b).EncodePNG(state.Board(), sideLength); err != nil {
			return fmt.Errorf("chessimg: failed to encode state %s as png intermediate with error - %s", state, err.Error())
		}
		log.Println(b.String())
		pngImg, err := png.Decode(b)
		if err != nil {
			return fmt.Errorf("chessimg: failed to decode png intermediate for state %s with error - %s", state, err.Error())
		}
		img := image.NewPaletted(pngImg.Bounds(), palette)
		anim.Delay = append(anim.Delay, int(delay.Seconds()*100.0))
		anim.Image = append(anim.Image, img)
	}
	return gif.EncodeAll(e.w, &anim)
}

func (e *Encoder) EncodePNG(board chess.Board, sideLength int) error {
	// check width and height params
	if sideLength <= 0 {
		return errors.New("chessimg: width and height must be greater than zero")
	}
	// create temp svg file
	svgFile, err := os.Create("temp.svg")
	if err != nil {
		return fmt.Errorf("chessimg: failed to create svg temp file with error - %s", err.Error())
	}
	defer svgFile.Close()
	// write svg to temp
	svgEn := e.copy(svgFile)
	if err := svgEn.EncodeSVG(board); err != nil {
		return err
	}
	// rsvg-convert -w 400 board.svg -o board.png
	// http://manpages.ubuntu.com/manpages/precise/man1/rsvg-convert.1.html
	l := fmt.Sprint(sideLength)
	if err := exec.Command("rsvg-convert", "-w", l, "-h", l, "temp.svg", "-o", "board.png").Run(); err != nil {
		return fmt.Errorf("chessimg: rsvg-convert command failed with error - %s", err.Error())
	}
	// write png to temp
	pngFile, err := os.Create("temp.png")
	if err != nil {
		return fmt.Errorf("chessimg: failed to create png temp file with error - %s", err.Error())
	}
	defer pngFile.Close()
	// copy png temp to writer
	_, err = io.Copy(e.w, pngFile)
	if err != nil {
		return fmt.Errorf("chessimg: failed to copy png temp file to writer with error - %s", err.Error())
	}
	// delete temp files
	if err := os.Remove("temp.svg"); err != nil {
		return fmt.Errorf("chessimg: failed to delete svg temp file with error - %s", err.Error())
	}
	if err := os.Remove("temp.png"); err != nil {
		return fmt.Errorf("chessimg: failed to delete png temp file with error - %s", err.Error())
	}
	return nil
}

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

func (e *Encoder) copy(w io.Writer) *Encoder {
	return &Encoder{
		w:     w,
		light: e.light,
		dark:  e.dark,
	}
}

func colorToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return colorful.Color{float64(r) / 255.0, float64(g) / 255.0, float64(b) / 255.0}.Hex()
}

func pieceXML(x, y int, p *chess.Piece) string {
	fileName := fmt.Sprintf("pieces/%s%s.svg", p.Color().String(), pieceTypeMap[p.Type()])
	svgStr := string(MustAsset(fileName))
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
