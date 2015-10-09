package chess

type File int

const (
	NoFile File = iota
	A
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

func fileFromStr(s string) File {
	file, in := strToFileMap[s]
	if !in {
		return NoFile
	}
	return file
}

type Rank int

const (
	NoRank Rank = iota
	R1
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

func rankFromStr(s string) Rank {
	rank, in := strToRankMap[s]
	if !in {
		return NoRank
	}
	return rank
}

type Square struct {
	file File
	rank Rank
}

func (s *Square) String() string {
	return s.file.String() + s.rank.String()
}

func squareFromStr(s string) *Square {
	return strToSquareMap[s]
}

// TODO should use direct access from the array
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

	strToSquareMap = map[string]*Square{
		"a1": A1, "a2": A2, "a3": A3, "a4": A4, "a5": A5, "a6": A6, "a7": A7, "a8": A8,
		"b1": B1, "b2": B2, "b3": B3, "b4": B4, "b5": B5, "b6": B6, "b7": B7, "b8": B8,
		"c1": C1, "c2": C2, "c3": C3, "c4": C4, "c5": C5, "c6": C6, "c7": C7, "c8": C8,
		"d1": D1, "d2": D2, "d3": D3, "d4": D4, "d5": D5, "d6": D6, "d7": D7, "d8": D8,
		"e1": E1, "e2": E2, "e3": E3, "e4": E4, "e5": E5, "e6": E6, "e7": E7, "e8": E8,
		"f1": F1, "f2": F2, "f3": F3, "f4": F4, "f5": F5, "f6": F6, "f7": F7, "f8": F8,
		"g1": G1, "g2": G2, "g3": G3, "g4": G4, "g5": G5, "g6": G6, "g7": G7, "g8": G8,
		"h1": H1, "h2": H2, "h3": H3, "h4": H4, "h5": H5, "h6": H6, "h7": H7, "h8": H8,
	}

	strToFileMap = map[string]File{
		"a": A, "b": B, "c": C, "d": D, "e": E, "f": F, "g": G,
	}

	strToRankMap = map[string]Rank{
		"1": R1, "2": R2, "3": R3, "4": R4, "5": R5, "6": R6, "7": R7, "8": R8,
	}
)
