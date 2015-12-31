package chess

import (
	"strconv"
	"strings"
)

// A Board represents a chess board with the relationships between squares and pieces.
type Board struct {
	pieces map[*Square]*Piece
	fenStr string
}

// Pieces returns the mapping of squares to pieces.
func (b *Board) Pieces() map[*Square]*Piece {
	cp := map[*Square]*Piece{}
	for k, v := range b.pieces {
		cp[k] = v
	}
	return cp
}

// String implements the fmt.Stringer interface and returns
// a string in the FEN board format: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR
func (b *Board) String() string {
	if b.fenStr != "" {
		return b.fenStr
	}
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
	b.fenStr = strings.Join(rankStrs, "/")
	return b.fenStr
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
func (b *Board) Draw() string {
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

// squaresForColor returns a slice of the squares that are
// occupied by the given color.
func (b *Board) squaresForColor(c Color) []*Square {
	squares := []*Square{}
	for _, sq := range allSquares {
		occupied := b.isOccupied(sq)
		if !occupied && c == NoColor {
			squares = append(squares, sq)
		} else if occupied && b.piece(sq).Color() == c {
			squares = append(squares, sq)
		}
	}
	return squares
}

// kingSquare returns the square that the king, of the given color,
// is placed on.  In non-traditional games it will return the square
// of the first king it finds.  If no king is on the board it will
// return nil.
func (b *Board) kingSquare(c Color) *Square {
	for _, sq := range allSquares {
		p := b.piece(sq)
		if p != nil && p.Type() == King && p.Color() == c {
			return sq
		}
	}
	return nil
}

// isSquareAttacked returns whether the other color can move to or capture
// one of the given squares.  It doesn't include game state in its calculation
// so it isn't sufficient for checkmate or stalemate calculations.
func (b *Board) isSquareAttacked(c Color, squares ...*Square) bool {
	for _, s1 := range b.squaresForColor(c.Other()) {
		for _, s2 := range squares {
			m := &Move{s1: s1, s2: s2, state: &GameState{board: b, turn: c.Other()}}
			pieceType := m.piece().Type()
			if (pieceType != Pawn || (pieceType == Pawn && pawnCaptureFilter(m))) && m.isValidWithCheck(false) {
				return true
			}
		}
	}
	return false
}

// inCheck returns whether the color is in check.
func (b *Board) inCheck(c Color) bool {
	kingSq := b.kingSquare(c)
	if kingSq == nil {
		return false
	}
	return b.isSquareAttacked(c, kingSq)
}

// emptyBetween returns true if the squares between
// s1 and s2 are empty, otherwise returns false.
func (b *Board) emptyBetween(s1, s2 *Square) bool {
	for _, sq := range s1.squaresTo(s2) {
		if b.isOccupied(sq) {
			return false
		}
	}
	return true
}

// isOccupied returns true if a piece is occupying
// the given square, otherwise returns false.
func (b *Board) isOccupied(sq *Square) bool {
	return b.pieces[sq] != nil
}

// piece returns the piece at the given square.
// If no piece is occupying the square, nil is
// returned.
func (b *Board) piece(sq *Square) *Piece {
	return b.pieces[sq]
}

// copy returns a copy of the board.
func (b *Board) copy() *Board {
	return &Board{pieces: b.Pieces()}
}

type squareColor int

const (
	noSquareColor squareColor = iota
	lightSquare
	darkSquare
)

func (b *Board) hasSufficientMaterial() bool {
	count := map[PieceType]int{}
	for _, p := range b.pieces {
		count[p.Type()]++
	}
	if count[Pawn] > 0 || count[Queen] > 0 || count[Rook] > 0 {
		return true
	}
	// 	king versus king
	if count[King] == 2 && count[Bishop] == 0 && count[Knight] == 0 {
		return false
	}
	// king and bishop versus king
	if count[King] == 2 && count[Bishop] == 1 && count[Knight] == 0 {
		return false
	}
	// king and knight versus king
	if count[King] == 2 && count[Bishop] == 0 && count[Knight] == 1 {
		return false
	}
	// king and bishop(s) versus king and bishop(s) with the bishops on the same colour.
	if count[King] == 2 && count[Knight] == 0 {
		color := noSquareColor
		for sq, p := range b.pieces {
			if p.Type() != Bishop {
				continue
			}
			c := colorForSquare(sq)
			if color == noSquareColor {
				color = c
				continue
			}
			if c != color {
				return true
			}
		}
		return false
	}
	return true
}

func colorForSquare(sq *Square) squareColor {
	if (int(sq.file)+int(sq.rank))%2 == 0 {
		return darkSquare
	}
	return lightSquare
}
