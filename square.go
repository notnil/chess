package chess

type File int

const (
	A File = iota + 1
	B
	C
	D
	E
	F
	G
	H
)

func (f File) String() string {
	s := "abcdefgh"
	i := int(f)
	return s[i-1 : i]
}

func Files() [8]File {
	return [8]File{A, B, C, D, E, F, G, H}
}

type Rank int

const (
	R1 Rank = iota + 1
	R2
	R3
	R4
	R5
	R6
	R7
	R8
)

func (r Rank) String() string {
	s := "12345678"
	i := int(r)
	return s[i-1 : i]
}

func Ranks() [8]Rank {
	return [8]Rank{R1, R2, R3, R4, R5, R6, R7, R8}
}

type Square struct {
	file File
	rank Rank
}

func (s *Square) String() string {
	return s.file.String() + s.rank.String()
}

func squareFromStr(s string) *Square {
	for _, sq := range allSquares {
		if sq.String() == s {
			return sq
		}
	}
	return nil
}

func getSquare(f File, r Rank) *Square {
	for _, s := range allSquares {
		if s.file == f && s.rank == r {
			return s
		}
	}
	return nil
}

func (s *Square) fileDif(o *Square) int {
	return abs(int(s.file) - int(o.file))
}

func (s *Square) rankDif(o *Square) int {
	return abs(int(s.rank) - int(o.rank))
}

func (s *Square) squaresTo(o *Square) []*Square {
	fileStep := sign(int(o.file) - int(s.file))
	rankStep := sign(int(o.rank) - int(s.rank))

	squares := []*Square{}
	f := int(s.file)
	r := int(s.rank)
	for f != int(o.file) || r != int(o.rank) {
		sq := getSquare(File(f), Rank(r))
		if sq != s && sq != o {
			squares = append(squares, sq)
		}
		f += fileStep
		r += rankStep
	}
	return squares
}

type squareSlice []*Square

func (a squareSlice) has(s *Square) bool {
	for _, sq := range a {
		if sq == s {
			return true
		}
	}
	return false
}

var (
	A1 = &Square{A, R1}
	A2 = &Square{A, R2}
	A3 = &Square{A, R3}
	A4 = &Square{A, R4}
	A5 = &Square{A, R5}
	A6 = &Square{A, R6}
	A7 = &Square{A, R7}
	A8 = &Square{A, R8}

	B1 = &Square{B, R1}
	B2 = &Square{B, R2}
	B3 = &Square{B, R3}
	B4 = &Square{B, R4}
	B5 = &Square{B, R5}
	B6 = &Square{B, R6}
	B7 = &Square{B, R7}
	B8 = &Square{B, R8}

	C1 = &Square{C, R1}
	C2 = &Square{C, R2}
	C3 = &Square{C, R3}
	C4 = &Square{C, R4}
	C5 = &Square{C, R5}
	C6 = &Square{C, R6}
	C7 = &Square{C, R7}
	C8 = &Square{C, R8}

	D1 = &Square{D, R1}
	D2 = &Square{D, R2}
	D3 = &Square{D, R3}
	D4 = &Square{D, R4}
	D5 = &Square{D, R5}
	D6 = &Square{D, R6}
	D7 = &Square{D, R7}
	D8 = &Square{D, R8}

	E1 = &Square{E, R1}
	E2 = &Square{E, R2}
	E3 = &Square{E, R3}
	E4 = &Square{E, R4}
	E5 = &Square{E, R5}
	E6 = &Square{E, R6}
	E7 = &Square{E, R7}
	E8 = &Square{E, R8}

	F1 = &Square{F, R1}
	F2 = &Square{F, R2}
	F3 = &Square{F, R3}
	F4 = &Square{F, R4}
	F5 = &Square{F, R5}
	F6 = &Square{F, R6}
	F7 = &Square{F, R7}
	F8 = &Square{F, R8}

	G1 = &Square{G, R1}
	G2 = &Square{G, R2}
	G3 = &Square{G, R3}
	G4 = &Square{G, R4}
	G5 = &Square{G, R5}
	G6 = &Square{G, R6}
	G7 = &Square{G, R7}
	G8 = &Square{G, R8}

	H1 = &Square{H, R1}
	H2 = &Square{H, R2}
	H3 = &Square{H, R3}
	H4 = &Square{H, R4}
	H5 = &Square{H, R5}
	H6 = &Square{H, R6}
	H7 = &Square{H, R7}
	H8 = &Square{H, R8}

	allSquares = [64]*Square{
		A1, A2, A3, A4, A5, A6, A7, A8,
		B1, B2, B3, B4, B5, B6, B7, B8,
		C1, C2, C3, C4, C5, C6, C7, C8,
		D1, D2, D3, D4, D5, D6, D7, D8,
		E1, E2, E3, E4, E5, E6, E7, E8,
		F1, F2, F3, F4, F5, F6, F7, F8,
		G1, G2, G3, G4, G5, G6, G7, G8,
		H1, H2, H3, H4, H5, H6, H7, H8,
	}
)
