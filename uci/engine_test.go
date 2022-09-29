package uci_test

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
)

var StockfishPath string

func init() {
	dir, _ := os.Getwd()
	StockfishPath = filepath.Join(dir, "..", "stockfish")
}

func Example() {
	// set up engine to use stockfish exe
	eng, err := uci.New(StockfishPath)
	if err != nil {
		panic(err)
	}
	defer eng.Close()
	// initialize uci with new game
	if err := eng.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
		panic(err)
	}
	// have stockfish play speed chess against itself (10 msec per move)
	game := chess.NewGame()
	for game.Outcome() == chess.NoOutcome {
		cmdPos := uci.CmdPosition{Position: game.Position()}
		cmdGo := uci.CmdGo{MoveTime: time.Second / 100}
		if err := eng.Run(cmdPos, cmdGo); err != nil {
			panic(err)
		}
		move := eng.SearchResults().BestMove
		if err := game.Move(move); err != nil {
			panic(err)
		}
	}
	fmt.Println(game.String())
}

func TestEngine(t *testing.T) {
	eng, err := uci.New(StockfishPath)
	if err != nil {
		t.Fatal(err)
	}
	defer eng.Close()
	setOpt := uci.CmdSetOption{Name: "UCI_Elo", Value: "1500"}
	setPos := uci.CmdPosition{Position: chess.StartingPosition()}
	setGo := uci.CmdGo{MoveTime: time.Second / 10}
	if err := eng.Run(uci.CmdUCI, uci.CmdIsReady, setOpt, uci.CmdUCINewGame, setPos, setGo); err != nil {
		t.Fatal(err)
	}
	if eng.SearchResults().BestMove.S2() != chess.E4 {
		t.Fatal("expected a different move")
	}
	pos := &chess.Position{}
	pos.UnmarshalText([]byte("r4r2/1b2bppk/ppq1p3/2pp3n/5P2/1P2P3/PBPPQ1PP/R4RK1 w - - 0 2"))
	setPos.Position = pos
	if err := eng.Run(uci.CmdUCINewGame, setPos, setGo); err != nil {
		t.Fatal(err)
	}
	if eng.SearchResults().BestMove.S2() != chess.H5 {
		t.Fatal("expected a different move")
	}
}

func TestStop(t *testing.T) {
	eng, err := uci.New(StockfishPath)
	if err != nil {
		t.Fatal(err)
	}
	defer eng.Close()
	go func() {
		if err := eng.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame, uci.CmdGo{Infinite: true}); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second)
	if err := eng.Run(uci.CmdStop); err != nil {
		t.Fatal(err)
	}
	if eng.SearchResults().BestMove.S2() != chess.D4 {
		t.Fatal("expected a different move")
	}
}

func TestLogger(t *testing.T) {
	t.SkipNow()

	b := bytes.NewBuffer([]byte{})
	logger := log.New(b, "", 0)
	eng, err := uci.New(StockfishPath, uci.Debug, uci.Logger(logger))
	if err != nil {
		t.Fatal(err)
	}
	defer eng.Close()
	setOpt := uci.CmdSetOption{Name: "UCI_Elo", Value: "1500"}
	setPos := uci.CmdPosition{Position: chess.StartingPosition()}
	setGo := uci.CmdGo{MoveTime: time.Second / 10}
	if err := eng.Run(uci.CmdUCI, uci.CmdIsReady, setOpt, uci.CmdUCINewGame, setPos, setGo); err != nil {
		t.Fatal(err)
	}
	expected := infoRegex.ReplaceAllString(logOutput, "")
	expectedSplit := strings.Split(expected, "\n")
	actual := b.String()
	actual = infoRegex.ReplaceAllString(actual, "")
	actualSplit := strings.Split(actual, "\n")
	for i := 0; i < len(expectedSplit); i++ {
		expected = strings.TrimSpace(expectedSplit[i])
		actual = strings.TrimSpace(actualSplit[i])
		if expected != actual {
			t.Fatalf("expected %s but got %s", expected, actual)
		}
	}
}

func TestEngineOutputParsing(t *testing.T){
	scanner := bufio.NewScanner(strings.NewReader(logOutput))

	results, err := uci.ProcessEngineOutput(scanner, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(results.MultiPV) != 1 {
		t.Fatalf("expected 1 multipv results but got %d", len(results.MultiPV))
	}

	if results.MultiPV[0].Score.CP != 50 {
		t.Fatalf("expected score of 50 but got %d", results.MultiPV[0].Score.CP)
	}
	if results.Info.Score.CP != 50 {
		t.Fatalf("expected score of 50 but got %d", results.Info.Score.CP)
	}

	// Parse BestMove and assert it is expected
	if results.BestMove.S2() != chess.E4 {
		t.Fatalf("expected E4 to be best move square, got %s", results.BestMove.S2())
	}
}

func TestMultiPVEngineOutputParsing(t *testing.T){
	scanner := bufio.NewScanner(strings.NewReader(multipvTestingString))

	results, err := uci.ProcessEngineOutput(scanner, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(results.MultiPV) != 10 {
		t.Fatalf("expected 10 multipv results but got %d", len(results.MultiPV))
	}

	if results.MultiPV[0].Score.CP != 47 {
		t.Fatalf("expected score of 47 but got %d", results.MultiPV[0].Score.CP)
	}
	if results.MultiPV[1].Score.CP != 31 {
		t.Fatalf("expected score of 31 but got %d", results.MultiPV[1].Score.CP)
	}
	if results.Info.Score.CP != 47 {
		t.Fatalf("expected score of 47 but got %d", results.Info.Score.CP)
	}

	// Parse BestMove and assert it is expected
	if results.BestMove.S2() != chess.E4 {
		t.Fatalf("expected E4 to be best move square, got %s", results.BestMove.S2())
	}
}

var (
	infoRegex = regexp.MustCompile("(?m)[\r\n]+^.*info.*$")
)

const (
	logOutput = `uci
Stockfish 14.1 by the Stockfish developers (see AUTHORS file)
id name Stockfish 14.1
id author the Stockfish developers (see AUTHORS file)
	
option name Debug Log File type string default 
option name Threads type spin default 1 min 1 max 512
option name Hash type spin default 16 min 1 max 33554432
option name Clear Hash type button
option name Ponder type check default false
option name MultiPV type spin default 1 min 1 max 500
option name Skill Level type spin default 20 min 0 max 20
option name Move Overhead type spin default 10 min 0 max 5000
option name Slow Mover type spin default 100 min 10 max 1000
option name nodestime type spin default 0 min 0 max 10000
option name UCI_Chess960 type check default false
option name UCI_AnalyseMode type check default false
option name UCI_LimitStrength type check default false
option name UCI_Elo type spin default 1350 min 1350 max 2850
option name UCI_ShowWDL type check default false
option name SyzygyPath type string default <empty>
option name SyzygyProbeDepth type spin default 1 min 1 max 100
option name Syzygy50MoveRule type check default true
option name SyzygyProbeLimit type spin default 7 min 0 max 7
option name Use NNUE type check default true
option name EvalFile type string default nn-13406b1dcbe0.nnue
uciok
isready
readyok
setoption name UCI_Elo value 1500
ucinewgame
position fen rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
go movetime 100
info string NNUE evaluation using nn-13406b1dcbe0.nnue enabled
info depth 1 seldepth 1 multipv 1 score cp 38 nodes 20 nps 20000 tbhits 0 time 1 pv d2d4
info depth 2 seldepth 2 multipv 1 score cp 82 nodes 51 nps 51000 tbhits 0 time 1 pv e2e4 a7a6
info depth 3 seldepth 3 multipv 1 score cp 55 nodes 154 nps 154000 tbhits 0 time 1 pv e2e4 c7c6 d2d4
info depth 4 seldepth 4 multipv 1 score cp 22 nodes 807 nps 269000 tbhits 0 time 3 pv g1f3 d7d5 d2d4 g8f6
info depth 5 seldepth 5 multipv 1 score cp 54 nodes 1061 nps 353666 tbhits 0 time 3 pv e2e4 c7c5 g1f3
info depth 6 seldepth 6 multipv 1 score cp 54 nodes 1761 nps 440250 tbhits 0 time 4 pv e2e4 c7c5 g1f3 d7d5 e4d5 d8d5
info depth 7 seldepth 8 multipv 1 score cp 50 nodes 5459 nps 545900 tbhits 0 time 10 pv e2e4 e7e5 g1f3 b8c6 b1c3 g8f6 d2d4
info depth 8 seldepth 8 multipv 1 score cp 50 nodes 6998 nps 583166 tbhits 0 time 12 pv e2e4 e7e5 g1f3 b8c6 d2d4 e5d4 f3d4 g8f6 b1c3
info depth 9 seldepth 11 multipv 1 score cp 54 nodes 12053 nps 573952 tbhits 0 time 21 pv e2e4 e7e5 g1f3 b8c6 d2d4 e5d4 f3d4 g8f6 d4c6
info depth 10 seldepth 13 multipv 1 score cp 35 nodes 28785 nps 564411 tbhits 0 time 51 pv e2e4 e7e6 d2d4 d7d5 e4d5 e6d5 g1f3 f8d6 f1d3 g8f6 e1g1 e8g8
info depth 11 seldepth 14 multipv 1 score cp 50 nodes 34551 nps 575850 tbhits 0 time 60 pv e2e4 e7e5 g1f3 b8c6 d2d4 e5d4 f3d4 g8f6 b1c3 f8b4
info depth 12 seldepth 14 multipv 1 score cp 50 nodes 55039 nps 534359 tbhits 0 time 103 pv e2e4 e7e5 g1f3 b8c6 d2d4 e5d4 f3d4 g8f6 b1c3 f8b4
bestmove e2e4 ponder c7c5`
)

const (
	multipvTestingString = `
info depth 19 currmove f2f3 currmovenumber 13
info depth 19 currmove g1h3 currmovenumber 14
info depth 19 currmove f2f4 currmovenumber 15
info depth 19 currmove b1a3 currmovenumber 16
info depth 19 currmove a2a4 currmovenumber 17
info depth 19 currmove h2h4 currmovenumber 18
info depth 19 currmove b2b4 currmovenumber 19
info depth 19 currmove g2g4 currmovenumber 20
info depth 19 currmove d2d3 currmovenumber 10
info depth 19 currmove b2b3 currmovenumber 11

info depth 19 seldepth 23 multipv 1 score cp 47 nodes 5953192 nps 550813 hashfull 953 tbhits 0 time 10808 pv e2e4 c7c5 g1f3 b8c6 d2d4 c5d4 f3d4 g7g6 c2c4 g8f6 b1c3 d7d6 f2f3 c6d4 d1d4 f8g7 c1e3 c8e6 f1e2 e8g8 d4d2 a8c8 c3d5
info depth 19 seldepth 23 multipv 2 score cp 31 nodes 5953192 nps 550813 hashfull 953 tbhits 0 time 10808 pv g1f3 g8f6 c2c4 e7e6 b1c3 d7d5 d2d4 f8b4 c1g5 b8d7 c4d5 e6d5 d1c2 h7h6 g5f6 b4c3 b2c3 d7f6 e2e3 e8g8 f1d3
info depth 19 seldepth 22 multipv 3 score cp 30 nodes 5953192 nps 550813 hashfull 953 tbhits 0 time 10808 pv c2c4 g8f6 g1f3 e7e6 b1c3 d7d5 d2d4 f8b4 c1g5 b8d7 c4d5 e6d5 d1c2 h7h6 g5f6 d8f6 a2a3 b4c3 c2c3 c7c6 e2e3 e8g8
info depth 19 seldepth 22 multipv 4 score cp 29 nodes 5953192 nps 550813 hashfull 953 tbhits 0 time 10808 pv d2d4 g8f6 c2c4 e7e6 g1f3 d7d5 b1c3 f8b4 c1g5 b8d7 c4d5 e6d5 d1c2 h7h6 g5f6 b4c3 b2c3 d7f6 e2e3 e8g8 f1d3 c8g4
info depth 19 seldepth 24 multipv 5 score cp 23 nodes 5953192 nps 550813 hashfull 953 tbhits 0 time 10808 pv g2g3 d7d5 g1f3 c7c5 f1g2 g8f6 e1g1 e7e6 d2d4 c5d4 f3d4 e6e5 d4b3 c8e6 c2c4 b8c6 c4d5 f6d5 b1c3 d5c3 b2c3 a8c8
info depth 19 seldepth 21 multipv 6 score cp 13 nodes 5953192 nps 550813 hashfull 953 tbhits 0 time 10808 pv e2e3 d7d5 d2d4 g8f6 g1f3 e7e6 f1d3 c7c5 b2b3 b7b6 e1g1 f8d6 b1d2 c5d4 e3d4 b8c6 a2a3 e8g8 c1b2
info depth 19 seldepth 22 multipv 7 score cp 4 nodes 5953192 nps 550813 hashfull 953 tbhits 0 time 10808 pv c2c3 g8f6 d2d4 c7c5 g1f3 d7d5 d4c5 e7e6 c1e3 f8e7 e3d4 d8c7 e2e3 e7c5 d4c5 c7c5 f1d3 e8g8 e1g1
info depth 19 seldepth 25 multipv 8 score cp 0 nodes 5953192 nps 550813 hashfull 953 tbhits 0 time 10808 pv a2a3 e7e5 e2e4 g8f6 b1c3 d7d5 e4d5 f6d5 d1h5 d5f6 h5e5 f8e7 c3b5 b8a6 e5d4 c8d7 b5c3 e8g8 f1a6 b7a6 g1f3 c7c5 d4f4 f6h5 f4e5
info depth 19 seldepth 28 multipv 9 score cp 0 nodes 5953192 nps 550813 hashfull 953 tbhits 0 time 10808 pv b1c3 d7d5 d2d4 g8f6 c1f4 c7c5 e2e3 c5d4 e3d4 a7a6 f1d3 b8c6 c3e2 d8b6 c2c3 b6b2 g1f3 c8g4 a2a4 g4f3
info depth 19 seldepth 26 multipv 10 score cp -22 nodes 5953192 nps 550813 hashfull 953 tbhits 0 time 10808 pv d2d3 d7d5 e2e4 d5e4 d3e4 d8d1 e1d1 e7e5 c1e3 g8f6 f2f3 c8e6 b1d2 b8d7 d2c4 e8c8 d1e1 f6e8 b2b3 e8d6 c4d6 f8d6
bestmove e2e4 ponder c7c5`
)
