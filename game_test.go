package chess

import (
	"log"
	"strings"
	"testing"
)

func TestCheckmate(t *testing.T) {
	fenStr := "rn1qkbnr/pbpp1ppp/1p6/4p3/2B1P3/5Q2/PPPP1PPP/RNB1K1NR w KQkq - 0 1"
	fen, err := FEN(fenStr)
	if err != nil {
		t.Fatal(err)
	}
	g := NewGame(fen)
	if m, err := g.MoveStr("Qxf7#"); err != nil {
		t.Fatal(err)
	} else if m.piece != WhiteQueen || m.GetTags() != Check|Capture {
		t.Fatalf("expected %s move with tags %d but got %s %s %d", "WhiteQueen", Check|Capture, m.piece.Color().Name(), m.piece.Type().String(), m.GetTags())
	}
	if g.Method() != Checkmate {
		t.Fatalf("expected method %s but got %s", Checkmate, g.Method())
	}
	if g.Outcome() != WhiteWon {
		t.Fatalf("expected outcome %s but got %s", WhiteWon, g.Outcome())
	}

	// Checkmate on castle
	fenStr = "Q7/5Qp1/3k2N1/7p/8/4B3/PP3PPP/R3K2R w KQ - 0 31"
	fen, err = FEN(fenStr)
	if err != nil {
		t.Fatal(err)
	}
	g = NewGame(fen)
	if m, err := g.MoveStr("O-O-O"); err != nil {
		t.Fatal(err)
	} else if m.piece != WhiteKing || m.GetTags() != Check|QueenSideCastle {
		t.Fatalf("expected %s move with tags %d but got %s %s %d", "WhiteKing", Check|QueenSideCastle, m.piece.Color().Name(), m.piece.Type().String(), m.GetTags())
	}

	if g.Method() != Checkmate {
		t.Fatalf("expected method %s but got %s", Checkmate, g.Method())
	}
	if g.Outcome() != WhiteWon {
		t.Fatalf("expected outcome %s but got %s", WhiteWon, g.Outcome())
	}

}

func TestCheckmateFromFen(t *testing.T) {
	fenStr := "rn1qkbnr/pbpp1Qpp/1p6/4p3/2B1P3/8/PPPP1PPP/RNB1K1NR b KQkq - 0 1"
	fen, err := FEN(fenStr)
	if err != nil {
		t.Fatal(err)
	}
	g := NewGame(fen)
	if g.Method() != Checkmate {
		t.Error(g.Position().Board().Draw())
		t.Fatalf("expected method %s but got %s", Checkmate, g.Method())
	}
	if g.Outcome() != WhiteWon {
		t.Fatalf("expected outcome %s but got %s", WhiteWon, g.Outcome())
	}
}

func TestStalemate(t *testing.T) {
	fenStr := "k1K5/8/8/8/8/8/8/1Q6 w - - 0 1"
	fen, err := FEN(fenStr)
	if err != nil {
		t.Fatal(err)
	}
	g := NewGame(fen)
	if _, err := g.MoveStr("Qb6"); err != nil {
		t.Fatal(err)
	}
	if g.Method() != Stalemate {
		t.Fatalf("expected method %s but got %s", Stalemate, g.Method())
	}
	if g.Outcome() != Draw {
		t.Fatalf("expected outcome %s but got %s", Draw, g.Outcome())
	}
}

// position shouldn't result in stalemate because pawn can move http://en.lichess.org/Pc6mJDZN#138
func TestInvalidStalemate(t *testing.T) {
	fenStr := "8/3P4/8/8/8/7k/7p/7K w - - 2 70"
	fen, err := FEN(fenStr)
	if err != nil {
		t.Fatal(err)
	}
	g := NewGame(fen)
	if m, err := g.MoveStr("d8=Q"); err != nil {
		t.Fatal(err)
	} else if m.piece != WhitePawn || m.promo != Queen || m.GetTags() != 0 {
		t.Fatalf("expected %s move with tags %d but got %s %s %d", "WhitePawn", 0, m.piece.Color().Name(), m.piece.Type().String(), m.GetTags())
	}
	if g.Outcome() != NoOutcome {
		t.Fatalf("expected outcome %s but got %s", NoOutcome, g.Outcome())
	}
}

func TestMoveSquares(t *testing.T) {
	// Check if moving via Square objects vs strings produces the same result
	type StrPair struct {
		from, to string
	}

	g1 := NewGame(UseNotation(UCINotation{}))
	g2 := NewGame(UseNotation(UCINotation{}))

	moves := []StrPair{
		{"e2", "e4"}, {"g8", "f6"},
		{"e4", "e5"}, {"f6", "d5"},
		{"c2", "c4"}, {"d5", "b6"},
		{"d2", "d4"}, {"d7", "d6"},
		{"f2", "f4"}}

	for i, m := range moves {
		mStr := m.from + m.to
		m1, err := g1.MoveStr(mStr)

		if err != nil {
			t.Fatalf("Error on MoveStr %d. %s: %s", i+1, mStr, err)
		}

		m2, err := g2.MoveSquares(NewSquareFromStr(m.from), NewSquareFromStr(m.to), NoPieceType)

		if err != nil {
			t.Fatalf("Error on MoveSquare %d. %s: %s", i+1, mStr, err)
		}

		if m1.String() != m2.String() || m1.GetTags() != m2.GetTags() {
			t.Fatalf("MoveStr and MoveSquare do not match for %d. %s.\nMoveStr:\t\t%s tags: %d\nMoveSquare:\t\t%s tags: %d", i+1, mStr, m1.String(), m1.GetTags(), m2.String(), m2.GetTags())
		}
	}

	if g1.String() != g2.String() {
		t.Fatalf("Games do not match.\nMoveStr:\n%s\n\nMoveSquare:\n%s", g1.String(), g2.String())
	}

	fen, _ := FEN("8/8/8/8/8/8/p7/8 b - - 0 1")
	g3 := NewGame(fen)

	m, err := g3.MoveSquares(A2, A1, Knight)

	if err != nil {
		t.Fatal(err)
	}

	correct := &Move{s1: A2, s2: A1, promo: Knight, piece: BlackPawn}

	if m.String() != correct.String() {
		t.Fatalf("Incorrect move %s. Expected %s", m.String(), correct.String())
	}

	piece := g3.Position().Board().Piece(A1)

	if piece != BlackKnight {
		t.Fatalf("Incorrect promotion %s %s. Expected Black Knight", piece.Color().Name(), piece.Type().String())
	}
}

func TestThreeFoldRepition(t *testing.T) {
	g := NewGame()
	moves := []string{
		"Nf3", "Nf6", "Ng1", "Ng8",
		"Nf3", "Nf6", "Ng1", "Ng8",
	}
	for _, m := range moves {
		if _, err := g.MoveStr(m); err != nil {
			t.Fatal(err)
		}
	}
	if err := g.Draw(ThreefoldRepetition); err != nil {
		for _, pos := range g.Positions() {
			log.Println(pos.String())
		}
		t.Fatalf("%s - %d reps", err.Error(), g.numOfRepitions())
	}
}

func TestInvalidThreeFoldRepition(t *testing.T) {
	g := NewGame()
	moves := []string{
		"Nf3", "Nf6", "Ng1", "Ng8",
	}
	for _, m := range moves {
		if _, err := g.MoveStr(m); err != nil {
			t.Fatal(err)
		}
	}
	if err := g.Draw(ThreefoldRepetition); err == nil {
		t.Fatal("should require three repeated board states")
	}
}

func TestFiveFoldRepition(t *testing.T) {
	g := NewGame()
	moves := []string{
		"Nf3", "Nf6", "Ng1", "Ng8",
		"Nf3", "Nf6", "Ng1", "Ng8",
		"Nf3", "Nf6", "Ng1", "Ng8",
		"Nf3", "Nf6", "Ng1", "Ng8",
	}
	for _, m := range moves {
		if _, err := g.MoveStr(m); err != nil {
			t.Fatal(err)
		}
	}
	if g.Outcome() != Draw || g.Method() != FivefoldRepetition {
		t.Fatal("should automatically draw after five repetitions")
	}
}

func TestFiftyMoveRule(t *testing.T) {
	fen, _ := FEN("2r3k1/1q1nbppp/r3p3/3pP3/pPpP4/P1Q2N2/2RN1PPP/2R4K b - b3 100 60")
	g := NewGame(fen)
	if err := g.Draw(FiftyMoveRule); err != nil {
		t.Fatal(err)
	}
}

func TestInvalidFiftyMoveRule(t *testing.T) {
	fen, _ := FEN("2r3k1/1q1nbppp/r3p3/3pP3/pPpP4/P1Q2N2/2RN1PPP/2R4K b - b3 99 60")
	g := NewGame(fen)
	if err := g.Draw(FiftyMoveRule); err == nil {
		t.Fatal("should require fifty moves")
	}
}

func TestSeventyFiveMoveRule(t *testing.T) {
	fen, _ := FEN("2r3k1/1q1nbppp/r3p3/3pP3/pPpP4/P1Q2N2/2RN1PPP/2R4K b - b3 149 80")
	g := NewGame(fen)
	if _, err := g.MoveStr("Kf8"); err != nil {
		t.Fatal(err)
	}
	if g.Outcome() != Draw || g.Method() != SeventyFiveMoveRule {
		t.Fatal("should automatically draw after seventy five moves w/ no pawn move or capture")
	}
}

func TestInsufficentMaterial(t *testing.T) {
	fens := []string{
		"8/2k5/8/8/8/3K4/8/8 w - - 1 1",
		"8/2k5/8/8/8/3K1N2/8/8 w - - 1 1",
		"8/2k5/8/8/8/3K1B2/8/8 w - - 1 1",
		"8/2k5/2b5/8/8/3K1B2/8/8 w - - 1 1",
		"4b3/2k5/2b5/8/8/3K1B2/8/8 w - - 1 1",
	}
	for _, f := range fens {
		fen, err := FEN(f)
		if err != nil {
			t.Fatal(err)
		}
		g := NewGame(fen)
		if g.Outcome() != Draw || g.Method() != InsufficientMaterial {
			log.Println(g.Position().Board().Draw())
			t.Fatalf("%s should automatically draw by insufficent material", f)
		}
	}
}

func TestSufficentMaterial(t *testing.T) {
	fens := []string{
		"8/2k5/8/8/8/3K1B2/4N3/8 w - - 1 1",
		"8/2k5/8/8/8/3KBB2/8/8 w - - 1 1",
		"8/2k1b3/8/8/8/3K1B2/8/8 w - - 1 1",
		"8/2k5/8/8/4P3/3K4/8/8 w - - 1 1",
		"8/2k5/8/8/8/3KQ3/8/8 w - - 1 1",
		"8/2k5/8/8/8/3KR3/8/8 w - - 1 1",
	}
	for _, f := range fens {
		fen, err := FEN(f)
		if err != nil {
			t.Fatal(err)
		}
		g := NewGame(fen)
		if g.Outcome() != NoOutcome {
			log.Println(g.Position().Board().Draw())
			t.Fatalf("%s should not find insufficent material", f)
		}
	}
}

func TestSerializationCycle(t *testing.T) {
	g := NewGame()
	g.MoveStr("e4")
	g.MoveStr("e5")
	pgn, err := PGN(strings.NewReader(g.String()))
	if err != nil {
		t.Fatal(err)
	}
	cp := NewGame(pgn)
	if cp.String() != g.String() {
		t.Fatalf("expected %s but got %s", g.String(), cp.String())
	}
}

func TestInitialNumOfValidMoves(t *testing.T) {
	g := NewGame()
	if len(g.ValidMoves()) != 20 {
		t.Fatal("should find 20 valid moves from the initial position")
	}
}

func TestTagPairs(t *testing.T) {
	g := NewGame()
	g.AddTagPair("Draw Offer", "White")
	tagPair := g.GetTagPair("Draw Offer")
	if tagPair == nil {
		t.Fatalf("expected %s but got %s", "White", "nil")
	}
	if tagPair.Value != "White" {
		t.Fatalf("expected %s but got %s", "White", tagPair.Value)
	}
	g.RemoveTagPair("Draw Offer")
	tagPair = g.GetTagPair("Draw Offer")
	if tagPair != nil {
		t.Fatalf("expected %s but got %s", "nil", "not nil")
	}
}

func TestPositionHash(t *testing.T) {
	g1 := NewGame()
	for _, s := range []string{"Nc3", "e5", "Nf3"} {
		g1.MoveStr(s)
	}
	g2 := NewGame()
	for _, s := range []string{"Nf3", "e5", "Nc3"} {
		g2.MoveStr(s)
	}
	if g1.Position().Hash() != g2.Position().Hash() {
		t.Fatalf("expected position hashes to be equal but got %s and %s", g1.Position().Hash(), g2.Position().Hash())
	}
}

func BenchmarkStalemateStatus(b *testing.B) {
	fenStr := "k1K5/8/8/8/8/8/8/1Q6 w - - 0 1"
	fen, err := FEN(fenStr)
	if err != nil {
		b.Fatal(err)
	}
	g := NewGame(fen)
	if _, err := g.MoveStr("Qb6"); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		g.Position().Status()
	}
}

func BenchmarkInvalidStalemateStatus(b *testing.B) {
	fenStr := "8/3P4/8/8/8/7k/7p/7K w - - 2 70"
	fen, err := FEN(fenStr)
	if err != nil {
		b.Fatal(err)
	}
	g := NewGame(fen)
	if _, err := g.MoveStr("d8=Q"); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		g.Position().Status()
	}
}

func BenchmarkPositionHash(b *testing.B) {
	fenStr := "8/3P4/8/8/8/7k/7p/7K w - - 2 70"
	fen, err := FEN(fenStr)
	if err != nil {
		b.Fatal(err)
	}
	g := NewGame(fen)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		g.Position().Hash()
	}
}
