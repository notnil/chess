package chess

import "testing"

type bitboardTestPair struct {
	initial  uint64
	reversed uint64
}

var (
	tests = []bitboardTestPair{
		{
			uint64(1),
			uint64(9223372036854775808),
		},
		{
			uint64(18446744073709551615),
			uint64(18446744073709551615),
		},
		{
			uint64(0),
			uint64(0),
		},
	}
)

func TestBitboardReverse(t *testing.T) {
	for _, p := range tests {
		r := uint64(bitboard(p.initial).Reverse())
		if r != p.reversed {
			t.Fatalf("bitboard reverse of %s expected %s but got %s", intStr(p.initial), intStr(p.reversed), intStr(r))
		}
	}
}

func TestBitboardOccupied(t *testing.T) {
	m := map[Square]bool{
		B3: true,
	}
	bb := newBitboard(m)
	if bb.Occupied(B3) != true {
		t.Fatalf("bitboard occupied of %s expected %t but got %t", bb, true, false)
	}
}

func BenchmarkBitboardReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u := uint64(9223372036854775807)
		bitboard(u).Reverse()
	}
}

func intStr(i uint64) string {
	return bitboard(i).String()
}
