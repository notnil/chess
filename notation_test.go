package chess

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

type validNotationTest struct {
	Pos1        *Position
	Pos2        *Position
	AlgText     string
	LongAlgText string
	Description string
}

func TestValidDecoding(t *testing.T) {
	f, err := os.Open("assets/valid_notation_tests.json")
	if err != nil {
		t.Fatal(err)
		return
	}

	validTests := []validNotationTest{}
	if err := json.NewDecoder(f).Decode(&validTests); err != nil {
		t.Fatal(err)
		return
	}

	for _, test := range validTests {
		t.Run(test.Description, func(t *testing.T) {
			for name, alg := range map[string]Notation{
				"AlgebraicNotation":     AlgebraicNotation{},
				"LongAlgebraicNotation": LongAlgebraicNotation{},
			} {
				t.Run(name, func(t *testing.T) {
					moveText := test.AlgText
					if name == "LongAlgebraicNotation" {
						moveText = test.LongAlgText
					}

					m, err := alg.Decode(test.Pos1, moveText)
					if err != nil {
						movesStrList := []string{}
						for _, m := range test.Pos1.ValidMoves() {
							s := alg.Encode(test.Pos1, m)
							movesStrList = append(movesStrList, s)
						}
						t.Fatalf("starting from board \n%s\n expected move to be valid error\n%s.\nValid Moves: %s\n", test.Pos1.board.Draw(), err, strings.Join(movesStrList, ","))
					}

					postPos := test.Pos1.Update(m)
					if test.Pos2.String() != postPos.String() {
						t.Fatalf("starting from board \n%s\n after move %s\n expected board to be %s\n%s\n but was %s\n%s\n",
							test.Pos1.board.Draw(), m.String(), test.Pos2.String(),
							test.Pos2.board.Draw(), postPos.String(), postPos.board.Draw())
					}
				})
			}
		})
	}
}

type notationDecodeTest struct {
	N       Notation
	Pos     *Position
	Text    string
	PostPos *Position
}

var (
	invalidDecodeTests = []notationDecodeTest{
		{
			// opening for white
			N:    AlgebraicNotation{},
			Pos:  unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			Text: "e5",
		},
		{
			// http://en.lichess.org/W91M4jms#14
			N:    AlgebraicNotation{},
			Pos:  unsafeFEN("rn1qkb1r/pp3ppp/2p1pn2/3p4/2PP4/2NQPN2/PP3PPP/R1B1K2R b KQkq - 0 7"),
			Text: "Nd7",
		},
		{
			// http://en.lichess.org/W91M4jms#17
			N:       AlgebraicNotation{},
			Pos:     unsafeFEN("r2qk2r/pp1n1ppp/2pbpn2/3p4/2PP4/1PNQPN2/P4PPP/R1B1K2R w KQkq - 1 9"),
			Text:    "O-O-O-O",
			PostPos: unsafeFEN("r2qk2r/pp1n1ppp/2pbpn2/3p4/2PP4/1PNQPN2/P4PPP/R1B2RK1 b kq - 0 9"),
		},
		{
			// http://en.lichess.org/W91M4jms#23
			N:    AlgebraicNotation{},
			Pos:  unsafeFEN("3r1rk1/pp1nqppp/2pbpn2/3p4/2PP4/1PNQPN2/PB3PPP/3RR1K1 b - - 5 12"),
			Text: "dx4",
		},
		{
			// should not assume pawn for unknown piece type "n"
			N:    AlgebraicNotation{},
			Pos:  unsafeFEN("rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq e6 0 2"),
			Text: "nf3",
		},
	}
)

func TestInvalidDecoding(t *testing.T) {
	for _, test := range invalidDecodeTests {
		if _, err := test.N.Decode(test.Pos, test.Text); err == nil {
			t.Fatalf("starting from board\n%s\n expected move notation %s to be invalid", test.Pos.board.Draw(), test.Text)
		}
	}
}

func TestConvertToUCI(t *testing.T) {
	tests := []struct {
		name            string
		longAlgNotation string
		expectedUCI     string
	}{
		{
			name:            "Pawn Move (same as UCI)",
			longAlgNotation: "d6c7",
			expectedUCI:     "d6c7",
		},
		{
			name:            "UCI with Promotion",
			longAlgNotation: "d2d1q",
			expectedUCI:     "d2d1q",
		},
		{
			name:            "Named Piece Move, Check",
			longAlgNotation: "Nh4f5+",
			expectedUCI:     "h4f5",
		},
		{
			name:            "Optional Hyphen, Check",
			longAlgNotation: "Nh4-f5+",
			expectedUCI:     "h4f5",
		},
		{
			name:            "Capture, Checkmate",
			longAlgNotation: "Qd4xf6#",
			expectedUCI:     "d4f6",
		},
		{
			name:            "Pawn Promote",
			longAlgNotation: "a7b8=Q",
			expectedUCI:     "a7b8q",
		},
		{
			name:            "Pawn Promote, Checkmate",
			longAlgNotation: "a7b8=Q#",
			expectedUCI:     "a7b8q",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if res := convertToUCI(test.longAlgNotation); res != test.expectedUCI {
				t.Errorf("'%s' not converted to UCI as expected. Got '%s', Expected '%s'", test.longAlgNotation, res, test.expectedUCI)
			}
		})
	}
}
