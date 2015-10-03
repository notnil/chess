package chess

import "fmt"

type Move struct {
	s1    *Square
	s2    *Square
	promo *PieceType
	state *GameState
}

func (m *Move) S1() *Square {
	return m.s1
}

func (m *Move) S2() *Square {
	return m.s2
}

func (m *Move) Promo() *PieceType {
	return m.promo
}

func (m *Move) PreMoveState() *GameState {
	return m.state
}

func (m *Move) String() string {
	return fmt.Sprintf("{s1:%s s2:%s FEN:%s}", m.s1, m.s2, m.state)
}

func (m *Move) isValid() bool {
	p := m.piece()
	return filterForPiece(p)(m)
}

func (m *Move) piece() *Piece {
	return m.state.Board.piece(m.s1)
}

func filterForPiece(p *Piece) moveFilter {
	filters := []moveFilter{occupiedFilter}
	switch p.Type() {
	case King:
		filters = append(filters, kingFilter)
	case Queen:
		filters = append(filters, queenFilter)
	case Rook:
		filters = append(filters, rookFilter)
	case Bishop:
		filters = append(filters, bishopFilter)
	case Knight:
		filters = append(filters, knightFilter)
	case Pawn:
		filters = append(filters, pawnFilter)
	}
	if p != nil && p.Type() != Knight {
		filters = append(filters, blockedFilter)
	}
	return moveFilters(filters).chainAnd()
}

type moveFilter func(m *Move) bool

type moveFilters []moveFilter

func (a moveFilters) chainAnd() moveFilter {
	return func(m *Move) bool {
		for _, f := range a {
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
		return m.state.Board.isOccupied(m.s1)
	}

	// filters squares that have pieces that can move on the game's turn
	turnFilter = func(m *Move) bool {
		return m.piece().Color() == m.state.Turn
	}

	// filters squares that have a piece in between s1 and s2 (not used for knights)
	blockedFilter = func(m *Move) bool {
		squares := m.s1.squaresTo(m.s2)
		for _, sq := range squares {
			if m.state.Board.isOccupied(sq) {
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
		p := m.piece()
		rankStep := rankStep(p.Color())
		startRank := Rank(int(backRank(p.Color())) + rankStep)
		sameFile := m.s1.file == m.s2.file
		upOne := int(m.s2.rank) == int(m.s1.rank)+rankStep
		if upOne && sameFile {
			return true
		}
		upTwoFirstMove := m.s1.rank == startRank && int(m.s2.rank) == int(m.s1.rank)+(2*rankStep)
		if upTwoFirstMove && sameFile {
			return true
		}
		capture := upOne && abs(int(m.s2.file)-int(m.s1.file)) == 1 && m.state.Board.isOccupied(m.s2)
		return capture
	}
)
