package chess

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

// String implements the fmt.Stringer interface
func (o Outcome) String() string {
	return string(o)
}

// A Method is the way in which the outcome occured.
type Method uint8

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
