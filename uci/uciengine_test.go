package uci

import (
	"fmt"
	"log"
	"os"
	"testing"
	"text/tabwriter"
	chess "github.com/notnil/chess"
)

var stdout = os.Stdout

func TestUciEngine(t *testing.T) {
	var logger *log.Logger = log.New(stdout, "", log.LstdFlags)

	e, err := Run("c:/lc0/lc0.exe", nil, logger)
	if err != nil {
		t.Fatalf("%s", err)
	}
	defer e.Quit()

	opt := e.Options()
	w := tabwriter.NewWriter(stdout, 1, 8, 0, ' ', 0)
	for k, v := range opt {
		fmt.Fprintln(w, k, "\t", v)
	}
	w.Flush()

	position, err := chess.DecodeFEN("5r2/pqN1Qpk1/2r3pp/2n1R3/5R2/6P1/4PPKP/8 w - - 0 1")
	e.SetPosition(position)
	//board := MustParseFen()
	//e.SetPosition(board)
	for info := range e.SearchDepth(3) {
		if info.Err() != nil {
			t.Fatalf("%s", info.Err())
		}
		if m, ok := info.BestMove(); ok {
			log.Println("bestmove:", m.String()) //m.San(board))
		} else if pv := info.Pv(); pv != nil {
			log.Println("pv:", pv)
			log.Println("stats:", info.Stats())
		} else {
			log.Println("stats:", info.Stats()) //info. .Line)
		}
	}
}
