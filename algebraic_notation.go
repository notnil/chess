package chess

import "fmt"

// AlgebraicNotation (or Standard Algebraic Notation) is the
// official chess notation used by FIDE. Examples: e2, e5,
// O-O (short castling), e8=Q (promotion)
type AlgebraicNotation struct{}

// String implements the fmt.Stringer interface and returns
// the notation's name.
func (_ AlgebraicNotation) String() string {
	return "Algebraic Notation"
}

// Encode implements the Encoder interface.
func (_ AlgebraicNotation) Encode(pos *Position, m *Move) string {
	checkChar := getCheckChar(pos, m)
	if m.HasTag(KingSideCastle) {
		return "O-O" + checkChar
	} else if m.HasTag(QueenSideCastle) {
		return "O-O-O" + checkChar
	}
	p := pos.Board().Piece(m.S1())
	pChar := charFromPieceType(p.Type())
	s1Str := formS1(pos, m)
	capChar := ""
	if m.HasTag(Capture) || m.HasTag(EnPassant) {
		capChar = "x"
		if p.Type() == Pawn && s1Str == "" {
			capChar = m.s1.File().String() + "x"
		}
	}
	promoText := charForPromo(m.promo)
	return pChar + s1Str + capChar + m.s2.String() + promoText + checkChar
}

// Decode implements the Decoder interface.
func (_ AlgebraicNotation) Decode(pos *Position, s string) (*Move, error) {
	s = removeSubstrings(s, "?", "!", "+", "#", "e.p.")
	for _, m := range pos.ValidMoves() {
		str := AlgebraicNotation{}.Encode(pos, m)
		str = removeSubstrings(str, "?", "!", "+", "#", "e.p.")
		if str == s {
			return m, nil
		}
	}
	return nil, fmt.Errorf("chess: could not decode algebraic notation %s for position %s", s, pos.String())
}
