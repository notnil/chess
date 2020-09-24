package chess

import (
	"fmt"
	"regexp"
	"strings"
)

// Encoder is the interface implemented by objects that can
// encode a move into a string given the position.  It is not
// the encoders responsibility to validate the move.
type Encoder interface {
	Encode(pos *Position, m *Move) string
}

// Decoder is the interface implemented by objects that can
// decode a string into a move given the position. It is not
// the decoders responsibility to validate the move.  An error
// is returned if the string could not be decoded.
type Decoder interface {
	Decode(pos *Position, s string) (*Move, error)
}

// Notation is the interface implemented by objects that can
// encode and decode moves.
type Notation interface {
	Encoder
	Decoder
}

// convertToUCI Returns equivalent LongAlgebraicNotation in Universal Chess
// Interface format.
//
// If the provided string is not in Long Algebraic Notation, return the
// string unmodified.
func convertToUCI(s string) string {

	// Identify only the squares and promotion if present in the string
	re := regexp.MustCompile("[KQBNR]?([a-h][1-8])[-x]?([a-h][1-8])(=?([QqBbNnRr]))?[+#]?")
	matches := re.FindStringSubmatch(s)

	// Two squares were not found
	if len(matches) == 0 || matches[1] == "" || matches[2] == "" {
		return s
	}

	return fmt.Sprintf("%s%s%s",
		matches[1],
		matches[2],
		strings.ToLower(matches[4]),
	)
}

func getCheckChar(pos *Position, move *Move) string {
	if !move.HasTag(Check) {
		return ""
	}
	nextPos := pos.Update(move)
	if nextPos.Status() == Checkmate {
		return "#"
	}
	return "+"
}

func formS1(pos *Position, m *Move) string {
	p := pos.board.Piece(m.s1)
	if p.Type() == Pawn {
		return ""
	}

	var req, fileReq, rankReq bool
	moves := pos.ValidMoves()

	for _, mv := range moves {
		if mv.s1 != m.s1 && mv.s2 == m.s2 && p == pos.board.Piece(mv.s1) {
			req = true

			if mv.s1.File() == m.s1.File() {
				rankReq = true
			}

			if mv.s1.Rank() == m.s1.Rank() {
				fileReq = true
			}
		}
	}

	var s1 = ""

	if fileReq || !rankReq && req {
		s1 = m.s1.File().String()
	}

	if rankReq {
		s1 += m.s1.Rank().String()
	}

	return s1
}

func charForPromo(p PieceType) string {
	c := charFromPieceType(p)
	if c != "" {
		c = "=" + c
	}
	return c
}

func charFromPieceType(p PieceType) string {
	switch p {
	case King:
		return "K"
	case Queen:
		return "Q"
	case Rook:
		return "R"
	case Bishop:
		return "B"
	case Knight:
		return "N"
	}
	return ""
}

func pieceTypeFromChar(c string) PieceType {
	switch c {
	case "q":
		return Queen
	case "r":
		return Rook
	case "b":
		return Bishop
	case "n":
		return Knight
	}
	return NoPieceType
}

func removeSubstrings(s string, subs ...string) string {
	for _, sub := range subs {
		s = strings.Replace(s, sub, "", -1)
	}
	return s
}
