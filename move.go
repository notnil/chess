package chess

type Move struct {
	s1    *Square
	s2    *Square
	promo *PieceType
	b     Board
	state *GameState
}

func (m *Move) isValid() bool {
	return false
}

type moveFilter func(m *Move) bool

func (f moveFilter) and(filters ...moveFilter) moveFilter {
	cp := append([]moveFilter(nil), filters...)
	cp = append(cp, f)
	return func(m *Move) bool {
		for _, f := range cp {
			if !f(m) {
				return false
			}
		}
		return true
	}
}

var (
	// filters squares where s1 doesn't have a piece
	occupiedFilter = func(m *Move) bool {
		return !m.b.isOccupied(m.s1)
	}

	// filters squares that have a piece in between s1 and s2 (not used for knights)
	blockedFilter = func(m *Move) bool {
		squares := m.s1.squaresTo(m.s2)
		for _, sq := range squares {
			if m.b.isOccupied(sq) {
				return false
			}
		}
		return true
	}

	// filters invalid moves for the king w/o regaurd to check or castleing
	kingFilter = func(m *Move) bool {
		return m.s1.fileDif(m.s2) <= 1 && m.s1.rankDif(m.s2) <= 1
	}

	// filters invalid moves for the queen
	queenFilter = func(m *Move) bool {
		return rookFilter(m) || bishopFilter(m)
	}

	// filters invalid moves for the rook
	rookFilter = func(m *Move) bool {
		return (m.s1.file == m.s2.file || m.s1.rank == m.s2.rank)
	}

	// filters invalid moves for the bishop
	bishopFilter = func(m *Move) bool {
		return m.s1.fileDif(m.s2) == m.s1.rankDif(m.s2)
	}

	// filters invalid moves for the knight
	knightFilter = func(m *Move) bool {
		fileDif := m.s1.fileDif(m.s2)
		rankDif := m.s1.rankDif(m.s2)
		return (fileDif == 1 && rankDif == 2) || (fileDif == 2 && rankDif == 1)
	}

	// filters invalid moves for the pawn w/o reguard for promotion and en passant
	pawnFilter = func(m *Move) bool {
		p := m.b.piece(m.s1)
		rankStep := rankStep(p.Color())
		startRank := Rank(int(backRank(p.Color())) + rankStep)
		sameFile := m.s1.file == m.s2.file
		upOne := int(m.s2.rank) == int(m.s1.rank)+rankStep
		upTwoFirstMove := m.s1.rank == startRank && int(m.s2.rank) == int(m.s1.rank)+(2*rankStep)
		capture := upOne && abs(int(m.s2.file)-int(m.s1.file)) == 1 && m.b.isOccupied(m.s2)
		return (upOne && sameFile) || (upTwoFirstMove && sameFile) || capture
	}
)
