package chess

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type pgnTest struct {
	PostPos *Position
	PGN     string
}

var (
	validPGNs = []pgnTest{
		{
			PostPos: unsafeFEN("4r3/6P1/2p2P1k/1p6/pP2p1R1/P1B5/2P2K2/3r4 b - - 0 45"),
			PGN:     mustParsePGN("fixtures/pgns/0001.pgn"),
		},
		{
			PostPos: unsafeFEN("4r3/6P1/2p2P1k/1p6/pP2p1R1/P1B5/2P2K2/3r4 b - - 0 45"),
			PGN:     mustParsePGN("fixtures/pgns/0002.pgn"),
		},
		{
			PostPos: unsafeFEN("2r2rk1/pp1bBpp1/2np4/2pp2p1/1bP5/1P4P1/P1QPPPBP/3R1RK1 b - - 0 3"),
			PGN:     mustParsePGN("fixtures/pgns/0003.pgn"),
		},
		{
			PostPos: unsafeFEN("r3kb1r/2qp1pp1/b1n1p2p/pp2P3/5n1B/1PPQ1N2/P1BN1PPP/R3K2R w KQkq - 1 14"),
			PGN:     mustParsePGN("fixtures/pgns/0004.pgn"),
		},
		{
			PostPos: unsafeFEN("rnbqkbnr/ppp2ppp/4p3/3p4/3PP3/8/PPP2PPP/RNBQKBNR w KQkq d6 0 3"),
			PGN:     mustParsePGN("fixtures/pgns/0008.pgn"),
		},
		{
			PostPos: unsafeFEN("r1bqkbnr/1ppp1ppp/p1n5/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R w KQkq - 0 4"),
			PGN:     mustParsePGN("fixtures/pgns/0009.pgn"),
		},
		{
			PostPos: unsafeFEN("r1bqkbnr/1ppp1ppp/p1n5/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R w KQkq - 0 4"),
			PGN:     mustParsePGN("fixtures/pgns/0010.pgn"),
		},
		{
			PostPos: unsafeFEN("8/8/6p1/4R3/6kQ/r2P1pP1/5P2/6K1 b - - 3 42"),
			PGN:     mustParsePGN("fixtures/pgns/0011.pgn"),
		},
		{
			PostPos: StartingPosition(),
			PGN:     mustParsePGN("fixtures/pgns/0012.pgn"),
		},
	}
)

func TestValidPGNs(t *testing.T) {
	for _, test := range validPGNs {
		game, err := decodePGN(test.PGN)
		if err != nil {
			t.Fatalf("recieved unexpected pgn error %s", err.Error())
		}
		if game.Position().String() != test.PostPos.String() {
			t.Fatalf("expected board to be \n%s\nFEN:%s\n but got \n%s\n\nFEN:%s\n",
				test.PostPos.board.Draw(), test.PostPos.String(),
				game.Position().board.Draw(), game.Position().String())
		}
	}
}

type commentTest struct {
	PGN         string
	MoveNumber  int
	CommentText string
}

var (
	commentTests = []commentTest{
		{
			PGN:         mustParsePGN("fixtures/pgns/0005.pgn"),
			MoveNumber:  7,
			CommentText: `(-0.25 → 0.39) Inaccuracy. cxd4 was best. [%eval 0.39] [%clk 0:05:05]`,
		},
		{
			PGN:         mustParsePGN("fixtures/pgns/0009.pgn"),
			MoveNumber:  5,
			CommentText: `This opening is called the Ruy Lopez.`,
		},
		{
			PGN:         mustParsePGN("fixtures/pgns/0010.pgn"),
			MoveNumber:  5,
			CommentText: `This opening is called the Ruy Lopez.`,
		},
	}
)

func TestCommentsDetection(t *testing.T) {
	for _, test := range commentTests {
		game, err := decodePGN(test.PGN)
		if err != nil {
			t.Fatal(err)
		}
		comment := strings.Join(game.Comments()[test.MoveNumber], " ")
		if comment != test.CommentText {
			t.Fatalf("expected pgn comment to be %s but got %s", test.CommentText, comment)
		}
	}
}

func TestNewGameComments(t *testing.T) {
	for _, test := range commentTests {
		pgn, err := PGN(strings.NewReader(test.PGN))
		if err != nil {
			t.Fatal(err)
		}
		game := NewGame(pgn)
		comment := strings.Join(game.Comments()[test.MoveNumber], " ")
		if comment != test.CommentText {
			t.Fatalf("expected pgn comment to be %s but got %s", test.CommentText, comment)
		}
	}
}

func TestWriteComments(t *testing.T) {
	pgn := mustParsePGN("fixtures/pgns/0005.pgn")
	game, err := decodePGN(pgn)
	if err != nil {
		t.Fatal(err)
	}
	game, err = decodePGN(game.String())
	if err != nil {
		t.Fatal(err)
	}
	if len(game.Comments()[7]) != 2 {
		t.Fatalf("expected %d comments for move 7 but got %d", 2, len(game.Comments()[7]))
	}
}

func TestScanner(t *testing.T) {
	for _, fname := range []string{"fixtures/pgns/0006.pgn", "fixtures/pgns/0007.pgn"} {
		f, err := os.Open(fname)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		scanner := NewScanner(f)
		games := []*Game{}
		for scanner.Scan() {
			game := scanner.Next()
			games = append(games, game)
		}
		if len(games) != 5 {
			t.Fatalf(fname+" expected 5 games but got %d", len(games))
		}
	}
}

func TestScannerWithNested(t *testing.T) {
	fname := "fixtures/pgns/0013.pgn"
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := NewScanner(f)
	games := []*Game{}
	for scanner.Scan() {
		err = scanner.Err()
		if err != nil && err != io.EOF {
			t.Fatalf(fname+" Unexpected non-nil/non-EOF err %v", err)
		}
		game := scanner.Next()
		moveList := game.Moves()
		if len(moveList) == 0 {
			continue
		}
		games = append(games, game)
	}
	err = scanner.Err()
	if err != io.EOF {
		t.Fatalf(fname+" Unexpected non-EOF err %v", err)
	}
	if len(games) != 3 {
		for idx, g := range games {
			fmt.Printf("Parsed game %v: %v\n\n", idx, g)
		}
		t.Fatalf(fname+" expected 3 games but got %d", len(games))
	}
}

func TestScannerWithNestedAndExpand(t *testing.T) {
	fname := "fixtures/pgns/0013.pgn"
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scannerOpts := ScannerOpts{ExpandVariations: true}
	scanner := NewScannerWithOptions(f, scannerOpts)
	games := []*Game{}
	for scanner.Scan() {
		err = scanner.Err()
		if err != nil && err != io.EOF {
			t.Fatalf(fname+" Unexpected non-nil/non-EOF err %v", err)
		}
		game := scanner.Next()
		moveList := game.Moves()
		if len(moveList) == 0 {
			continue
		}
		games = append(games, game)
	}
	err = scanner.Err()
	if err != io.EOF {
		t.Fatalf(fname+" Unexpected non-EOF err %v", err)
	}
	if len(games) != 10 {
		for idx, g := range games {
			fmt.Printf("Parsed game %v: %v\n\n", idx, g)
		}
		t.Fatalf(fname+" expected 10 games but got %d", len(games))
	}
}

func BenchmarkPGN(b *testing.B) {
	pgn := mustParsePGN("fixtures/pgns/0001.pgn")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		opt, _ := PGN(strings.NewReader(pgn))
		NewGame(opt)
	}
}

func mustParsePGN(fname string) string {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return string(b)
}
func TestGamesFromPGN(t *testing.T) {
	for _, test := range validPGNs {
		reader := strings.NewReader(test.PGN)
		games, err := GamesFromPGN(reader)
		if err != nil {
			t.Fatalf("fail to read games from valid pgn: %s", err.Error())
		}
		if len(games) != 1 {
			t.Fatalf("expected to get 1 game from pgn, got %d", len(games))
		}
	}
}
