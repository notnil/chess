package chess

import (
	"strconv"
	"strings"
)

type Board map[*Square]*Piece

// String implements the fmt.Stringer interface and returns
// a string in the format: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR
func (b Board) String() string {
	rankStrs := []string{}
	for _, r := range []Rank{R8, R7, R6, R5, R4, R3, R2, R1} {
		s := ""
		skip := 0
		for _, f := range Files() {
			sq := getSquare(f, r)
			p := b.piece(sq)
			if p == nil {
				skip++
				continue
			}
			if skip != 0 {
				s += strconv.Itoa(skip)
				skip = 0
			}
			s += p.getFENChar()
		}
		if skip > 0 {
			s += strconv.Itoa(skip)
		}
		rankStrs = append(rankStrs, s)
	}
	return strings.Join(rankStrs, "/")
}

// Draw returns visual representation of the board useful for debugging.  Ex.
// A B C D E F G H
// 8- - - - - - - -
// 7- ♟ - - - - - ♝
// 6- - - - ♘ - - -
// 5- - - - ♟ - - -
// 4- - - - - - - -
// 3- - - - - - - -
// 2- ♖ - - - - - -
// 1- ♗ - - - - - -
func (b Board) Draw() string {
	s := "\n A B C D E F G H\n"
	for _, r := range []Rank{R8, R7, R6, R5, R4, R3, R2, R1} {
		s += r.String()
		for _, f := range Files() {
			p := b.piece(getSquare(f, r))
			if p == nil {
				s += "-"
			} else {
				s += p.String()
			}
			s += " "
		}
		s += "\n"
	}
	return s
}

func (b Board) isOccupied(sq *Square) bool {
	return b[sq] != nil
}

func (b Board) piece(sq *Square) *Piece {
	return b[sq]
}

func (b Board) copy() Board {
	n := map[*Square]*Piece{}
	for k, v := range b {
		n[k] = v
	}
	return n
}

// squaresForColor returns a slice of the squares that are
// occupied by the given color.
func (b Board) squaresForColor(c Color) []*Square {
	squares := []*Square{}
	for _, sq := range allSquares {
		if b.isOccupied(sq) && b.piece(sq).Color() == c {
			squares = append(squares, sq)
		}
	}
	return squares
}

// kingSquare returns the square that the king, of the given color,
// is placed on.  In non-traditional games it will return the square
// of the first king it finds.  If no king is on the board it will
// return nil.
func (b Board) kingSquare(c Color) *Square {
	for _, sq := range b.squaresForColor(c) {
		if b.piece(sq).Type() == King {
			return sq
		}
	}
	return nil
}

// isSquareAttacked returns whether the other color can move to or capture
// the square.  It doesn't include game state in its calculation so it isn't
// sufficient for checkmate or stalemate calculations.
func (b Board) isSquareAttacked(c Color, s2 *Square) bool {
	for _, s1 := range b.squaresForColor(c.Other()) {
		m := &Move{s1: s1, s2: s2, state: &GameState{Board: b, Turn: c.Other()}}
		if m.isValid() {
			return true
		}
	}
	return false
}

// inCheck returns whether the color is in check.
func (b Board) inCheck(c Color) bool {
	kingSq := b.kingSquare(c)
	if kingSq == nil {
		return false
	}
	return b.isSquareAttacked(c, kingSq)
}
