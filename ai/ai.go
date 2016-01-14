package ai

import (
	"math/rand"
	"time"

	"github.com/loganjspears/chess/chess"
)

type AI interface {
	Move(pos *chess.Position) *chess.Move
}

type Athena struct {
	Ply      int
	scoreMap map[[16]byte]float64
}

func New(ply int) *Athena {
	return &Athena{
		Ply:      ply,
		scoreMap: map[[16]byte]float64{},
	}
}

func (a *Athena) Move(gs *chess.GameState) *chess.Move {
	rand.Seed(time.Now().UnixNano())
	move, _ := a.minMax(gs, gs.Turn(), a.Ply, 0)
	return move
}

func (a *Athena) PositionsEvaluated() int {
	return len(a.scoreMap)
}

func (a *Athena) minMax(pos *chess.Position, c chess.Color, maxPly, ply int) (*chess.Move, float64) {
	var topMove *chess.Move
	topScore := -1000.0
	for _, m := range pos.ValidMoves() {
		postPos := pos.Update(m)
		if maxPly != ply {
			plyMove, _ := a.minMax(postPos, c.Other(), maxPly, ply+1)
			if plyMove != nil {
				postPos = postPos.Update(plyMove)
			}
		}
		scr := a.score(postPos) + (rand.Float64() / 100)
		if c == chess.Black {
			scr *= -1.0
		}
		if scr > topScore {
			topMove = m
			topScore = scr
		}
	}
	return topMove, topScore
}

func (a *Athena) score(pos *chess.Position) float64 {
	// hash := pos.Hash()
	// if score, in := a.scoreMap[hash]; in {
	// 	return score
	// }
	// outcome, _ := pos.Outcome()
	// switch outcome {
	// case chess.WhiteWon:
	// 	return 1000.0
	// case chess.BlackWon:
	// 	return -1000.0
	// case chess.Draw:
	// 	return 0.0
	// }

	total := 0.0
	for _, piece := range gs.Board().Pieces() {
		score := pieceScore(gs, piece)
		if piece.Color() == chess.White {
			total += score
		} else {
			total -= score
		}
	}
	a.scoreMap[hash] = total
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
