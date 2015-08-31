package chess

import "testing"

type testStruct struct {
	Board Board
	Turn  color
	Moves []*move
	Err   bool
}

// taken from http://en.lichess.org/YvHkKnvo
func TestPartialGame(t *testing.T) {
	moves := []*move{
		{E2, E4, nil},
		{D7, D5, nil},
		{E4, D5, nil},
		{D8, D5, nil},
		{B1, C3, nil},
		{D5, A5, nil},
		{A2, A3, nil},
		{B8, C6, nil},
		{B2, B4, nil},
		{C6, B4, nil},
		{C3, B5, nil},
		{A7, A6, nil},
		{C2, C3, nil},
	}
	g := NewGame()
	for _, m := range moves {
		if err := g.Move(m.s1, m.s2, m.promo); err != nil {
			t.Fatal(err, m, "valid moves", g.board.validMoves(m.s1), g.board)
		}
	}
}

func TestMoves(t *testing.T) {
	for _, tStruct := range tests {
		for _, m := range tStruct.Moves {
			boardCopy := map[*Square]*Piece{}
			for sq, p := range tStruct.Board {
				boardCopy[sq] = p
			}
			g := &Game{
				moves: []*move{},
				board: boardCopy,
				turn:  tStruct.Turn,
			}
			p := g.board.piece(m.s1)
			err := g.Move(m.s1, m.s2, m.promo)
			if tStruct.Err && err == nil {
				t.Errorf("Expected error from piece %s move %+v", p, m)
			} else if !tStruct.Err && err != nil {
				t.Errorf("Received unexpected error %s from piece %s move %s", err.Error(), p, m)
			}
		}
	}
}

func TestCheckmate(t *testing.T) {
	g := &Game{
		moves: []*move{},
		board: map[*Square]*Piece{
			C8: BKing,
			C6: WKing,
			H1: WRook,
		},
		turn:   white,
		status: Ongoing,
	}
	if err := g.Move(H1, H8, nil); err != nil {
		t.Error(err)
	}

	expected := WhiteWon
	if g.status != expected {
		t.Errorf("expected %s got %s", expected, g.status)
	}
}

func TestStalemate(t *testing.T) {
	g := &Game{
		moves: []*move{},
		board: map[*Square]*Piece{
			F7: WKing,
			F5: WQueen,
			H8: BKing,
		},
		turn:   white,
		status: Ongoing,
	}
	if err := g.Move(F5, G6, nil); err != nil {
		t.Error(err)
	}
	expected := Draw
	if g.status != expected {
		t.Errorf("expected %s got %s", expected, g.status)
	}
}

func TestInvalidCastling(t *testing.T) {
	// test king previous move
	g1 := &Game{
		moves: []*move{
			&move{s1: E2, s2: E1, promo: nil},
		},
		board: map[*Square]*Piece{
			A1: WRook,
			E1: WKing,
			H1: WRook,
		},
		turn:   white,
		status: Ongoing,
	}
	if err := g1.Move(E1, G1, nil); err == nil {
		t.Error("allowed castle with a previous move by the king")
	}
	// test rook previous move
	g2 := &Game{
		moves: []*move{
			&move{s1: H2, s2: H1, promo: nil},
		},
		board: map[*Square]*Piece{
			A1: WRook,
			E1: WKing,
			H1: WRook,
		},
		turn:   white,
		status: Ongoing,
	}
	if err := g2.Move(E1, G1, nil); err == nil {
		t.Error("allowed castle with a previous move by the rook")
	}
	// test pieces between
	g3 := &Game{
		moves: []*move{},
		board: map[*Square]*Piece{
			A1: WRook,
			B1: WKnight,
			E1: WKing,
			H1: WRook,
		},
		turn:   white,
		status: Ongoing,
	}
	if err := g3.Move(E1, C1, nil); err == nil {
		t.Error("allowed castle with a piece in between the king and rook")
	}
	// The king is not currently in check.
	g4 := &Game{
		moves: []*move{},
		board: map[*Square]*Piece{
			A1: WRook,
			E1: WKing,
			H1: WRook,
			H4: BBishop,
		},
		turn:   white,
		status: Ongoing,
	}
	if err := g4.Move(E1, G1, nil); err == nil {
		t.Error("allowed castle with the king in check")
	}
	// The king does not pass through a square that is attacked by an enemy piece.[4]
	g5 := &Game{
		moves: []*move{},
		board: map[*Square]*Piece{
			A1: WRook,
			E1: WKing,
			H1: WRook,
			H3: BBishop,
		},
		turn:   white,
		status: Ongoing,
	}
	if err := g5.Move(E1, G1, nil); err == nil {
		t.Error("allowed castle with the king traveling through an attacked square")
	}
}

func TestValidCastle(t *testing.T) {
	g := NewGame()
	moves := []*move{
		{E2, E4, nil},
		{E7, E5, nil},
		{G1, F3, nil},
		{B8, C6, nil},
		{F1, B5, nil},
		{A7, A6, nil},
		{B5, A4, nil},
		{G8, F6, nil},
		{E1, G1, nil},
	}
	for _, m := range moves {
		if err := g.Move(m.s1, m.s2, m.promo); err != nil {
			t.Error(err)
		}
	}

	p := g.board.piece(F1)
	if p == nil || p != WRook {
		t.Error("Rook should have moved.")
	}
}
func TestEnPassant(t *testing.T) {
	// allow if pawn just moved two up
	g1 := &Game{
		moves: []*move{
			&move{s1: A2, s2: A4, promo: nil},
		},
		board: map[*Square]*Piece{
			A4: WPawn,
			B3: WPawn,
			A5: BPawn,
			B4: BPawn,
		},
		turn:   black,
		status: Ongoing,
	}
	if err := g1.Move(B4, A3, nil); err != nil {
		t.Error(err)
	}
	// disallow if pawn didn't just move two up
	g2 := &Game{
		moves: []*move{
			&move{s1: B2, s2: B3, promo: nil},
		},
		board: map[*Square]*Piece{
			A4: WPawn,
			B4: WPawn,
			A5: BPawn,
			B4: BPawn,
		},
		turn:   black,
		status: Ongoing,
	}
	if err := g2.Move(B4, A3, nil); err == nil {
		t.Error("shouldn't allow en passant w/ prior move being two up")
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
		Turn:  white,
		Board: map[*Square]*Piece{E4: WKing},
		Moves: validKingE4Moves,
		Err:   false,
	}
	invalidKingE4Test = &testStruct{
		Turn:  white,
		Board: map[*Square]*Piece{E4: WKing},
		Moves: allMovesExcluding(E4, validKingE4Moves),
		Err:   true,
	}
	validKingA1Moves = []*move{
		&move{s1: A1, s2: A2, promo: nil},
		&move{s1: A1, s2: B1, promo: nil},
		&move{s1: A1, s2: B2, promo: nil},
	}
	validKingA1Test = &testStruct{
		Turn:  white,
		Board: map[*Square]*Piece{A1: WKing},
		Moves: validKingA1Moves,
		Err:   false,
	}
	invalidKingA1Test = &testStruct{
		Turn:  white,
		Board: map[*Square]*Piece{A1: WKing},
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
		Turn:  white,
		Board: map[*Square]*Piece{C5: WQueen},
		Moves: validQueenC5Moves,
		Err:   false,
	}
	invalidQueenC5Test = &testStruct{
		Turn:  white,
		Board: map[*Square]*Piece{C5: WQueen},
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
		Turn:  white,
		Board: map[*Square]*Piece{C5: WRook},
		Moves: validRookC5Moves,
		Err:   false,
	}
	invalidRookC5Test = &testStruct{
		Turn:  white,
		Board: map[*Square]*Piece{C5: WRook},
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
		Turn:  white,
		Board: map[*Square]*Piece{G2: WBishop},
		Moves: validBishopG2Moves,
		Err:   false,
	}
	invalidBishopG2Test = &testStruct{
		Turn:  white,
		Board: map[*Square]*Piece{G2: WBishop},
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
		Turn:  white,
		Board: map[*Square]*Piece{D4: WKnight},
		Moves: validKnightD4Moves,
		Err:   false,
	}
	invalidKnightD4Test = &testStruct{
		Turn:  white,
		Board: map[*Square]*Piece{D4: WKnight},
		Moves: allMovesExcluding(D4, validKnightD4Moves),
		Err:   true,
	}
	validWhitePawnE2Moves = []*move{
		&move{s1: E2, s2: E3, promo: nil},
		&move{s1: E2, s2: E4, promo: nil},
	}
	validWhitePawnE2Test = &testStruct{
		Turn:  white,
		Board: map[*Square]*Piece{E2: WPawn},
		Moves: validWhitePawnE2Moves,
		Err:   false,
	}
	invalidWhitePawnE2Test = &testStruct{
		Turn:  white,
		Board: map[*Square]*Piece{E2: WPawn},
		Moves: allMovesExcluding(E2, validWhitePawnE2Moves),
		Err:   true,
	}
	validBlackPawnA6Moves = []*move{
		&move{s1: A6, s2: A5, promo: nil},
	}
	validBlackPawnA6Test = &testStruct{
		Turn:  black,
		Board: map[*Square]*Piece{A6: BPawn},
		Moves: validBlackPawnA6Moves,
		Err:   false,
	}
	invalidBlackPawnA6Test = &testStruct{
		Turn:  black,
		Board: map[*Square]*Piece{A6: BPawn},
		Moves: allMovesExcluding(A6, validBlackPawnA6Moves),
		Err:   true,
	}
	blackPawnCaptureTest = &testStruct{
		Turn:  black,
		Board: map[*Square]*Piece{G7: BPawn, H6: WRook},
		Moves: []*move{
			&move{s1: G7, s2: H6, promo: nil},
			&move{s1: G7, s2: G6, promo: nil},
			&move{s1: G7, s2: G5, promo: nil},
		},
		Err: false,
	}
	blockedMoveTest = &testStruct{
		Turn:  white,
		Board: map[*Square]*Piece{C6: WRook, C4: WPawn},
		Moves: []*move{
			&move{s1: C6, s2: C4, promo: nil},
			&move{s1: C6, s2: C3, promo: nil},
			&move{s1: C6, s2: C2, promo: nil},
		},
		Err: true,
	}
	validCheckTest = &testStruct{
		Turn: white,
		Board: map[*Square]*Piece{
			A2: BBishop,
			B4: WKnight,
			C6: WRook,
			D7: WPawn,
			E6: WKing,
			F1: BRook,
			H8: BKing,
		},
		Moves: []*move{
			&move{s1: E6, s2: E7, promo: nil},
			&move{s1: E6, s2: E5, promo: nil},
			&move{s1: E6, s2: D6, promo: nil},
			&move{s1: B4, s2: D5, promo: nil},
		},
		Err: false,
	}
	invalidCheckTest = &testStruct{
		Turn: white,
		Board: map[*Square]*Piece{
			A2: BBishop,
			B4: WKnight,
			C6: WRook,
			D7: WPawn,
			E6: WKing,
			F1: BRook,
			H8: BKing,
		},
		Moves: []*move{
			&move{s1: E6, s2: D5, promo: nil},
			&move{s1: E6, s2: F6, promo: nil},
			&move{s1: C6, s2: C3, promo: nil},
		},
		Err: true,
	}
	validWhitePawnPromomtion = &testStruct{
		Turn: white,
		Board: map[*Square]*Piece{
			A7: WPawn,
			B8: BKnight,
		},
		Moves: []*move{
			&move{s1: A7, s2: A8, promo: WQueen},
			&move{s1: A7, s2: A8, promo: WRook},
			&move{s1: A7, s2: B8, promo: WBishop},
			&move{s1: A7, s2: B8, promo: WKnight},
		},
		Err: false,
	}
	invalidWhitePawnPromomtion = &testStruct{
		Turn: white,
		Board: map[*Square]*Piece{
			A7: WPawn,
			B8: BKnight,
		},
		Moves: []*move{
			&move{s1: A7, s2: A8, promo: WKing},
			&move{s1: A7, s2: A8, promo: nil},
			&move{s1: A7, s2: A8, promo: WPawn},
		},
		Err: true,
	}
	validWhiteCastling = &testStruct{
		Turn: white,
		Board: map[*Square]*Piece{
			A1: WRook,
			E1: WKing,
			H1: WRook,
		},
		Moves: []*move{
			&move{s1: E1, s2: G1, promo: nil},
			&move{s1: E1, s2: C1, promo: nil},
		},
		Err: false,
	}
	validBlackCastling = &testStruct{
		Turn: black,
		Board: map[*Square]*Piece{
			A8: BRook,
			E8: BKing,
			H8: BRook,
		},
		Moves: []*move{
			&move{s1: E8, s2: G8, promo: nil},
			&move{s1: E8, s2: C8, promo: nil},
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
		blockedMoveTest,
		validCheckTest,
		invalidCheckTest,
		validWhitePawnPromomtion,
		invalidWhitePawnPromomtion,
		validWhiteCastling,
		validBlackCastling,
	}
)
