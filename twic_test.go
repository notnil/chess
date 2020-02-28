package chess

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestTWICPGNScanner2(t *testing.T) {
	games := []*Game{}

	inputPath := GetTwicArchiveFiles(1)

	in, err := os.Open(inputPath)
	if err != nil {
		t.Fatalf("unexpected pgn read error %s", err.Error())
	}
	defer in.Close()

	scanner := bufio.NewScanner(in)
	scanner.Split(splitPGN)
	i := 0
	for scanner.Scan() {
		game := scanner.Text()
		decgame, err := decodePGN(game)
		if err != nil {
			fmt.Println("Error decoding game ", i)
		} else {
			games = append(games, decgame)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}

func TestPGNScanner1(t *testing.T) {
	games := []*Game{}

	inputPath := "./assets/twic0920.pgn"
	in, err := os.Open(inputPath)
	if err != nil {
		t.Fatalf("unexpected pgn read error %s", err.Error())
	}
	defer in.Close()

	scanner := bufio.NewScanner(in)
	scanner.Split(splitPGN)
	i := 0
	for scanner.Scan() {
		game := scanner.Text()
		decgame, err := decodePGN(game)
		if err != nil {
			t.Fatalf("Error decoding game %d", i)
		}
		games = append(games, decgame)
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	if i != 2245 {
		t.Fatalf("2245 games should have been read from %s. %d games read", inputPath, i)
	}

	// create FEN for all games
	a := []string{}
	for _, game := range games {
		for _, pos := range game.Positions() {
			a = append(a, pos.String())
		}
	}

	if a[0] != "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1" {
		t.Fatalf("Fen not correct. Should be: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1, but is: %s ", a[0])
	}

	if a[1058] != "6k1/p2q1b2/1p1n1B1p/2p3p1/2P5/1PN3QP/P5P1/6K1 b - - 1 37" {
		t.Fatalf("Fen not correct. Should be: 6k1/p2q1b2/1p1n1B1p/2p3p1/2P5/1PN3QP/P5P1/6K1 b - - 1 37, but is: %s ", a[1058])
	}

	if a[184726] != "2rrn1k1/2qNbp1p/4p3/p3P3/2NnQpP1/7P/P4BB1/1R2R1K1 b - - 2 30" {
		t.Fatalf("Fen not correct. Should be: 2rrn1k1/2qNbp1p/4p3/p3P3/2NnQpP1/7P/P4BB1/1R2R1K1 b - - 2 30, but is: %s ", a[184726])
	}
}
