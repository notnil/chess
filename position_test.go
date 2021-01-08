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
