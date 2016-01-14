package chess

// A MoveTag represents a notable consequence of a move.
type MoveTag int8

const (
	// KingSideCastle indicates that the move is a king side castle.
	KingSideCastle MoveTag = iota + 1
	// QueenSideCastle indicates that the move is a queen side castle.
	QueenSideCastle
	// Capture indicates that the move captures a piece.
	Capture
	// EnPassant indicates that the move captures via en passant.
	EnPassant
	// Check indicates that the move puts the opposing player in check.
	Check
	// inCheck indicates that the move puts the moving player in check and
	// is therefore invalid.
	inCheck
)

// A Move is the moving of a piece from one square to another.
type Move struct {
	s1    Square
	s2    Square
	promo PieceType
	tags  []MoveTag
}

func (m *Move) String() string {
	return m.s1.String() + m.s2.String() + m.promo.String()
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

// HasTag returns true if the move contains the MoveTag given.
func (m *Move) HasTag(tag MoveTag) bool {
	if m.tags == nil {
		return false
	}
	for _, t := range m.tags {
		if t == tag {
			return true
		}
	}
	return false
}
