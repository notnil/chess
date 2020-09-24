package chess

import "fmt"

// LongAlgebraicNotation is a more computer friendly alternative to algebraic
// notation.  This notation uses the same format as the UCI (Universal Chess
// Interface).  Examples: e2e4, e7e5, e1g1 (white short castling), e7e8q (for promotion)
type LongAlgebraicNotation struct{}

// String implements the fmt.Stringer interface and returns
// the notation's name.
func (_ LongAlgebraicNotation) String() string {
	return "Long Algebraic Notation"
}

// Encode implements the Encoder interface.
func (_ LongAlgebraicNotation) Encode(pos *Position, m *Move) string {
	return m.S1().String() + m.S2().String() + m.Promo().String()
}

// Decode returns the details for a move based off the provided string and
// board state.
//
// While the algorithm is named "LongAlgebraicNotation", this function access
// both the long algebraic notation and Uiversal Chess Interface notation.
//
// Decode implements the Decoder interface.
func (_ LongAlgebraicNotation) Decode(pos *Position, s string) (*Move, error) {

	s = convertToUCI(s)

	l := len(s)
	err := fmt.Errorf(`chess: failed to decode long algebraic notation text "%s" for position %s`, s, pos.String())

	if l < 4 || l > 5 {
		return nil, fmt.Errorf("%w. Length of move string must be in the range [4,5]. Got '%d'", err, l)
	}

	s1, ok := strToSquareMap[s[0:2]]
	if !ok {
		return nil, fmt.Errorf("%w. First square is not valid. Got '%s'", err, s[0:2])
	}
	s2, ok := strToSquareMap[s[2:4]]
	if !ok {
		return nil, fmt.Errorf("%w. Second square is not valid. Got '%s'", err, s[2:4])
	}

	promo := NoPieceType
	if l == 5 {
		promo = pieceTypeFromChar(s[4:5])
		if promo == NoPieceType {
			return nil, fmt.Errorf("%w. Invalid  Piece Type. Got '%s'", err, s[4:5])
		}
	}

	m := &Move{s1: s1, s2: s2, promo: promo}
	p := pos.Board().Piece(s1)
	if p.Type() == King {
		if (s1 == E1 && s2 == G1) || (s1 == E8 && s2 == G8) {
			m.addTag(KingSideCastle)
		} else if (s1 == E1 && s2 == C1) || (s1 == E8 && s2 == C8) {
			m.addTag(QueenSideCastle)
		}
	} else if p.Type() == Pawn && s2 == pos.enPassantSquare {
		m.addTag(EnPassant)
		m.addTag(Capture)
	}

	c1 := p.Color()
	c2 := pos.Board().Piece(s2).Color()
	if c2 != NoColor && c1 != c2 {
		m.addTag(Capture)
	}

	return m, nil
}
