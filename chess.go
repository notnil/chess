package chess

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidMove          = errors.New("chess: attempted invalid move")
	ErrOutOfTurn            = errors.New("chess: attempted to move out of turn")
	ErrGameAlreadyCompleted = errors.New("chess: attempted move after checkmate")
	ErrInvalidPromotion     = errors.New("chess: attempted invalid pawn promotion")
	ErrInvalidCastling      = errors.New("chess: attempted invalid castling")
)

type Status int

const (
	Ongoing Status = iota
	Stalemate
	BlackCheckmated
	WhiteCheckmated
	BlackResigned
	WhiteResigned
	Draw
)

type Game struct {
	moves  []*move
	board  Board
	turn   color
	status Status
}

type config struct{}

func NewGame(options ...func(*config)) *Game {
	c := &config{}
	for _, o := range options {
		o(c)
	}
	return &Game{
		moves:  []*move{},
		board:  newBoard(),
		turn:   white,
		status: Ongoing,
	}
}

func (g *Game) Resign() {
	switch g.turn {
	case white:
		g.status = WhiteResigned
	case black:
		g.status = BlackResigned
	}
}

func (g *Game) Move(s1 *Square, s2 *Square, promo *Piece) error {
	if g.Finished() {
		return ErrGameAlreadyCompleted
	}
	if g.board.piece(s1).color() != g.turn {
		return ErrOutOfTurn
	}
	if err := g.makeMove(s1, s2, promo); err != nil {
		return err
	}
	return nil
}

func (g *Game) Finished() bool {
	return g.status != Ongoing
}

func (g *Game) makeMove(s1 *Square, s2 *Square, promo *Piece) error {
	m := &move{s1: s1, s2: s2, promo: promo}
	valid := squareSlice(g.board.validMoves(s1)).has(s2)
	if !valid {
		if g.board.isCastling(m) {
			if err := g.checkCastling(m); err != nil {
				return err
			}
		} else {
			return ErrInvalidMove
		}
	}
	cp := g.board.copy()
	cp.move(m)
	if cp.inCheck(g.turn) {
		return ErrInvalidMove
	}
	if err := g.checkPromotion(m); err != nil {
		return err
	}
	g.moves = append(g.moves, m)
	g.board.move(m)
	g.turn = g.turn.other()
	g.status = g.calcStatus()
	return nil
}

func (g *Game) checkPromotion(m *move) error {
	promoRank := rank1
	if g.turn == white {
		promoRank = rank8
	}

	pawnMoving := g.board.piece(m.s1).pieceType() == pawn
	needsPromo := pawnMoving && promoRank == m.s2.rank
	hasPromo := m.promo != nil && g.turn == m.promo.color() && m.promo.pieceType().isPromotable()
	if needsPromo != hasPromo {
		return ErrInvalidPromotion
	}
	return nil
}

func (g *Game) checkCastling(m *move) error {
	if !g.board.isCastling(m) {
		return nil
	}

	isAttacked := func(squares ...*Square) bool {
		for _, s := range squares {
			if g.board.isSquareAttacked(g.turn, s) {
				return true
			}
		}
		return false
	}

	backRow := [8]*Square{A1, B1, C1, D1, E1, F1, G1, H1}
	if g.turn == black {
		backRow = [8]*Square{A8, B8, C8, D8, E8, F8, G8, H8}
	}
	kingSide := m.s1 == backRow[4] && m.s2 == backRow[6]
	queenSide := m.s1 == backRow[4] && m.s2 == backRow[2]
	kingMoved := g.hasPieceMoved(backRow[4])
	kingRookMoved := g.hasPieceMoved(backRow[7])
	kingSideOccupied := !g.board.emptyBetween(backRow[4], backRow[7])
	kingSideAttacked := isAttacked(backRow[4], backRow[5], backRow[6])
	queenRookMoved := g.hasPieceMoved(backRow[0])
	queenSideOccupied := !g.board.emptyBetween(backRow[0], backRow[4])
	queenSideAttacked := isAttacked(backRow[2], backRow[3], backRow[4])
	if kingSide && (kingMoved || kingRookMoved || kingSideOccupied || kingSideAttacked) {
		return ErrInvalidCastling
	}
	if queenSide && (kingMoved || queenRookMoved || queenSideOccupied || queenSideAttacked) {
		return ErrInvalidCastling
	}
	return nil
}

func (g *Game) calcStatus() Status {
	if g.status != Ongoing {
		return g.status
	}
	if g.board.inCheckmate(g.turn) {
		switch g.turn {
		case white:
			return WhiteCheckmated
		case black:
			return BlackCheckmated
		}
	}
	if g.board.inStalemate(g.turn) {
		return Stalemate
	}
	return g.status
}

func (g *Game) hasPieceMoved(s *Square) bool {
	for _, m := range g.moves {
		if m.s1 == s || m.s2 == s {
			return true
		}
	}
	return false
}

type move struct {
	s1    *Square
	s2    *Square
	promo *Piece
}

func (m *move) String() string {
	return fmt.Sprintf("%s %s", m.s1, m.s2)
}

func newBoard() Board {
	return map[*Square]*Piece{
		A1: WRook, B1: WKnight, C1: WBishop, D1: WQueen, E1: WKing, F1: WBishop, G1: WKnight, H1: WRook, // white back row
		A2: WPawn, B2: WPawn, C2: WPawn, D2: WPawn, E2: WPawn, F2: WPawn, G2: WPawn, H2: WPawn, // white front row
		A7: BPawn, B7: BPawn, C7: BPawn, D7: BPawn, E7: BPawn, F7: BPawn, G7: BPawn, H7: BPawn, // black front row
		A8: BRook, B8: BKnight, C8: BBishop, D8: BQueen, E8: BKing, F8: BBishop, G8: BKnight, H8: BRook, // black back row
	}
}
