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
