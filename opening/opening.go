// Package opening implements chess opening determination and exploration.
package opening

import (
	"bytes"

	"github.com/torresomarmx/chess-lib"
)

// A Opening represents a specific sequence of moves from the staring position.
type Opening struct {
	code  string
	title string
	pgn   string
	game  *chess.Game
}

// Code returns the Encyclopaedia of Chess Openings (ECO) code.
func (o *Opening) Code() string {
	return o.code
}

// Title returns the Encyclopaedia of Chess Openings (ECO) title of the opening.
func (o *Opening) Title() string {
	return o.title
}

// PGN returns the opening in PGN format.
func (o *Opening) PGN() string {
	return o.pgn
}

// Game returns the opening as a game.
func (o *Opening) Game() *chess.Game {
	if o.game == nil {
		pgn, _ := chess.PGN(bytes.NewBufferString(o.pgn))
		o.game = chess.NewGame(pgn)
	}
	return o.game
}

// Book is an opening book that returns openings for move sequences
type Book interface {
	// Find returns the most specific opening for the list of moves.  If no opening is found, Find returns nil.
	Find(moves []*chess.Move) *Opening
	// Possible returns the possible openings after the moves given.  If moves is empty or nil all openings are returned.
	Possible(moves []*chess.Move) []*Opening
}
