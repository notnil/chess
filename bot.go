package chess

import (
	"sort"

	"github.com/notnil/chess"
	"github.com/notnil/opening"
)

/*
TODO
CPU Profile
*/

type Bot struct {
	Debug   bool
	UseBook bool
}

func (b *Bot) Move(game *chess.Game) (*chess.Move, float64) {
	if game == nil {
		return nil, 0.0
	}
	var move *chess.Move
	if b.UseBook {
		move = openingMove(game)
		if move != nil {
			return move, 0.0
		}
	}
	pos := game.Position()
	turn := pos.Turn()
	moves := pos.ValidMoves()
	sort.Sort(byMoveImportance(moves))

	type moveScore struct {
		move  *chess.Move
		score float64
	}
	ch := make(chan moveScore)
	for _, m := range moves {
		go func(m *chess.Move) {
			newPos := pos.Update(m)
			score := alphaBetaMin(turn, newPos, -10000.0, 10000.0, 5)
			ch <- moveScore{m, score}
		}(m)
	}
	bestScore := -10000.0
	for i := 0; i < len(moves); i++ {
		moveScr := <-ch
		if moveScr.score > bestScore {
			bestScore = moveScr.score
			move = moveScr.move
		}
	}
	return move, bestScore
}

func alphaBetaMax(maxColor chess.Color, pos *chess.Position, alpha, beta float64, depthleft int) float64 {
	if depthleft == 0 {
		return score(maxColor, pos)
	}
	moves := pos.ValidMoves()
	sort.Sort(byMoveImportance(moves))
	for _, m := range moves {
		newPos := pos.Update(m)
		score := alphaBetaMin(maxColor, newPos, alpha, beta, depthleft-1)
		if score >= beta {
			return beta
		}
		if score > alpha {
			alpha = score
		}
	}
	return alpha
}

func alphaBetaMin(maxColor chess.Color, pos *chess.Position, alpha, beta float64, depthleft int) float64 {
	if depthleft == 0 {
		return score(maxColor, pos)
	}
	for _, m := range pos.ValidMoves() {
		newPos := pos.Update(m)
		score := alphaBetaMax(maxColor, newPos, alpha, beta, depthleft-1)
		if score <= alpha {
			return alpha
		}
		if score < beta {
			beta = score
		}
	}
	return beta
}

var (
	scoreCache = map[string]float64{}
)

func score(maxColor chess.Color, pos *chess.Position) float64 {
	turn := pos.Turn()
	status := pos.Status()
	if status == chess.Stalemate {
		return 0.0
	} else if status == chess.Checkmate {
		if turn == maxColor {
			return 1000.0
		}
		return -1000.0
	}
	total := 0.0
	for _, piece := range pos.Board().SquareMap() {
		score := pieceScore(pos, piece)
		if piece.Color() == maxColor {
			total += score
		} else {
			total -= score
		}
	}
	return total
}

func pieceScore(pos *chess.Position, piece chess.Piece) float64 {
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

func openingMove(game *chess.Game) *chess.Move {
	prevMoves := game.Moves()
	moveIndex := len(prevMoves)
	opennings := opening.Possible(game.Moves())
	sort.Sort(byOpeningLength(opennings))
	for _, op := range opennings {
		opGame := op.Game()
		opMoves := opGame.Moves()
		opMoveIndex := len(opGame.Moves())
		if opMoveIndex > moveIndex {
			return opMoves[moveIndex]
		}
	}
	return nil
}

type byOpeningLength []*opening.Opening

func (a byOpeningLength) Len() int           { return len(a) }
func (a byOpeningLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byOpeningLength) Less(i, j int) bool { return len(a[i].PGN()) > len(a[j].PGN()) }

type byMoveImportance []*chess.Move

func (a byMoveImportance) Len() int      { return len(a) }
func (a byMoveImportance) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byMoveImportance) Less(i, j int) bool {
	if a[i].HasTag(chess.Check) && !a[j].HasTag(chess.Check) {
		return true
	}
	if a[i].HasTag(chess.Capture) && !a[j].HasTag(chess.Capture) {
		return true
	}
	return false
}
