package uci_test

import (
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
	if eng.SearchResults().BestMove.S2() != chess.E4 {
		t.Fatal("expected a different move")
	}
}

func TestLogger(t *testing.T) {
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
	t.Log(actual)
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

var (
	infoRegex = regexp.MustCompile("(?m)[\r\n]+^.*info.*$")
)

const (
	logOutput = `uci
Stockfish 14.1 by the Stockfish developers (see AUTHORS file)
id name Stockfish 14.1
id author the Stockfish developers (see AUTHORS file)
	
option name Debug Log File type string default 
option name Contempt type spin default 24 min -100 max 100
option name Analysis Contempt type combo default Both var Off var White var Black var Both
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
option name EvalFile type string default nn-82215d0fd0df.nnue
uciok
isready
readyok
setoption name UCI_Elo value 1500
ucinewgame
position fen rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
go movetime 100
info string NNUE evaluation using nn-82215d0fd0df.nnue enabled
info depth 1 seldepth 1 multipv 1 score cp 3 nodes 20 nps 20000 tbhits 0 time 1 pv d2d4
info depth 2 seldepth 2 multipv 1 score cp 50 nodes 42 nps 42000 tbhits 0 time 1 pv d2d4 a7a6
info depth 3 seldepth 3 multipv 1 score cp 8 nodes 153 nps 153000 tbhits 0 time 1 pv a2a3 e7e6 d2d4
info depth 4 seldepth 4 multipv 1 score cp 9 nodes 303 nps 303000 tbhits 0 time 1 pv c2c3 c7c5 d2d4 c5d4
info depth 5 seldepth 5 multipv 1 score cp 9 nodes 532 nps 266000 tbhits 0 time 2 pv c2c3 c7c5 d2d4 c5d4 c3d4
info depth 6 seldepth 6 multipv 1 score cp 10 nodes 1004 nps 502000 tbhits 0 time 2 pv c2c3 g8f6 d2d4 d7d5 b1d2
info depth 7 seldepth 7 multipv 1 score cp 37 nodes 1680 nps 560000 tbhits 0 time 3 pv c2c4 e7e5 e2e3 g8f6 g1f3
info depth 8 seldepth 9 multipv 1 score cp 24 nodes 4717 nps 673857 tbhits 0 time 7 pv c2c4 e7e6 d2d4 d7d5 g1f3 d5c4 e2e3 g8f6 f1c4
info depth 9 seldepth 13 multipv 1 score cp 35 nodes 7299 nps 729900 tbhits 0 time 10 pv c2c4 c7c5 b1c3 b7b6 e2e4 c8b7 d2d4
info depth 10 seldepth 12 multipv 1 score cp 44 nodes 15571 nps 778550 tbhits 0 time 20 pv d2d4 g8f6 c2c4 e7e6 g1f3 d7d5 b1c3 d5c4 d1a4 b8d7 a4c4
info depth 11 seldepth 15 multipv 1 score cp 77 nodes 32300 nps 807500 tbhits 0 time 40 pv e2e4 c7c6 d2d4 d7d5 e4e5 c8f5 g1f3 e7e6
info depth 12 seldepth 16 multipv 1 score cp 22 nodes 79712 nps 830333 tbhits 0 time 96 pv e2e4 c7c5 b1c3 e7e6 d2d4 c5d4 d1d4 b8c6 d4e3 f8b4 c1d2 g8f6 e4e5
info depth 13 seldepth 16 multipv 1 score cp 22 nodes 83419 nps 825930 tbhits 0 time 101 pv e2e4 c7c5 b1c3 e7e6 d2d4 c5d4 d1d4 b8c6 d4e3 f8b4 c1d2 g8f6 e4e5
bestmove e2e4 ponder c7c5`
)
