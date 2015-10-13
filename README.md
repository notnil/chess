# chess

package chess is a go library designed to accomplish the following:
- chess game / turn management
- move validation
- PGN encoding / decoding
- FEN encoding / decoding

## Usage

Using API
```go
game := chess.NewGame()
game.Move(chess.E2,chess.E4,chess.NoPromo)
```

Using Algebraic Notation
```go
game := chess.NewGame()
game.MoveAlg("e4")
```

Random Game
```go
package main

import (
	"fmt"

	"github.com/loganjspears/chess"
)

func main() {
    game := chess.NewGame()
    fmt.Println(game.State().Board().Draw())
    for g.Outcome() == chess.NoOutcome {
        moves := game.ValidMoves()
        move := moves[rand.Intn(len(moves))]
        game.Move(move.S1(),move.S2(),move.Promo())
        fmt.Println(game.State().Board().Draw())
    }
    fmt.Printf("Game completed. %s by %s.\n", game.Outcome(), game.Method())
    fmt.Println(game.String())    
}
```

## chesscli

chesscli is a fully functioning client of the chess package.  It takes moves from algebraic notation and updates the state of the game.  I use it to record live games and test the package.
