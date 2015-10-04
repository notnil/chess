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

func (m *Move) PostMoveState() *GameState {
	// TODO account for castleing
	b := m.state.Board.copy()
	b[m.s2] = b[m.s1]
	delete(b, m.s1)
	if m.promo != nil {
		b[m.s2] = getPiece(*m.promo, m.state.Turn)
	}

	moveCount := m.state.MoveCount
	if m.state.Turn == Black {
		moveCount++
	}
	return &GameState{
		Board:           b,
		Turn:            m.state.Turn.Other(),
		CastleRights:    m.state.CastleRights, // TODO
		EnPassantSquare: nil,                  // TODO
		HalfMoveClock:   0,                    // TODO
		MoveCount:       moveCount,
	}
}

func (m *Move) String() string {
	return fmt.Sprintf("{s1:%s s2:%s FEN:%s}", m.s1, m.s2, m.state)
}

func (m *Move) isValid() bool {
	p := m.piece()
	isValid := filterForPiece(p)(m)
	inCheck := m.PostMoveState().Board.inCheck(m.state.Turn)
	return isValid && !inCheck
}

func (m *Move) piece() *Piece {
	return m.state.Board.piece(m.s1)
}

func filterForPiece(p *Piece) moveFilter {
	if p == nil {
		return s1Filter
	}
	filters := []moveFilter{s1Filter, turnFilter}
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
	if p.Type() != Knight {
		filters = append(filters, blockedFilter)
	}
	if p.Type() != Pawn {
		filters = append(filters, s2Filter)
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
	s1Filter = func(m *Move) bool {
		return m.state.Board.isOccupied(m.s1)
	}

	// filters squares that have pieces that can move on the game's turn
	turnFilter = func(m *Move) bool {
		return m.piece().Color() == m.state.Turn
	}

	// filters squares where s2 isn't empty or the other color's piece (not used for pawns)
	s2Filter = func(m *Move) bool {
		return !m.state.Board.isOccupied(m.s2) || m.state.Board.piece(m.s2).Color() != m.state.Turn
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

	// filters invalid moves for the king castleing
	// castleFilter = func(m *Move) bool {
	// 	return m.s1.fileDif(m.s2) <= 1 && m.s1.rankDif(m.s2) <= 1
	// }

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

	// filters invalid moves for the pawn
	pawnFilter = func(m *Move) bool {
		return pawnUpOneFilter(m) || pawnUpTwoFilter(m) || pawnCaptureFilter(m) || pawnEnPassantFilter(m)
	}

	// filters invalid moves for the pawn moving one square
	pawnUpOneFilter = func(m *Move) bool {
		p := m.piece()
		rankStep := rankStep(p.Color())
		sameFile := m.s1.file == m.s2.file
		upOne := int(m.s2.rank) == int(m.s1.rank)+rankStep
		requiresPromo := m.s2.rank == backRank(p.Color().Other())
		promotable := m.promo != nil && m.promo.isPromotable()
		return (upOne && sameFile && !requiresPromo) || (upOne && sameFile && requiresPromo && promotable)
	}

	// filters invalid moves for the pawn moving two squares
	pawnUpTwoFilter = func(m *Move) bool {
		p := m.piece()
		rankStep := rankStep(p.Color())
		startRank := Rank(int(backRank(p.Color())) + rankStep)
		sameFile := m.s1.file == m.s2.file
		upTwoFirstMove := m.s1.rank == startRank && int(m.s2.rank) == int(m.s1.rank)+(2*rankStep)
		return upTwoFirstMove && sameFile
	}

	// filters invalid moves for the pawn while capturing
	pawnCaptureFilter = func(m *Move) bool {
		p := m.piece()
		rankStep := rankStep(p.Color())
		upOne := int(m.s2.rank) == int(m.s1.rank)+rankStep
		capture := abs(int(m.s2.file)-int(m.s1.file)) == 1 && m.state.Board.isOccupied(m.s2)
		return upOne && capture
	}

	// filters invalid moves for the pawn using en passant
	pawnEnPassantFilter = func(m *Move) bool {
		sq := m.state.EnPassantSquare
		return sq != nil && m.s2 == sq
	}
)
