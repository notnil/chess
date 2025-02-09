package chess

import (
	"testing"
)

func TestBoardTextSerialization(t *testing.T) {
	for _, fen := range []string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
		"rnbqkb1r/pp1p1np1/4pp2/2pP3p/8/2N2NP1/PPP1PP1P/R1BQKBR1",
	} {
		t.Run(fen, func(t *testing.T) {
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
		})
	}

}

func BenchmarkBoardTextSerialization(b *testing.B) {
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	board := &Board{}
	if err := board.UnmarshalText([]byte(fen)); err != nil {
		b.Fatal("recieved unexpected error", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.String()
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

func TestBoardRotation(t *testing.T) {
	fens := []string{
		"RP4pr/NP4pn/BP4pb/QP4pq/KP4pk/BP4pb/NP4pn/RP4pr",
		"RNBKQBNR/PPPPPPPP/8/8/8/8/pppppppp/rnbkqbnr",
		"rp4PR/np4PN/bp4PB/kp4PK/qp4PQ/bp4PB/np4PN/rp4PR",
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
	}
	g := NewGame()
	board := g.Position().Board()
	for i := 0; i < 4; i++ {
		board = board.Rotate()
		if fens[i] != board.String() {
			t.Fatalf("expected board string %s but got %s", fens[i], board.String())
		}
	}
}

func TestBoardFlip(t *testing.T) {
	g := NewGame()
	board := g.Position().Board()
	board = board.Flip(UpDown)
	b := "RNBQKBNR/PPPPPPPP/8/8/8/8/pppppppp/rnbqkbnr"
	if b != board.String() {
		t.Fatalf("expected board string %s but got %s", b, board.String())
	}
	board = board.Flip(UpDown)
	b = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	if b != board.String() {
		t.Fatalf("expected board string %s but got %s", b, board.String())
	}
	board = board.Flip(LeftRight)
	b = "rnbkqbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBKQBNR"
	if b != board.String() {
		t.Fatalf("expected board string %s but got %s", b, board.String())
	}
	board = board.Flip(LeftRight)
	b = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	if b != board.String() {
		t.Fatalf("expected board string %s but got %s", b, board.String())
	}
}

func TestBoardTranspose(t *testing.T) {
	g := NewGame()
	board := g.Position().Board()
	board = board.Transpose()
	b := "rp4PR/np4PN/bp4PB/qp4PQ/kp4PK/bp4PB/np4PN/rp4PR"
	if b != board.String() {
		t.Fatalf("expected board string %s but got %s", b, board.String())
	}
}
