package chess

import (
	"strconv"
	"strings"
)

// bitboard is a board representation encoded in an unsigned 64-bit integer.  The
// 64 board positions begin with A1 as the most significant bit and H8 as the least.
type bitboard uint64

func newBitboard(m map[Square]bool) bitboard {
	s := ""
	for sq := 0; sq < numOfSquaresInBoard; sq++ {
		if m[Square(sq)] {
			s += "1"
		} else {
			s += "0"
		}
	}
	bb, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return bitboard(bb)
}

func (b bitboard) MSB() int {
	i := strings.Index(b.String(), "1")
	if i == -1 {
		return 0
	}
	return i
}

func (b bitboard) Mapping() map[Square]bool {
	s := b.String()
	m := map[Square]bool{}
	for sq := 0; sq < numOfSquaresInBoard; sq++ {
		if s[sq:sq+1] == "1" {
			m[Square(sq)] = true
		}
	}
	return m
}

func (b bitboard) Occupied(sq Square) bool {
	return (uint64(b) >> uint64(63-sq) & 1) == 1
}

func (b bitboard) String() string {
	s := strconv.FormatUint(uint64(b), 2)
	return strings.Repeat("0", numOfSquaresInBoard-len(s)) + s
}

func (b bitboard) Reverse() bitboard {
	var u uint64
	for sq := 0; sq < 64; sq++ {
		u = (u << 1) + (uint64(b) & 1)
		b = b >> 1
	}
	return bitboard(u)
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
func (b bitboard) Draw() string {
	s := "\n A B C D E F G H\n"
	for r := 7; r >= 0; r-- {
		s += rank(r).String()
		for f := 0; f < numOfSquaresInRow; f++ {
			sq := getSquare(file(f), rank(r))
			if b.Occupied(sq) {
				s += "1"
			} else {
				s += "0"
			}
			s += " "
		}
		s += "\n"
	}
	return s
}
