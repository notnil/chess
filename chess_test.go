package chess

import "testing"

type testStruct struct {
	Board Board
	Moves []*move
	Err   bool
}

func TestMoves(t *testing.T) {
	for _, tStruct := range tests {
		for _, m := range tStruct.Moves {
			boardCopy := map[*Square]Piece{}
			for sq, p := range tStruct.Board {
				boardCopy[sq] = p
			}
			g := &Game{
				moves: []*move{},
				board: boardCopy,
			}
			p := g.board.piece(m.s1)
			err := g.Move(m.s1, m.s2, nil)
			if tStruct.Err && err == nil {
				t.Errorf("Expected error from piece %s move %+v", p, m)
			} else if !tStruct.Err && err != nil {
				t.Errorf("Received unexpected error %s from piece %s move %s", err.Error(), p, m)
			}
		}
	}
}

func allMovesExcluding(s1 *Square, moves []*move) []*move {
	excluded := []*move{}
	for _, s := range allSquares {
		found := false
		for _, m := range moves {
			if m.s2 == s {
				found = true
				break
			}
		}
		if !found {
			excluded = append(excluded, &move{
				s1: s1, s2: s, promo: nil,
			})
		}
	}
	return excluded
}

var (
	validKingE4Moves = []*move{
		&move{s1: E4, s2: D3, promo: nil},
		&move{s1: E4, s2: D4, promo: nil},
		&move{s1: E4, s2: D5, promo: nil},
		&move{s1: E4, s2: E3, promo: nil},
		&move{s1: E4, s2: E5, promo: nil},
		&move{s1: E4, s2: F3, promo: nil},
		&move{s1: E4, s2: F4, promo: nil},
		&move{s1: E4, s2: F5, promo: nil},
	}
	validKingE4Test = &testStruct{
		Board: map[*Square]Piece{E4: WKing},
		Moves: validKingE4Moves,
		Err:   false,
	}
	invalidKingE4Test = &testStruct{
		Board: map[*Square]Piece{E4: WKing},
		Moves: allMovesExcluding(E4, validKingE4Moves),
		Err:   true,
	}
	validKingA1Moves = []*move{
		&move{s1: A1, s2: A2, promo: nil},
		&move{s1: A1, s2: B1, promo: nil},
		&move{s1: A1, s2: B2, promo: nil},
	}
	validKingA1Test = &testStruct{
		Board: map[*Square]Piece{A1: WKing},
		Moves: validKingA1Moves,
		Err:   false,
	}
	invalidKingA1Test = &testStruct{
		Board: map[*Square]Piece{A1: WKing},
		Moves: allMovesExcluding(A1, validKingA1Moves),
		Err:   true,
	}
	validQueenC5Moves = []*move{
		&move{s1: C5, s2: A5, promo: nil},
		&move{s1: C5, s2: B5, promo: nil},
		&move{s1: C5, s2: D5, promo: nil},
		&move{s1: C5, s2: E5, promo: nil},
		&move{s1: C5, s2: F5, promo: nil},
		&move{s1: C5, s2: G5, promo: nil},
		&move{s1: C5, s2: H5, promo: nil},
		&move{s1: C5, s2: C1, promo: nil},
		&move{s1: C5, s2: C2, promo: nil},
		&move{s1: C5, s2: C3, promo: nil},
		&move{s1: C5, s2: C4, promo: nil},
		&move{s1: C5, s2: C6, promo: nil},
		&move{s1: C5, s2: C7, promo: nil},
		&move{s1: C5, s2: C8, promo: nil},
		&move{s1: C5, s2: A7, promo: nil},
		&move{s1: C5, s2: B6, promo: nil},
		&move{s1: C5, s2: D4, promo: nil},
		&move{s1: C5, s2: E3, promo: nil},
		&move{s1: C5, s2: F2, promo: nil},
		&move{s1: C5, s2: G1, promo: nil},
		&move{s1: C5, s2: A3, promo: nil},
		&move{s1: C5, s2: B4, promo: nil},
		&move{s1: C5, s2: D6, promo: nil},
		&move{s1: C5, s2: E7, promo: nil},
		&move{s1: C5, s2: F8, promo: nil},
	}
	validQueenC5Test = &testStruct{
		Board: map[*Square]Piece{C5: WQueen},
		Moves: validQueenC5Moves,
		Err:   false,
	}
	invalidQueenC5Test = &testStruct{
		Board: map[*Square]Piece{C5: WQueen},
		Moves: allMovesExcluding(C5, validQueenC5Moves),
		Err:   true,
	}
	validRookC5Moves = []*move{
		&move{s1: C5, s2: A5, promo: nil},
		&move{s1: C5, s2: B5, promo: nil},
		&move{s1: C5, s2: D5, promo: nil},
		&move{s1: C5, s2: E5, promo: nil},
		&move{s1: C5, s2: F5, promo: nil},
		&move{s1: C5, s2: G5, promo: nil},
		&move{s1: C5, s2: H5, promo: nil},
		&move{s1: C5, s2: C1, promo: nil},
		&move{s1: C5, s2: C2, promo: nil},
		&move{s1: C5, s2: C3, promo: nil},
		&move{s1: C5, s2: C4, promo: nil},
		&move{s1: C5, s2: C6, promo: nil},
		&move{s1: C5, s2: C7, promo: nil},
		&move{s1: C5, s2: C8, promo: nil},
	}
	validRookC5Test = &testStruct{
		Board: map[*Square]Piece{C5: WRook},
		Moves: validRookC5Moves,
		Err:   false,
	}
	invalidRookC5Test = &testStruct{
		Board: map[*Square]Piece{C5: WRook},
		Moves: allMovesExcluding(C5, validRookC5Moves),
		Err:   true,
	}
	validBishopG2Moves = []*move{
		&move{s1: G2, s2: A8, promo: nil},
		&move{s1: G2, s2: B7, promo: nil},
		&move{s1: G2, s2: C6, promo: nil},
		&move{s1: G2, s2: D5, promo: nil},
		&move{s1: G2, s2: E4, promo: nil},
		&move{s1: G2, s2: F3, promo: nil},
		&move{s1: G2, s2: H1, promo: nil},
		&move{s1: G2, s2: F1, promo: nil},
		&move{s1: G2, s2: H3, promo: nil},
	}
	validBishopG2Test = &testStruct{
		Board: map[*Square]Piece{G2: WBishop},
		Moves: validBishopG2Moves,
		Err:   false,
	}
	invalidBishopG2Test = &testStruct{
		Board: map[*Square]Piece{G2: WBishop},
		Moves: allMovesExcluding(G2, validBishopG2Moves),
		Err:   true,
	}
	validKnightD4Moves = []*move{
		&move{s1: D4, s2: B3, promo: nil},
		&move{s1: D4, s2: B5, promo: nil},
		&move{s1: D4, s2: C2, promo: nil},
		&move{s1: D4, s2: C6, promo: nil},
		&move{s1: D4, s2: E2, promo: nil},
		&move{s1: D4, s2: E6, promo: nil},
		&move{s1: D4, s2: F3, promo: nil},
		&move{s1: D4, s2: F5, promo: nil},
	}
	validKnightD4Test = &testStruct{
		Board: map[*Square]Piece{D4: WKnight},
		Moves: validKnightD4Moves,
		Err:   false,
	}
	invalidKnightD4Test = &testStruct{
		Board: map[*Square]Piece{D4: WKnight},
		Moves: allMovesExcluding(D4, validKnightD4Moves),
		Err:   true,
	}
	validWhitePawnE2Moves = []*move{
		&move{s1: E2, s2: E3, promo: nil},
		&move{s1: E2, s2: E4, promo: nil},
	}
	validWhitePawnE2Test = &testStruct{
		Board: map[*Square]Piece{E2: WPawn},
		Moves: validWhitePawnE2Moves,
		Err:   false,
	}
	invalidWhitePawnE2Test = &testStruct{
		Board: map[*Square]Piece{E2: WPawn},
		Moves: allMovesExcluding(E2, validWhitePawnE2Moves),
		Err:   true,
	}
	validBlackPawnA6Moves = []*move{
		&move{s1: A6, s2: A5, promo: nil},
	}
	validBlackPawnA6Test = &testStruct{
		Board: map[*Square]Piece{A6: BPawn},
		Moves: validBlackPawnA6Moves,
		Err:   false,
	}
	invalidBlackPawnA6Test = &testStruct{
		Board: map[*Square]Piece{A6: BPawn},
		Moves: allMovesExcluding(A6, validBlackPawnA6Moves),
		Err:   true,
	}

	blackPawnCaptureTest = &testStruct{
		Board: map[*Square]Piece{G7: BPawn, H6: WRook},
		Moves: []*move{
			&move{s1: G7, s2: H6, promo: nil},
			&move{s1: G7, s2: G6, promo: nil},
			&move{s1: G7, s2: G5, promo: nil},
		},
		Err: false,
	}

	tests = []*testStruct{
		validKingE4Test,
		invalidKingE4Test,
		validKingA1Test,
		invalidKingA1Test,
		validQueenC5Test,
		invalidQueenC5Test,
		validRookC5Test,
		invalidRookC5Test,
		validBishopG2Test,
		invalidBishopG2Test,
		validKnightD4Test,
		invalidKnightD4Test,
		validWhitePawnE2Test,
		invalidWhitePawnE2Test,
		validBlackPawnA6Test,
		invalidBlackPawnA6Test,
		blackPawnCaptureTest,
	}
)
