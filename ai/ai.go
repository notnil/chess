package ai

import (
	"math/rand"
	"time"

	"github.com/loganjspears/chess"
)

type AI interface {
	Move(gs *chess.GameState) *chess.Move
}

type Random struct{}

func (r Random) Move(gs *chess.GameState) *chess.Move {
	moves := gs.ValidMoves()
	rand.Seed(time.Now().UnixNano())
	return moves[rand.Intn(len(moves))]
}
