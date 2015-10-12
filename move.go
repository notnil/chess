package chess

import "strings"

type Move struct {
	s1       *Square
	s2       *Square
	promo    PieceType
	state    *GameState
	notation string
}

func (m *Move) S1() *Square {
	return m.s1
}

func (m *Move) S2() *Square {
	return m.s2
}

func (m *Move) Promo() PieceType {
	return m.promo
}

func (m *Move) PreMoveState() *GameState {
	return m.state
}

func (m *Move) String() string {
	if m.notation == "" {
		m.notation = encodeMove(m)
	}
	return m.notation
}

func (m *Move) isValid() bool {
	return m.isValidWithCheck(true)
}

func (m *Move) isValidWithCheck(check bool) bool {
	p := m.piece()
	isValid := filterForPiece(p)(m)
	if !isValid {
		return false
	}
	if !check {
		return true
	}
	inCheck := m.postMoveState().board.inCheck(m.state.turn)
	return !inCheck
}

func (m *Move) isCapture() bool {
	return (m.state.board.isOccupied(m.s2) && m.state.board.piece(m.s2).Color() != m.state.turn) || pawnEnPassantFilter(m)
}

func (m *Move) postMoveState() *GameState {
	moveCount := m.state.moveCount
	if m.state.turn == Black {
		moveCount++
	}
	halfMove := m.state.halfMoveClock
	if m.piece().Type() == Pawn || m.state.board.isOccupied(m.s2) || m.isCastling() {
		halfMove = 0
	} else {
		halfMove++
	}
	return &GameState{
		board:           m.postBoard(),
		turn:            m.state.turn.Other(),
		castleRights:    m.postCastleRights(),
		enPassantSquare: m.postEnPassantSquare(),
		halfMoveClock:   halfMove,
		moveCount:       moveCount,
	}
}

func (m *Move) piece() *Piece {
	return m.state.board.piece(m.s1)
}

func (m *Move) postBoard() Board {
	b := m.state.board.copy()
	if m.isCastling() {
		var rookS1, rookS2 *Square
		switch m.s2 {
		case G1:
			rookS1 = H1
			rookS2 = F1
		case G8:
			rookS1 = H8
			rookS2 = F8
		case C1:
			rookS1 = A1
			rookS2 = D1
		case C8:
			rookS1 = A8
			rookS2 = D8
		}
		b[rookS2] = b[rookS1]
		delete(b, rookS1)
	}
	if m.piece().Type() == Pawn && pawnEnPassantFilter(m) {
		rank := Rank(rankStep(m.state.turn.Other()) + int(m.state.enPassantSquare.rank))
		capSq := getSquare(m.s2.file, rank)
		delete(b, capSq)
	}
	b[m.s2] = b[m.s1]
	delete(b, m.s1)
	if m.promo != NoPiece {
		b[m.s2] = getPiece(m.promo, m.state.turn)
	}
	return b
}

func (m *Move) postEnPassantSquare() *Square {
	p := m.piece()
	if p.Type() == Pawn && pawnUpTwoFilter(m) {
		return m.s1.squaresTo(m.s2)[0]
	}
	return nil
}

func (m *Move) postCastleRights() CastleRights {
	cr := string(m.state.castleRights)
	p := m.piece()
	if p == WhiteKing || m.s1 == H1 {
		cr = strings.Replace(cr, "K", "", -1)
	}
	if p == WhiteKing || m.s1 == A1 {
		cr = strings.Replace(cr, "Q", "", -1)
	}
	if p == BlackKing || m.s1 == H8 {
		cr = strings.Replace(cr, "k", "", -1)
	}
	if p == BlackKing || m.s1 == A8 {
		cr = strings.Replace(cr, "q", "", -1)
	}
	if cr == "" {
		cr = "-"
	}
	return CastleRights(cr)
}

func (m *Move) isCastling() bool {
	p := m.piece()
	if p == nil {
		return false
	}
	backRow := [8]*Square{A1, B1, C1, D1, E1, F1, G1, H1}
	if m.state.turn == Black {
		backRow = [8]*Square{A8, B8, C8, D8, E8, F8, G8, H8}
	}
	kingSide := m.s1 == backRow[4] && m.s2 == backRow[6]
	queenSide := m.s1 == backRow[4] && m.s2 == backRow[2]
	return p.Type() == King && (kingSide || queenSide)
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
		filters = append(filters, promotionFilter, s2Filter)
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
	// filters moves where s1 doesn't have a piece
	s1Filter = func(m *Move) bool {
		return m.state.board.isOccupied(m.s1)
	}

	// filters moves for pieces of the wrong color
	turnFilter = func(m *Move) bool {
		return m.piece().Color() == m.state.turn
	}

	// filters moves into pieces of the same color
	s2Filter = func(m *Move) bool {
		return !m.state.board.isOccupied(m.s2) || m.state.board.piece(m.s2).Color() != m.state.turn
	}

	// filters moves that have a piece in between s1 and s2 (not used for knights)
	blockedFilter = func(m *Move) bool {
		return m.state.board.emptyBetween(m.s1, m.s2)
	}

	// filters moves that have a promotion (not used for pawns)
	promotionFilter = func(m *Move) bool {
		return m.promo == NoPiece
	}

	// filters invalid moves for the king
	kingFilter = func(m *Move) bool {
		kingMove := m.s1.fileDif(m.s2) <= 1 && m.s1.rankDif(m.s2) <= 1
		return kingMove || castleFilter(m)
	}

	castleFilter moveFilter

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
		promotable := m.promo.isPromotable()
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
		capture := abs(int(m.s2.file)-int(m.s1.file)) == 1 && m.state.board.isOccupied(m.s2)
		return upOne && capture
	}

	// filters invalid moves for the pawn using en passant
	pawnEnPassantFilter = func(m *Move) bool {
		p := m.piece()
		rankStep := rankStep(p.Color())
		upOne := int(m.s2.rank) == int(m.s1.rank)+rankStep
		capture := abs(int(m.s2.file)-int(m.s1.file)) == 1 && m.s2 == m.state.enPassantSquare
		return upOne && capture
	}
)

func init() {
	// Must be in init() because of initialization loop
	// filters invalid moves for the king castleing
	castleFilter = func(m *Move) bool {
		if !m.isCastling() {
			return false
		}
		c := m.state.turn
		canCastleKingSide := m.state.castleRights.CanCastle(c, KingSide)
		canCastleQueenSide := m.state.castleRights.CanCastle(c, QueenSide)
		if !(canCastleKingSide || canCastleQueenSide) {
			return false
		}
		backRow := [8]*Square{A1, B1, C1, D1, E1, F1, G1, H1}
		if c == Black {
			backRow = [8]*Square{A8, B8, C8, D8, E8, F8, G8, H8}
		}
		kingSide := m.s1 == backRow[4] && m.s2 == backRow[6]
		queenSide := m.s1 == backRow[4] && m.s2 == backRow[2]
		kingSideOccupied := !m.state.board.emptyBetween(backRow[4], backRow[7])
		kingSideAttacked := m.state.board.isSquareAttacked(m.state.turn, backRow[4], backRow[5], backRow[6])
		if kingSide && (!canCastleKingSide || kingSideOccupied || kingSideAttacked) {
			return false
		}
		queenSideOccupied := !m.state.board.emptyBetween(backRow[0], backRow[4])
		queenSideAttacked := m.state.board.isSquareAttacked(m.state.turn, backRow[2], backRow[3], backRow[4])
		if queenSide && (!canCastleQueenSide || queenSideOccupied || queenSideAttacked) {
			return false
		}
		return true
	}
}
