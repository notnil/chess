package chess

import "testing"

func TestPieceString(t *testing.T) {
	tables := []struct {
		piece PieceType
		str   string
	}{
		{King, "k"},
		{Queen, "q"},
		{Rook, "r"},
		{Bishop, "b"},
		{Knight, "n"},
		{Pawn, "p"},
	}

	for _, table := range tables {
		if table.piece.String() != table.str {
			t.Errorf("String version of piece was incorrect.")
		}
	}
}
