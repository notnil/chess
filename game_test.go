package chess

import (
	"strings"
	"testing"
)

func TestCheckmate(t *testing.T) {
	fenStr := "rn1qkbnr/pbpp1ppp/1p6/4p3/2B1P3/5Q2/PPPP1PPP/RNB1K1NR w KQkq - 0 1"
	fen, err := FEN(strings.NewReader(fenStr))
	if err != nil {
		t.Fatal(err)
	}
	g := NewGame(fen)
	if err := g.Move(F3, F7, NoPiece); err != nil {
		t.Fatal(err)
	}
	if g.Method() != Checkmate {
		t.Fatalf("expected method %s but got %s", Checkmate, g.Method())
	}
	if g.Outcome() != WhiteWon {
		t.Fatalf("expected outcome %s but got %s", WhiteWon, g.Outcome())
	}
}

func TestStalemate(t *testing.T) {
	fenStr := "k1K5/8/8/8/8/8/8/1Q6 w KQkq - 0 1"
	fen, err := FEN(strings.NewReader(fenStr))
	if err != nil {
		t.Fatal(err)
	}
	g := NewGame(fen)
	if err := g.Move(B1, B6, NoPiece); err != nil {
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
	fen, err := FEN(strings.NewReader(fenStr))
	if err != nil {
		t.Fatal(err)
	}
	g := NewGame(fen)
	if err := g.Move(D7, D8, Queen); err != nil {
		t.Fatal(err)
	}
	if g.Outcome() != NoOutcome {
		t.Fatalf("expected outcome %s but got %s", NoOutcome, g.Outcome())
	}
}

func TestTakeBack(t *testing.T) {
	g := NewGame()
	if err := g.MoveAlg("e4"); err != nil {
		t.Fatal(err)
	}
	g = g.TakeBack(1)
	if startFEN != g.FEN() {
		t.Fatalf("take back expected fen %s but got %s", startFEN, g.FEN())
	}
	if len(NewGame().moves) != len(g.moves) {
		t.Fatalf("take back expected %d move but got %d", len(NewGame().moves), len(g.moves))
	}
}
