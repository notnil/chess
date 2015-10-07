package chess

import (
	"errors"
	"fmt"
)

const (
	startFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)

type Status int

const (
	InProgress Status = iota + 1
	Complete
)

type Outcome int

const (
	NoOutcome Outcome = iota
	WhiteWon
	BlackWon
	Draw
)

type Method int

const (
	NoMethod Method = iota
	Checkmate
	Resignation
	DrawOffer
	Stalemate
)

type Game struct {
	moves   []*Move
	state   *GameState
	status  Status
	outcome Outcome
	method  Method
}

func NewGame() *Game {
	state, _ := fen(startFEN)
	return &Game{
		moves:   []*Move{},
		state:   state,
		status:  InProgress,
		outcome: NoOutcome,
		method:  NoMethod,
	}
}

func (g *Game) Move(s1, s2 *Square, promo *PieceType) error {
	if g.status == Complete {
		return errors.New("chess: invalid move game complete")
	}
	move := &Move{
		s1:    s1,
		s2:    s2,
		promo: promo,
		state: g.state,
	}
	if !move.isValid() {
		return fmt.Errorf("chess: invalid move %s", move)
	}
	g.moves = append(g.moves, move)
	g.state = move.postMoveState()
	outcome, method := g.state.getOutcome()
	g.outcome = outcome
	g.method = method
	return nil
}

func (g *Game) ValidMoves() []*Move {
	return g.state.validMoves()
}

func (g *Game) Moves() []*Move {
	return append([]*Move(nil), g.moves...)
}

func (g *Game) GameState() *GameState {
	return g.state
}

func (g *Game) Status() Status {
	return g.status
}

func (g *Game) Outcome() Outcome {
	return g.outcome
}

func (g *Game) Method() Method {
	return g.method
}
