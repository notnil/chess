package chess

import "testing"

type algDecodeTest struct {
	PreState  *GameState
	AlgText   string
	PostState *GameState
}

type algEncodeTest struct {
	Move    *Move
	AlgText string
}

var (
	validAlgDecodeTests = []algDecodeTest{
		{
			// opening for white
			PreState:  unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			AlgText:   "e4",
			PostState: unsafeFEN("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"),
		},
		{
			// http://en.lichess.org/W91M4jms#14
			PreState:  unsafeFEN("rn1qkb1r/pp3ppp/2p1pn2/3p4/2PP4/2NQPN2/PP3PPP/R1B1K2R b KQkq - 0 7"),
			AlgText:   "Nbd7",
			PostState: unsafeFEN("r2qkb1r/pp1n1ppp/2p1pn2/3p4/2PP4/2NQPN2/PP3PPP/R1B1K2R w KQkq - 1 8"),
		},
		{
			// http://en.lichess.org/W91M4jms#17
			PreState:  unsafeFEN("r2qk2r/pp1n1ppp/2pbpn2/3p4/2PP4/1PNQPN2/P4PPP/R1B1K2R w KQkq - 1 9"),
			AlgText:   "O-O",
			PostState: unsafeFEN("r2qk2r/pp1n1ppp/2pbpn2/3p4/2PP4/1PNQPN2/P4PPP/R1B2RK1 b kq - 0 9"),
		},
		{
			// http://en.lichess.org/W91M4jms#23
			PreState:  unsafeFEN("3r1rk1/pp1nqppp/2pbpn2/3p4/2PP4/1PNQPN2/PB3PPP/3RR1K1 b - - 5 12"),
			AlgText:   "dxc4",
			PostState: unsafeFEN("3r1rk1/pp1nqppp/2pbpn2/8/2pP4/1PNQPN2/PB3PPP/3RR1K1 w - - 0 13"),
		},
		{
			// http://en.lichess.org/W91M4jms#34
			PreState:  unsafeFEN("3r2k1/pp2qppp/2p2n2/4b3/2P5/2N1P3/PB2QPPP/3R2K1 w - - 1 18"),
			AlgText:   "Rxd8+",
			PostState: unsafeFEN("3R2k1/pp2qppp/2p2n2/4b3/2P5/2N1P3/PB2QPPP/6K1 b - - 0 18"),
		},
	}

	invalidAlgDecodeTests = []algDecodeTest{
		{
			// opening for white
			PreState: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			AlgText:  "e5",
		},
		{
			// http://en.lichess.org/W91M4jms#14
			PreState: unsafeFEN("rn1qkb1r/pp3ppp/2p1pn2/3p4/2PP4/2NQPN2/PP3PPP/R1B1K2R b KQkq - 0 7"),
			AlgText:  "Nd7",
		},
		{
			// http://en.lichess.org/W91M4jms#17
			PreState:  unsafeFEN("r2qk2r/pp1n1ppp/2pbpn2/3p4/2PP4/1PNQPN2/P4PPP/R1B1K2R w KQkq - 1 9"),
			AlgText:   "O-O-O-O",
			PostState: unsafeFEN("r2qk2r/pp1n1ppp/2pbpn2/3p4/2PP4/1PNQPN2/P4PPP/R1B2RK1 b kq - 0 9"),
		},
		{
			// http://en.lichess.org/W91M4jms#23
			PreState: unsafeFEN("3r1rk1/pp1nqppp/2pbpn2/3p4/2PP4/1PNQPN2/PB3PPP/3RR1K1 b - - 5 12"),
			AlgText:  "dx4",
		},
	}

	algEncodeTests = []algEncodeTest{
		{
			Move:    &Move{s1: E2, s2: E4, state: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
			AlgText: "e4",
		},
		{
			Move:    &Move{s1: D1, s2: D3, state: unsafeFEN("rn1qkb1r/pp3ppp/2p1pn2/3p4/2PP4/2NbPN2/PP3PPP/R1BQK2R w KQkq - 0 7")},
			AlgText: "Qxd3",
		},
		{
			Move:    &Move{s1: B8, s2: D7, state: unsafeFEN("rn1qkb1r/pp3ppp/2p1pn2/3p4/2PP4/2NQPN2/PP3PPP/R1B1K2R b KQkq - 0 7")},
			AlgText: "Nbd7",
		},
		{
			Move:    &Move{s1: E1, s2: G1, state: unsafeFEN("r2qk2r/pp1n1ppp/2pbpn2/3p4/2PP4/1PNQPN2/P4PPP/R1B1K2R w KQkq - 1 9")},
			AlgText: "O-O",
		},
		{
			Move:    &Move{s1: A1, s2: D1, state: unsafeFEN("r4rk1/pp1nqppp/2pbpn2/3p4/2PP4/1PNQPN2/PB3PPP/R4RK1 w - - 2 11")},
			AlgText: "Rad1",
		},
		{
			Move:    &Move{s1: D8, s2: D1, state: unsafeFEN("3q2k1/pp3ppp/2p2n2/4b3/2P5/2N1P3/PB3PPP/3Q2K1 b - - 1 19")},
			AlgText: "Qxd1+",
		},
	}
)

func TestValidAlgDecoding(t *testing.T) {
	for _, test := range validAlgDecodeTests {
		move, err := decodeMove(test.PreState, test.AlgText)
		if err != nil || !move.isValid() {
			t.Fatalf("starting from board\n%s\n expected move notation %s to be valid", test.PreState.board.Draw(), test.AlgText)
		}
		postState := move.postMoveState()
		if postState.String() != test.PostState.String() {
			t.Fatalf("starting from board \n%s\n after move %s\n expected board to be %s\n%s\n but was %s\n%s\n",
				move.state.board.Draw(), move.String(), test.PostState.String(),
				test.PostState.board.Draw(), postState.String(), postState.board.Draw())
		}
	}
}

func TestInvalidAlgDecoding(t *testing.T) {
	for _, test := range invalidAlgDecodeTests {
		if _, err := decodeMove(test.PreState, test.AlgText); err == nil {
			t.Fatalf("starting from board\n%s\n expected move notation %s to be invalid", test.PreState.board.Draw(), test.AlgText)
		}
	}
}

func TestValidAlgEncoding(t *testing.T) {
	for _, test := range algEncodeTests {
		if !test.Move.isValid() {
			t.Fatalf("starting from board\n%s\n invalid move %s", test.Move.state.board.Draw(), test.Move.String())
		}
		actual := encodeMove(test.Move)
		if actual != test.AlgText {
			t.Fatalf("starting from board\n%s\n to board \n%s\nexpected move notation %s but got %s", test.Move.state.board.Draw(), test.Move.postMoveState().board.Draw(), test.AlgText, actual)
		}
	}
}
