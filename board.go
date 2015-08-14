package chess

type Board map[*Square]*Piece

func (b Board) copy() Board {
	n := map[*Square]*Piece{}
	for k, v := range b {
		n[k] = v
	}
	return n
}

func (b Board) occupied(s *Square) bool {
	_, occupied := b[s]
	return occupied
}

func (b Board) piece(s *Square) *Piece {
	p, _ := b[s]
	return p
}

func (b Board) move(m *move) {
	b[m.s2] = b[m.s1]
	delete(b, m.s1)
	if m.promo != nil {
		b[m.s2] = m.promo
	}
	if b.isCastling(m) {
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
}

func (b Board) isCastling(m *move) bool {
	p := b.piece(m.s1)
	if p == nil {
		return false
	}
	turn := b.piece(m.s1).color()
	backRow := [8]*Square{A1, B1, C1, D1, E1, F1, G1, H1}
	if turn == black {
		backRow = [8]*Square{A8, B8, C8, D8, E8, F8, G8, H8}
	}
	kingSide := m.s1 == backRow[4] && m.s2 == backRow[6]
	queenSide := m.s1 == backRow[4] && m.s2 == backRow[2]
	isKing := p != nil && p.pieceType() == king
	return isKing && (kingSide || queenSide)
}

func (b Board) isEnPassant(m *move) bool {
	p := b.piece(m.s1)
	if p == nil {
		return false
	}

	c := p.color()
	fifthRank := rank(int(c.backRank()) + (c.rankStep() * 4))
	sixthRank := rank(int(fifthRank) + c.rankStep())
	capPawnSq := square(m.s2.file, fifthRank)
	capPawn := b.piece(capPawnSq)

	isOnFifth := fifthRank == m.s1.rank
	isPawn := p.pieceType() == pawn
	isCapturing := m.s1.fileDif(m.s2) == 1 && m.s2.rank == sixthRank
	isCapPawn := capPawn != nil && capPawn.pieceType() == pawn && capPawn.color() == c.other()
	return isOnFifth && isPawn && isCapturing && isCapPawn
}

func (b Board) squaresForColor(c color) []*Square {
	squares := []*Square{}
	for _, sq := range allSquares {
		if b.occupied(sq) && b.piece(sq).color() == c {
			squares = append(squares, sq)
		}
	}
	return squares
}

func (b Board) hasValidMoves(c color) bool {
	for _, sq := range b.squaresForColor(c) {
		for _, s2 := range b.validMoves(sq) {
			cp := b.copy()
			m := &move{s1: sq, s2: s2, promo: nil} // TOOD might require a promotion
			cp.move(m)
			if !cp.inCheck(c) {
				return true
			}
		}
	}
	return false
}

func (b Board) inCheck(c color) bool {
	kingSq := b.kingSquare(c)
	// should only happen in unit tests
	if kingSq == nil {
		return false
	}
	return b.isSquareAttacked(c, kingSq)
}

func (b Board) isSquareAttacked(c color, s *Square) bool {
	for _, sq := range b.squaresForColor(c.other()) {
		if squareSlice(b.validMoves(sq)).has(s) {
			return true
		}
	}
	return false
}

func (b Board) inCheckmate(c color) bool {
	// should only happen in unit tests
	if b.kingSquare(c) == nil {
		return false
	}
	return b.inCheck(c) && !b.hasValidMoves(c)
}

func (b Board) inStalemate(c color) bool {
	kingSq := b.kingSquare(c)
	// should only happen in unit tests
	if kingSq == nil {
		return false
	}
	return !b.inCheck(c) && !b.hasValidMoves(c)
}

func (b Board) kingSquare(c color) *Square {
	for _, sq := range b.squaresForColor(c) {
		if b.piece(sq).pieceType() == king {
			return sq
		}
	}
	return nil
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

func (b Board) emptyBetween(s1 *Square, s2 *Square) bool {
	for _, s := range s1.squaresTo(s2) {
		if b.piece(s) != nil {
			return false
		}
	}
	return true
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
		return s1.fileDif(s2) <= 1 && s1.rankDif(s2) <= 1
	}

	queenFilter = func(b Board, s1 *Square, s2 *Square) bool {
		return rookFilter(b, s1, s2) || bishopFilter(b, s1, s2)
	}

	rookFilter = func(b Board, s1 *Square, s2 *Square) bool {
		return (s1.file == s2.file || s1.rank == s2.rank)
	}

	bishopFilter = func(b Board, s1 *Square, s2 *Square) bool {
		return s1.fileDif(s2) == s1.rankDif(s2)
	}

	knightFilter = func(b Board, s1 *Square, s2 *Square) bool {
		fileDif := s1.fileDif(s2)
		rankDif := s1.rankDif(s2)
		return (fileDif == 1 && rankDif == 2) || (fileDif == 2 && rankDif == 1)
	}

	// TODO En passant
	pawnFilter = func(b Board, s1 *Square, s2 *Square) bool {
		p := b.piece(s1)
		rankStep := p.color().rankStep()
		startRank := rank(int(p.color().backRank()) + rankStep)
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
