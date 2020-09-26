package chess

import (
	"fmt"
	"strings"
)

// LongAlgebraicNotation is a more verbose alternative to standard
// Algebraic notation. Where as standard notation might describe a move as
// "Nxc2+", the "Long" form would be "Nd4xc2+" (describing the origin square).
type LongAlgebraicNotation struct{}

// String implements the fmt.Stringer interface and returns
// the notation's name.
func (_ LongAlgebraicNotation) String() string {
	return "Long Algebraic Notation"
}

// Encode describes a given move for a position in Long Algebraic Notation
//
// Encode implements the Encoder interface.
func (_ LongAlgebraicNotation) Encode(pos *Position, m *Move) string {
	var b strings.Builder

	fmt.Fprintf(&b, pos.board.Piece(m.S1()).String())
	fmt.Fprintf(&b, m.S1().String())
	if m.HasTag(Capture) {
		fmt.Fprintf(&b, "x")
	}
	fmt.Fprintf(&b, m.S2().String())
	fmt.Fprintf(&b, m.Promo().String())

	if m.HasTag(Check) {
		fmt.Fprintf(&b, "+")
	} else {
		nextPos := pos.Update(m)
		if nextPos.Status() == Checkmate {
			fmt.Fprintf(&b, "#")
		}
	}

	return b.String()
}

// Decode translates the provided notation described in LongAlgebraicNotation
// into a Move
//
// Decode implements the Decoder interface.
func (_ LongAlgebraicNotation) Decode(pos *Position, s string) (*Move, error) {

	m := &Move{}
	err := fmt.Errorf(`chess: failed to decode Long Algebraic Notation "%s" for position %s`, s, pos.String())

	if len(s) >= 5 && s[0:5] == "O-O-O" {
		if pos.turn == White {
			m.s1 = E1
			m.s2 = C1
		} else {
			m.s1 = E8
			m.s2 = C8
		}
		m.addTag(QueenSideCastle)
		s = s[5:]
	} else if len(s) >= 3 && s[0:3] == "O-O" {
		if pos.turn == White {
			m.s1 = E1
			m.s2 = G1
		} else {
			m.s1 = E8
			m.s2 = G8
		}
		m.addTag(KingSideCastle)
		s = s[3:]
	} else {
		var pt PieceType
		pt, s = LongAlgebraicNotation{}.pieceTypeSubstr(s)

		var ok bool
		m.s1, ok = strToSquareMap[s[0:2]]
		if !ok {
			return nil, fmt.Errorf("%w. First square is not valid. Got '%s'", err, s[0:2])
		}
		s = s[2:]

		var isCap bool
		isCap, s = LongAlgebraicNotation{}.captureSubStr(s)
		if isCap {
			m.addTag(Capture)
		}

		m.s2, ok = strToSquareMap[s[0:2]]
		if !ok {
			return nil, fmt.Errorf("%w. Second square is not valid. Got '%s'", err, s[2:4])
		}
		s = s[2:]

		if pt == Pawn && m.s2 == pos.enPassantSquare {
			m.addTag(EnPassant)
			m.addTag(Capture)
		}

		m.promo, s = LongAlgebraicNotation{}.promotionSubstr(s)
	}

	_, s = LongAlgebraicNotation{}.checkSymbolSubstr(s)

	_, decodeErr := symbolToEvaluation(s)
	if decodeErr != nil {
		return nil, fmt.Errorf("%w. %s", err, decodeErr)
	}

	return m, nil
}

// pieceTypeSubstr returns the type of piece referenced by the notation
//
// The remaining string (without the piece type) is also returned
func (_ LongAlgebraicNotation) pieceTypeSubstr(s string) (PieceType, string) {
	if len(s) < 1 {
		return NoPieceType, s
	}
	switch s[0:1] {
	case "K":
		return King, s[1:]
	case "Q":
		return Queen, s[1:]
	case "R":
		return Rook, s[1:]
	case "N":
		return Knight, s[1:]
	case "B":
		return Bishop, s[1:]
	}
	return Pawn, s
}

// captureSubStr returns whether the notation references a capture
//
// The remaining string (without the capture notation) is also returned
func (_ LongAlgebraicNotation) captureSubStr(s string) (bool, string) {
	if len(s) < 1 {
		return false, s
	}
	switch s[0:1] {
	case "x":
		return true, s[1:]
	case "-":
		return false, s[1:]
	}
	return false, s
}

// promotionSubstr returns the PieceType the move will promote to
//
// The remaining string (without the promotion notation) is also returned
func (_ LongAlgebraicNotation) promotionSubstr(s string) (PieceType, string) {
	if len(s) < 2 || s[0:1] != "=" {
		return NoPieceType, s
	}

	return pieceTypeFromChar(strings.ToLower(s[1:2])), s[2:]
}

// checkSymbolSubstr returns the mark describing whether the move ends in
// check or checkmate
//
// The remaining string (without the result symbol) is also returned
func (_ LongAlgebraicNotation) checkSymbolSubstr(s string) (string, string) {
	if len(s) < 1 {
		return "", s
	}

	symbol := s[0:1]
	if symbol == "#" || symbol == "+" {
		return symbol, s[1:]
	}

	return "", s
}
