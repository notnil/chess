## Introduction

ucichessplayer is a go library which provides a common uci engine on top of the notnil/chess library, to support uci communication to chess engines via golang.

## Usage

There is no actual code written to setup chess engines, but there is an example in the test code in file uciengine_test.go

Output can be seen from running

```
go test ./uci -v
```

#### Example interaction

Example interaction with the ucichessplayer as per the example test:

```go
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

// start from a certain position
fenStr := "5r2/pqN1Qpk1/2r3pp/2n1R3/5R2/6P1/4PPKP/8 w - - 0 1"
fen, _ := chess.FEN(fenStr)
game := chess.NewGame(fen)
e.SetPosition(game.Position())
	
// use the engine for searching 8 plies deep
for info := range e.SearchDepth(8) {
	if info.Err() != nil {
		t.Fatalf("%s", info.Err())
	}
	if m, ok := info.BestMove(); ok {
		log.Println("bestmove:", m.String()) //m.San(board))
	} else if pv := info.Pv(); pv != nil {
		log.Println("pv:", pv.ReadablePv(game.Position()))
		log.Println("stats:", info.Stats())
	} else {
		log.Println("stats:", info.Stats()) //info. .Line)
	}
}
```

The above shows chess move output from the engine in the golang debug window.
