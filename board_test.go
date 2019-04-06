package chess

import (
	"testing"
)

func TestBoardTextSerialization(t *testing.T) {
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	b := &Board{}
	if err := b.UnmarshalText([]byte(fen)); err != nil {
		t.Fatal("recieved unexpected error", err)
	}
	txt, err := b.MarshalText()
	if err != nil {
		t.Fatal("recieved unexpected error", err)
	}
	if fen != string(txt) {
		t.Fatalf("fen expected board string %s but got %s", fen, string(txt))
	}
}

func TestBoardBinarySerialization(t *testing.T) {
	g := NewGame()
	board := g.Position().Board()
	b, err := board.MarshalBinary()
	if err != nil {
		t.Fatal("recieved unexpected error", err)
	}
	cpBoard := &Board{}
	err = cpBoard.UnmarshalBinary(b)
	if err != nil {
		t.Fatal("recieved unexpected error", err)
	}
	s := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	if s != cpBoard.String() {
		t.Fatalf("expected board string %s but got %s", s, cpBoard.String())
	}
}
