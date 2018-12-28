package chess

import "testing"

type bitboardTestPair struct {
	initial  uint64
	reversed uint64
}

var (
	tests = []bitboardTestPair{
		bitboardTestPair{
			uint64(1),
			uint64(9223372036854775808),
		},
		bitboardTestPair{
			uint64(18446744073709551615),
			uint64(18446744073709551615),
		},
		bitboardTestPair{
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

func BenchmarkBitboardReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u := uint64(9223372036854775807)
		bitboard(u).Reverse()
	}
}

func TestBitboardOccupied(t *testing.T) {
	b := newBitboard(map[Square]bool{A2: true, H8: true})
	if b.Occupied(A1) {
		t.Error("A1 should not be occupied")
	}
	if !b.Occupied(A2) {
		t.Error("A2 should be occupied")
	}
	if b.Occupied(D4) {
		t.Error("A1 should not be occupied")
	}
	if !b.Occupied(H8) {
		t.Error("A2 should be occupied")
	}
}

func BenchmarkBitboardOccupied(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u := uint64(1)
		bitboard(u).Occupied(A1)
	}
}

func intStr(i uint64) string {
	return bitboard(i).String()
}
