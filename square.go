package chess

type file int

const (
	fileA file = iota
	fileB
	fileC
	fileD
	fileE
	fileF
	fileG
	fileH
)

func (f file) String() string {
	s := "abcdefgh"
	i := int(f)
	return s[i : i+1]
}

func files() []file {
	return []file{fileA, fileB, fileC, fileD, fileE, fileF, fileG, fileH}
}

func fileFromString(s string) *file {
	for _, f := range files() {
		if f.String() == s {
			return &f
		}
	}
	return nil
}

type rank int

const (
	rank1 rank = iota
	rank2
	rank3
	rank4
	rank5
	rank6
	rank7
	rank8
)

func (r rank) String() string {
	s := "12345678"
	i := int(r)
	return s[i : i+1]
}

func ranks() []rank {
	return []rank{rank1, rank2, rank3, rank4, rank5, rank6, rank7, rank8}
}

func rankFromString(s string) *rank {
	for _, r := range ranks() {
		if r.String() == s {
			return &r
		}
	}
	return nil
}

type Square struct {
	file file
	rank rank
}

var (
	A1 = &Square{fileA, rank1}
	A2 = &Square{fileA, rank2}
	A3 = &Square{fileA, rank3}
	A4 = &Square{fileA, rank4}
	A5 = &Square{fileA, rank5}
	A6 = &Square{fileA, rank6}
	A7 = &Square{fileA, rank7}
	A8 = &Square{fileA, rank8}

	B1 = &Square{fileB, rank1}
	B2 = &Square{fileB, rank2}
	B3 = &Square{fileB, rank3}
	B4 = &Square{fileB, rank4}
	B5 = &Square{fileB, rank5}
	B6 = &Square{fileB, rank6}
	B7 = &Square{fileB, rank7}
	B8 = &Square{fileB, rank8}

	C1 = &Square{fileC, rank1}
	C2 = &Square{fileC, rank2}
	C3 = &Square{fileC, rank3}
	C4 = &Square{fileC, rank4}
	C5 = &Square{fileC, rank5}
	C6 = &Square{fileC, rank6}
	C7 = &Square{fileC, rank7}
	C8 = &Square{fileC, rank8}

	D1 = &Square{fileD, rank1}
	D2 = &Square{fileD, rank2}
	D3 = &Square{fileD, rank3}
	D4 = &Square{fileD, rank4}
	D5 = &Square{fileD, rank5}
	D6 = &Square{fileD, rank6}
	D7 = &Square{fileD, rank7}
	D8 = &Square{fileD, rank8}

	E1 = &Square{fileE, rank1}
	E2 = &Square{fileE, rank2}
	E3 = &Square{fileE, rank3}
	E4 = &Square{fileE, rank4}
	E5 = &Square{fileE, rank5}
	E6 = &Square{fileE, rank6}
	E7 = &Square{fileE, rank7}
	E8 = &Square{fileE, rank8}

	F1 = &Square{fileF, rank1}
	F2 = &Square{fileF, rank2}
	F3 = &Square{fileF, rank3}
	F4 = &Square{fileF, rank4}
	F5 = &Square{fileF, rank5}
	F6 = &Square{fileF, rank6}
	F7 = &Square{fileF, rank7}
	F8 = &Square{fileF, rank8}

	G1 = &Square{fileG, rank1}
	G2 = &Square{fileG, rank2}
	G3 = &Square{fileG, rank3}
	G4 = &Square{fileG, rank4}
	G5 = &Square{fileG, rank5}
	G6 = &Square{fileG, rank6}
	G7 = &Square{fileG, rank7}
	G8 = &Square{fileG, rank8}

	H1 = &Square{fileH, rank1}
	H2 = &Square{fileH, rank2}
	H3 = &Square{fileH, rank3}
	H4 = &Square{fileH, rank4}
	H5 = &Square{fileH, rank5}
	H6 = &Square{fileH, rank6}
	H7 = &Square{fileH, rank7}
	H8 = &Square{fileH, rank8}
)

func (s *Square) String() string {
	return s.file.String() + s.rank.String()
}

func square(f file, r rank) *Square {
	for _, s := range allSquares {
		if s.file == f && s.rank == r {
			return s
		}
	}
	return nil
}

func squareFromString(s string) *Square {
	for _, sq := range allSquares {
		if sq.String() == s {
			return sq
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
		sq := square(file(f), rank(r))
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

var allSquares = [64]*Square{
	A1, A2, A3, A4, A5, A6, A7, A8,
	B1, B2, B3, B4, B5, B6, B7, B8,
	C1, C2, C3, C4, C5, C6, C7, C8,
	D1, D2, D3, D4, D5, D6, D7, D8,
	E1, E2, E3, E4, E5, E6, E7, E8,
	F1, F2, F3, F4, F5, F6, F7, F8,
	G1, G2, G3, G4, G5, G6, G7, G8,
	H1, H2, H3, H4, H5, H6, H7, H8,
}

func sign(i int) int {
	if i > 0 {
		return 1
	} else if i < 0 {
		return -1
	}
	return 0
}

func min(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}
