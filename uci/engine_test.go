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
	setEloOpt := uci.CmdSetOption{Name: "UCI_Elo", Value: "1500"}
	setWdlOpt := uci.CmdSetOption{Name: "UCI_ShowWDL", Value: "true"}
	setPos := uci.CmdPosition{Position: chess.StartingPosition()}
	setGo := uci.CmdGo{MoveTime: 5 * time.Second}
	if err := eng.Run(uci.CmdUCI, uci.CmdIsReady, setEloOpt, setWdlOpt, uci.CmdUCINewGame, setPos, setGo); err != nil {
		t.Fatal(err)
	}
	if eng.SearchResults().BestMove.S2() != chess.D4 {
		t.Fatal("expected a different move")
	}
	// these wdl values are sensitive to stockfish version, KNPS, cpu, and
	// search time. current values were set via stockfish version 17 on
	// AMD Ryzen 7 5800X3D
	if eng.SearchResults().Info.Score.Win != 86 {
		t.Fatalf("expected a different win probability: %v", eng.SearchResults().Info.Score.Win)
	}
	if eng.SearchResults().Info.Score.Draw != 894 {
		t.Fatalf("expected a different draw probability: %v", eng.SearchResults().Info.Score.Draw)
	}
	if eng.SearchResults().Info.Score.Loss != 20 {
		t.Fatalf("expected a different draw probability: %v", eng.SearchResults().Info.Score.Loss)
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
