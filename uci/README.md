# uci

## Introduction

**uci** is a client package for [Universal Chess Interface protocol](http://wbec-ridderkerk.nl/html/UCIProtocol.html) (UCI) compliant engines (such as [Stockfish](https://stockfishchess.org/)).

## Installation

**uci** requires an engine to be useful.  On macOS you can install Stockfish with:
```
brew install stockfish
```
All installation methods are listed on [Stockfish's Download Page](https://stockfishchess.org/download/)

## Supported Commands

Engine's Run method takes commands that implement the **Cmd** interface.  Here is a supported commands:

| Command  | Type  |Description |
| ------------- | ------------- | ------------- |
| uci  | CmdUCI  | tell engine to use the uci  |
| isready  | CmdIsReady  | this command can be used to wait for the engine to be ready again or to ping the engine to find out if it is still alive|
| setoption  | CmdSetOption  | this is sent to the engine when the user wants to change the internal parameters  |
| ucinewgame  | CmdNewGame  | this is sent to the engine when the next search (started with "position" and "go") will be from a different game.  |
| position  | CmdPosition  | set up the position described in fenstring on the internal board and play the moves on the internal chess board.  |
| go  | CmdGo  | start calculating on the current position set up with the "position" command.  |
| stop  | CmdStop  | stop calculating as soon as possible  |
| ponderhit  | CmdPonderHit  | This will be sent if the engine was told to ponder on the same move the user has played  |
| quit  | CmdQuit  | quit the program as soon as possible  |

## Example Stockfish v. Stockfish

```go
package main

import (
	"fmt"
	"time"

	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
)

func main() {
	// set up engine to use stockfish exe
	eng, err := uci.New("stockfish")
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
	// Output: 
	// 1.c4 c5 2.Nf3 e6 3.Nc3 Nc6 4.d4 cxd4 5.Nxd4 Nf6 6.a3 d5 7.cxd5 exd5 8.Bf4 Bc5 9.Ndb5 O-O 10.Nc7 d4 11.Na4 Be7 12.Nxa8 Bf5 13.g3 Qd5 14.f3 Rxa8 15.Bg2 Rd8 16.b4 Qe6 17.Nc5 Bxc5 18.bxc5 Nd5 19.O-O Nc3 20.Qd2 Nxe2+ 21.Kh1 d3 22.Bd6 Qd7 23.Rab1 h6 24.a4 Re8 25.g4 Bg6 26.a5 Ncd4 27.Qb4 Qe6 28.Qxb7 Nc2 29.Qxa7 Ne3 30.Rb8 Nxf1 31.Qb6 d2 32.Rxe8+ Qxe8 33.Qb3 Ne3 34.h3 Bc2 35.Qxc2 Nxc2 36.Kh2 d1=Q 37.h4 Qg1+ 38.Kh3 Ne1 39.h5 Qxg2+ 40.Kh4 Nxf3#  0-1
}
```