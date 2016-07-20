# chess
[![GoDoc](https://godoc.org/github.com/notnil/chess?status.svg)](https://godoc.org/github.com/notnil/chess)
[![Build Status](https://drone.io/github.com/notnil/chess/status.png)](https://drone.io/github.com/notnil/chess/latest)
[![Coverage Status](https://coveralls.io/repos/notnil/chess/badge.svg?branch=master&service=github)](https://coveralls.io/github/notnil/chess?branch=master)
[![Go Report Card](https://goreportcard.com/badge/notnil/chess)](https://goreportcard.com/report/notnil/chess)

package chess is a go library designed to accomplish the following:
- chess game / turn management
- move validation
- PGN encoding / decoding
- FEN encoding / decoding

## Usage

Using Moves
```go
game := chess.NewGame()
moves := game.ValidMoves()
game.Move(moves[0])
```

Using Algebraic Notation
```go
game := chess.NewGame()
game.MoveAlg("e4")
```

Using PGN
```go
pgn, _ := chess.PGN(pgnReader)
game := chess.NewGame(pgn)
```

Using FEN
```go
fen, _ := chess.FEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
game := chess.NewGame(fen)
```

Random Game
```go
package main

import (
	"fmt"

	"github.com/notnil/chess"
)

func main() {
    game := chess.NewGame()
	// generate moves until game is over
    for game.Outcome() == chess.NoOutcome {
		// select a random move
        moves := game.ValidMoves()
        move := moves[rand.Intn(len(moves))]
		game.Move(move)
    }
	// print outcome and game PGN
	fmt.Println(game.Position().Board().Draw())
	fmt.Printf("Game completed. %s by %s.\n", game.Outcome(), game.Method())
    fmt.Println(game.String())    
}
```
