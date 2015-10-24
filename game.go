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

// A Outcome is the result of a game.
type Outcome string

const (
	// NoOutcome indicates that a game is in progress or ended without a result.
	NoOutcome Outcome = "*"
	// WhiteWon indicates that white won the game.
	WhiteWon Outcome = "1-0"
	// BlackWon indicates that black won the game.
	BlackWon Outcome = "0-1"
	// Draw indicates that game was a draw.
	Draw Outcome = "1/2-1/2"
)

// A Method is the way in which the outcome occured.
type Method int

const (
	// NoMethod indicates that an outcome hasn't occured or that the method can't be determined.
	NoMethod Method = iota
	// Checkmate indicates that the game was won by a playing being checkmated.
	Checkmate
	// Resignation indicates that the game was won by player resigning.
	Resignation
	// DrawOffer indicates that the game was drawn by player agreeing to a draw offer.
	DrawOffer
	// Stalemate indicates that the game was drawn by player being stalemated.
	Stalemate
	// ThreefoldRepetition indicates that the game was drawn when the game
	// state was repeated three times and a player requested a draw.
	ThreefoldRepetition
	// FivefoldRepetition indicates that the game was automatically drawn
	// by the game state being repeated five times.
	FivefoldRepetition
	// FiftyMoveRule indicates that the game was drawn by the half
	// move clock being fifty or greater when a player requested a draw.
	FiftyMoveRule
	// SeventyFiveMoveRule indicates that the game was automatically drawn
	// when the half move clock was seventy five or greater.
	SeventyFiveMoveRule
	// InsufficientMaterial indicates that the game was automatically drawn
	// because there was insufficent material for checkmate.
	InsufficientMaterial
)

// TagPair represents metadata in a key value pairing.
type TagPair struct {
	Key   string
	Value string
}

// A Game represents a single chess game.
type Game struct {
	tagPairs []*TagPair
	moves    []*Move
	state    *GameState
	outcome  Outcome
	method   Method
}

// PGN takes a reader and returns a function that updates
// the game to reflect the PGN data.  The returned function
// is designed to be used in the NewGame constructor.  An
// error is returned if there is a problem parsing the PGN
// data.
func PGN(r io.Reader) (func(*Game), error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	game, err := decodePGN(string(b))
	if err != nil {
		return nil, err
	}
	return func(g *Game) {
		g.copy(game)
	}, nil
}

// FEN takes a reader and returns a function that updates
// the game to reflect the FEN data.  Since FEN doesn't include
// prior moves, the move list will be empty.  The returned
// function is designed to be used in the NewGame constructor.
// An error is returned if there is a problem parsing the FEN data.
func FEN(r io.Reader) (func(*Game), error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	state, err := decodeFEN(string(b))
	if err != nil {
		return nil, err
	}
	return func(g *Game) {
		g.updateState(state)
	}, nil
}

// TagPairs returns a function that sets the tag pairs
// to the given value.  The returned function is designed
// to be used in the NewGame constructor.
func TagPairs(tagPairs []*TagPair) func(*Game) {
	return func(g *Game) {
		g.tagPairs = append([]*TagPair(nil), tagPairs...)
	}
}

// NewGame defaults to returning a game in the standard
// opening position.  Options can be given to change
// the game's initial state.
func NewGame(options ...func(*Game)) *Game {
	state, _ := decodeFEN(startFEN)
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

// Move updates the game with the given move.  An error is returned
// if the move is invalid or the game has already been completed.
func (g *Game) Move(m *Move) error {
	if g.outcome != NoOutcome {
		return fmt.Errorf("chess: invalid move %s game %s by %s", m, g.Outcome(), g.Method())
	}
	if !m.isValid() {
		return fmt.Errorf("chess: invalid move %s", m)
	}
	g.moves = append(g.moves, m)
	g.updateState(m.postMoveState())
	return nil
}

// MoveSq moves the piece at s1 to s2, applies the given
// promotion, and updates the game.  An error is returned
// if the move is invalid or the game has already been completed.
func (g *Game) MoveSq(s1, s2 *Square, promo PieceType) error {
	move := &Move{
		s1:    s1,
		s2:    s2,
		promo: promo,
		state: g.state,
	}
	return g.Move(move)
}

// MoveAlg decodes the given string in algebraic notation
// and calls the Move function.  An error is returned if
// the move can't be decoded, the move is invalid, or the
// game has already been completed.
func (g *Game) MoveAlg(alg string) error {
	move, err := decodeMove(g.State(), alg)
	if err != nil {
		return err
	}
	return g.MoveSq(move.S1(), move.S2(), move.Promo())
}

// ValidMoves returns a list of valid moves in the
// current position.
func (g *Game) ValidMoves() []*Move {
	return g.state.validMoves()
}

// States returns the state history of the game.
func (g *Game) States() []*GameState {
	states := []*GameState{}
	for _, m := range g.moves {
		states = append(states, m.PreMoveState())
	}
	states = append(states, g.state)
	return states
}

// Moves returns the move history of the game.
func (g *Game) Moves() []*Move {
	return append([]*Move(nil), g.moves...)
}

// TagPairs returns the game's tag pairs.
func (g *Game) TagPairs() []*TagPair {
	return append([]*TagPair(nil), g.tagPairs...)
}

// State returns the game's current state.
func (g *Game) State() *GameState {
	return g.state
}

// Outcome returns the game outcome.
func (g *Game) Outcome() Outcome {
	return g.outcome
}

// Method returns the method in which the outcome occured.
func (g *Game) Method() Method {
	return g.method
}

// TakeBack returns a copy of the game with the most recent
// n moves removed.  If n is greater than the number of moves
// or is negative then the game is set back to its initial state.
func (g *Game) TakeBack(n int) *Game {
	cp := &Game{}
	cp.copy(g)
	if len(cp.moves) == 0 {
		return cp
	}
	var state *GameState
	if len(cp.moves) < n || n < 0 {
		state = cp.moves[0].PreMoveState()
		cp.moves = []*Move{}
	} else {
		i := len(cp.moves) - n
		state = cp.moves[i].PreMoveState()
		cp.moves = g.moves[:i]
	}
	cp.updateState(state)
	return cp
}

// FEN returns the FEN notation of the current state.
func (g *Game) FEN() string {
	return g.State().String()
}

// String implements the fmt.Stringer interface and returns
// the game's PGN.
func (g *Game) String() string {
	return encodePGN(g)
}

// MarshalText implements the encoding.TextMarshaler interface and
// encodes the game's PGN.
func (g *Game) MarshalText() (text []byte, err error) {
	return []byte(encodePGN(g)), nil
}

// UnmarshalText implements the encoding.TextUnarshaler interface and
// assumes the data is in the PGN format.
func (g *Game) UnmarshalText(text []byte) error {
	game, err := decodePGN(string(text))
	if err != nil {
		return err
	}
	g.copy(game)
	return nil
}

// Draw attempts to draw the game by the given method.  If the
// method is valid, then the game is updated to a draw by that
// method.  If the method isn't valid then an error is returned.
func (g *Game) Draw(method Method) error {
	switch method {
	case ThreefoldRepetition:
		if g.numOfRepitions() < 3 {
			return errors.New("chess: draw by ThreefoldRepetition requires at least three repetitions of the current board state")
		}
	case FiftyMoveRule:
		if g.state.halfMoveClock < 50 {
			return fmt.Errorf("chess: draw by FiftyMoveRule requires the half move clock to be at 50 or greater but is %d", g.state.halfMoveClock)
		}
	case DrawOffer:
	default:
		return fmt.Errorf("chess: unsupported draw method %s", method)
	}
	g.outcome = Draw
	g.method = method
	return nil
}

// Resign resigns the game for the given color.  If the game has
// already been completed then the game is not updated.
func (g *Game) Resign(color Color) {
	if g.outcome != NoOutcome || color == NoColor {
		return
	}
	if color == White {
		g.outcome = BlackWon
	} else {
		g.outcome = WhiteWon
	}
	g.method = Resignation
}

// EligibleDraws returns the draw methods that eligible for game's
// Draw method.
func (g *Game) EligibleDraws() []Method {
	draws := []Method{DrawOffer}
	if g.numOfRepitions() >= 3 {
		draws = append(draws, ThreefoldRepetition)
	}
	if g.state.halfMoveClock < 50 {
		draws = append(draws, FiftyMoveRule)
	}
	return draws
}

func (g *Game) copy(game *Game) {
	g.tagPairs = game.tagPairs
	g.moves = game.moves
	g.state = game.state
	g.outcome = game.outcome
	g.method = game.method
}

func (g *Game) updateState(state *GameState) {
	g.state = state
	outcome, method := state.getOutcome()
	g.outcome = outcome
	g.method = method

	// five fold rep creates automatic draw
	if g.numOfRepitions() >= 5 {
		g.outcome = Draw
		g.method = FivefoldRepetition
	}

	// 75 move rule creates automatic draw
	if g.state.halfMoveClock >= 75 && g.method != Checkmate {
		g.outcome = Draw
		g.method = SeventyFiveMoveRule
	}

	// insufficent material creates automatic draw
	if !g.state.board.hasSufficientMaterial() {
		g.outcome = Draw
		g.method = InsufficientMaterial
	}
}

func (g *Game) numOfRepitions() int {
	count := 0
	for _, gs := range g.States() {
		if g.state.samePosition(gs) {
			count++
		}
	}
	return count
}
