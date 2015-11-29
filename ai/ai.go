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

type Athena struct {
	Ply int
}

func (a Athena) Move(gs *chess.GameState) *chess.Move {
	rand.Seed(time.Now().UnixNano())
	return minMax(gs, a.Ply, 0)
}

func minMax(gs *chess.GameState, maxPly, ply int) *chess.Move {
	var topMove *chess.Move
	topScore := -1000.0

	for _, m := range gs.ValidMoves() {
		state := m.PostMoveState()
		if maxPly != ply {
			plyMove := minMax(state, maxPly, ply+1)
			state = plyMove.PostMoveState()
		}
		scr := score(state, gs.Turn()) + (rand.Float64() / 100)
		if scr > topScore {
			topMove = m
			topScore = scr
		}
	}
	return topMove
}

func score(gs *chess.GameState, color chess.Color) float64 {
	outcome, _ := gs.Outcome()
	switch outcome {
	case chess.WhiteWon:
		if color == chess.White {
			return 1000.0
		}
		return -1000.0
	case chess.BlackWon:
		if color == chess.Black {
			return 1000.0
		}
		return -1000.0
	case chess.Draw:
		return 0.0
	}

	total := 0.0
	for _, piece := range gs.Board() {
		score := pieceScore(gs, piece)
		if piece.Color() == color {
			total += score
		} else {
			total -= score
		}
	}
	return total
}

func pieceScore(gs *chess.GameState, piece *chess.Piece) float64 {
	switch piece.Type() {
	case chess.King:
		return 200.0
	case chess.Queen:
		return 9.0
	case chess.Rook:
		return 5.0
	case chess.Bishop:
		return 3.1
	case chess.Knight:
		return 3.0
	case chess.Pawn:
		return 1.0
	}
	return 0.0
}
