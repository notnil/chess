package chess

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidMove = errors.New("chess: invalid move")
	ErrOutOfTurn   = errors.New("chess: out of turn")
)

type Game struct {
	moves []*move
	board Board
	turn  color
}

type config struct{}

func NewGame(options ...func(*config)) *Game {
	c := &config{}
	for _, o := range options {
		o(c)
	}
	return &Game{
		moves: []*move{},
		board: newBoard(),
		turn:  white,
	}
}

func (g *Game) Move(s1 *Square, s2 *Square, promo *Piece) error {
	if g.board.piece(s1).color() != g.turn {
		return ErrOutOfTurn
	}
	valid := squareSlice(g.board.validMoves(s1)).has(s2)
	if !valid {
		return ErrInvalidMove
	}
	g.makeMove(s1, s2, promo)
	return nil
}

func (g *Game) makeMove(s1 *Square, s2 *Square, promo *Piece) {
	m := &move{
		s1:    s1,
		s2:    s2,
		promo: promo,
	}
	g.moves = append(g.moves, m)
	g.board.move(m)
	g.turn = g.turn.other()
}

type move struct {
	s1    *Square
	s2    *Square
	promo *Piece
}

func (m *move) String() string {
	return fmt.Sprintf("%s %s", m.s1, m.s2)
}

func (m *move) validate(b Board) error {
	return nil
}

func newBoard() Board {
	return map[*Square]Piece{
		A1: WRook, B1: WKnight, C1: WBishop, D1: WQueen, E1: WKing, F1: WBishop, G1: WKnight, H1: WRook, // white back row
		A2: WPawn, B2: WPawn, C2: WPawn, D2: WPawn, E2: WPawn, F2: WPawn, G2: WPawn, H2: WPawn, // white front row
		A7: BPawn, B7: BPawn, C7: BPawn, D7: BPawn, E7: BPawn, F7: BPawn, G7: BPawn, H7: BPawn, // black front row
		A8: BRook, B8: BKnight, C8: BBishop, D8: BQueen, E8: BKing, F8: BBishop, G8: BKnight, H8: BRook, // black back row
	}
}
