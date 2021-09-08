package chess

import "testing"

type newSquareTest struct {
	f  File
	r  Rank
	sq Square
}

func TestNewSquare(t *testing.T) {
	testCases := []newSquareTest{
		{FileA, Rank1, A1},
		{FileA, Rank8, A8},
		{FileH, Rank1, H1},
		{FileH, Rank8, H8},
		{FileB, Rank4, B4},
		{FileE, Rank8, E8},
		{FileH, Rank3, H3},
		{FileD, Rank7, D7},
	}

	for _, testCase := range testCases {
		square := NewSquare(testCase.f, testCase.r)
		if square != testCase.sq {
			t.Fatalf("expected %s, got %s", testCase.sq.String(), square.String())
		}
	}
}
