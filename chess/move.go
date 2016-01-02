package chess

// A Move is the moving of a piece from one square to another.
type Move struct {
	s1    Square
	s2    Square
	promo PieceType
}

func (m *Move) String() string {
	return m.s1.String() + m.s2.String()
}

// S1 returns the origin square of the move.
func (m *Move) S1() Square {
	return m.s1
}

// S2 returns the destination square of the move.
func (m *Move) S2() Square {
	return m.s2
}

// Promo returns promotion peice type of the move.
func (m *Move) Promo() PieceType {
	return m.promo
}
