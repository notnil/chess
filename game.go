package chess

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

const (
	startFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)

type Outcome string

const (
	NoOutcome Outcome = "*"
	WhiteWon  Outcome = "1-0"
	BlackWon  Outcome = "0-1"
	Draw      Outcome = "1/2-1/2"
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
	outcome Outcome
	method  Method
}

func FEN(r io.Reader) (func(*Game), error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	state, err := fen(string(b))
	if err != nil {
		return nil, err
	}
	return func(g *Game) {
		g.updateState(state)
	}, nil
}

func NewGame(options ...func(*Game)) *Game {
	state, _ := fen(startFEN)
	game := &Game{
		moves:   []*Move{},
		state:   state,
		outcome: NoOutcome,
		method:  NoMethod,
	}
	for _, f := range options {
		f(game)
	}
	return game
}

func (g *Game) Move(s1, s2 *Square, promo *PieceType) error {
	if g.outcome != NoOutcome {
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
	g.updateState(move.postMoveState())
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

func (g *Game) Outcome() Outcome {
	return g.outcome
}

func (g *Game) Method() Method {
	return g.method
}

func (g *Game) updateState(state *GameState) {
	g.state = state
	outcome, method := state.getOutcome()
	g.outcome = outcome
	g.method = method
}
