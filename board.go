package chess

type Board map[*Square]Piece

func (b Board) occupied(s *Square) bool {
	_, occupied := b[s]
	return occupied
}

func (b Board) piece(s *Square) Piece {
	p, _ := b[s]
	return p
}

func (b Board) move(m *move) {
	b[m.s2] = b[m.s1]
	delete(b, m.s1)
}

func (b Board) validMoves(s *Square) []*Square {
	if !b.occupied(s) {
		return []*Square{}
	}
	p := b.piece(s)
	f := filter(p.pieceType())
	validSquares := []*Square{}
	for _, sq := range allSquares {
		valid := f(b, s, sq)
		// log.Println(sq, "is", valid)
		if valid {
			validSquares = append(validSquares, sq)
		}
	}
	return validSquares
}

type squareFilter func(b Board, s1 *Square, s2 *Square) bool

type squareFilters []squareFilter

func (a squareFilters) chain() squareFilter {
	return func(b Board, s1 *Square, s2 *Square) bool {
		for _, f := range a {
			if !f(b, s1, s2) {
				return false
			}
		}
		return true
	}
}

func filter(p pieceType) squareFilter {
	filters := []squareFilter{occupiedFilter}
	var f squareFilter
	switch p {
	case king:
		f = kingFilter
	case queen:
		f = queenFilter
	case rook:
		f = rookFilter
	case bishop:
		f = bishopFilter
	case knight:
		f = knightFilter
	case pawn:
		f = pawnFilter
	}
	filters = append(filters, f)
	if p != knight && p != king {
		filters = append(filters, pathBlockedFilter)
	}
	return squareFilters(filters).chain()
}

var (
	occupiedFilter = func(b Board, s1 *Square, s2 *Square) bool {
		if !b.occupied(s2) {
			return true
		}
		return b.piece(s1).color() != b.piece(s2).color()
	}

	pathBlockedFilter = func(b Board, s1 *Square, s2 *Square) bool {
		squares := s1.squaresTo(s2)
		for _, sq := range squares {
			if b.occupied(sq) {
				return false
			}
		}
		return true
	}

	kingFilter = func(b Board, s1 *Square, s2 *Square) bool {
		fileDif := abs(int(s2.file) - int(s1.file))
		rankDif := abs(int(s2.rank) - int(s1.rank))
		return fileDif <= 1 && rankDif <= 1
	}

	queenFilter = func(b Board, s1 *Square, s2 *Square) bool {
		return rookFilter(b, s1, s2) || bishopFilter(b, s1, s2)
	}

	rookFilter = func(b Board, s1 *Square, s2 *Square) bool {
		return (s1.file == s2.file || s1.rank == s2.rank)
	}

	bishopFilter = func(b Board, s1 *Square, s2 *Square) bool {
		fileDif := abs(int(s2.file) - int(s1.file))
		rankDif := abs(int(s2.rank) - int(s1.rank))
		return fileDif == rankDif
	}

	knightFilter = func(b Board, s1 *Square, s2 *Square) bool {
		fileDif := abs(int(s2.file) - int(s1.file))
		rankDif := abs(int(s2.rank) - int(s1.rank))
		return (fileDif == 1 && rankDif == 2) || (fileDif == 2 && rankDif == 1)
	}

	// TODO En passant
	pawnFilter = func(b Board, s1 *Square, s2 *Square) bool {
		p := b.piece(s1)
		startRank := rank2
		rankStep := 1
		if p.color() == black {
			startRank = rank7
			rankStep = -1
		}
		sameFile := s1.file == s2.file
		upOne := int(s2.rank) == int(s1.rank)+rankStep
		upTwoFirstMove := s1.rank == startRank && int(s2.rank) == int(s1.rank)+(2*rankStep)
		capture := upOne && abs(int(s2.file)-int(s1.file)) == 1 && b.occupied(s2)
		return (upOne && sameFile) || (upTwoFirstMove && sameFile) || capture
	}
)

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
