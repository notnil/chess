package chess

import (
	"testing"
)

func TestPositionBinary(t *testing.T) {
	for _, fen := range validFENs {
		pos, err := decodeFEN(fen)
		if err != nil {
			t.Fatal(err)
		}
		b, err := pos.MarshalBinary()
		if err != nil {
			t.Fatal(err)
		}
		cp := &Position{}
		if err := cp.UnmarshalBinary(b); err != nil {
			t.Fatal(err)
		}
		if pos.String() != cp.String() {
			t.Fatalf("expected %s but got %s", pos.String(), cp.String())
		}
	}
}

func TestPositionUpdate(t *testing.T) {
	for _, fen := range validFENs {
		pos, err := decodeFEN(fen)
		if err != nil {
			t.Fatal(err)
		}

		{
			np := pos.Update(pos.ValidMoves()[0])
			if pos.Turn().Other() != np.turn {
				t.Fatal("expected other turn")
			}
			if pos.halfMoveClock+1 != np.halfMoveClock {
				t.Fatal("expected half move clock increment")
			}
			if pos.board.String() == np.board.String() {
				t.Fatal("expected board update")
			}
		}

		{
			np := pos.Update(nil)
			if pos.Turn().Other() != np.turn {
				t.Fatal("expected other turn")
			}
			if pos.halfMoveClock+1 != np.halfMoveClock {
				t.Fatal("expected half move clock increment")
			}
			if pos.board.String() != np.board.String() {
				t.Fatal("expected same board")
			}
		}
	}
}
