package chess

import "testing"

type pgnTest struct {
	PostState *GameState
	PGN       string
}

var (
	validPGNs = []pgnTest{
		{
			PostState: unsafeFEN("8/p3kppp/2p2n2/1p6/2P5/4PP2/PN3KPP/8 w - - 3 24"),
			PGN: `[Event "Troll Masters"]
            [Site "Gausdal NOR"]
            [Date "2001.01.05"]
            [Round "1"]
            [White "Edvardsen,R"]
            [Black "Carlsen,Magnus"]
            [Result "1/2-1/2"]
            [WhiteElo "2055"]
            [BlackElo ""]
            [ECO "D12"]

            1.d4 Nf6 2.Nf3 d5 3.e3 Bf5 4.c4 c6 5.Nc3 e6 6.Bd3 Bxd3 7.Qxd3 Nbd7 8.b3 Bd6
            9.O-O O-O 10.Bb2 Qe7 11.Rad1 Rad8 12.Rfe1 dxc4 13.bxc4 e5 14.dxe5 Nxe5 15.Nxe5 Bxe5
            16.Qe2 Rxd1 17.Rxd1 Rd8 18.Rxd8+ Qxd8 19.Qd1 Qxd1+ 20.Nxd1 Bxb2 21.Nxb2 b5
            22.f3 Kf8 23.Kf2 Ke7  1/2-1/2`,
		},
	}
)

func TestValidPGNs(t *testing.T) {
	for _, test := range validPGNs {
		game, err := decodePGN(test.PGN)
		if err != nil {
			t.Fatalf("recieved unexpected pgn error %s", err.Error())
		}
		if game.GameState().String() != test.PostState.String() {
			t.Fatalf("expected board to be \n%s\n but got \n%s\n",
				test.PostState.board.Draw(),
				game.GameState().board.Draw())
		}
	}
}
