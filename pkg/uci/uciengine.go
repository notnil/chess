// package engine defines a generic interface for communicating with chess
// engines.
package uci

import (
	"errors"

	"time"
	chess "github.com/MelleKoning/chess/pkg/chess"
)

var (
	// ErrTimeout indicates a timeout in communicating with the engine.
	ErrTimeout = errors.New("timeout in engine communication")
	// ErrExited indicates that the engine was closed.
	ErrExited = errors.New("engine was closed")
)

// Engine provides a generic interface to a running chess engine.
type UciEngine interface {
	// SetPosition sets the position to search.
	SetPosition(position *chess.Position)

	// Search starts an infinite search of the position; that is, until
	// Stop is called. During the search, Info's will be sent on the
	// channel that is returned. The last Info on the channel, and only the
	// last, will either return a non-nil Info.Err or an Info.BestMove,
	// after which the channel is closed.
	Search() <-chan UciEngineInfo

	// SearchDepth is like Search but tells the engine to stop at a certain
	// depth.
	SearchDepth(depth int) <-chan UciEngineInfo

	// SearchTime is like Search but tells the engine to stop after the
	// given time.
	SearchTime(t time.Duration) <-chan UciEngineInfo

	// SearchClock is like Search but informs the engine of the time
	// controls of the game and lets the engine decide how much time to
	// use. movesToGo is the number of moves to the next time control.
	SearchClock(wtime, btime, winc, binc time.Duration, movesToGo int) <-chan UciEngineInfo

	// Stop stops a search started by one of the SearchXXX functions.
	Stop()

	// Quit quits the engine process.
	Quit()

	// Ping pings the engine process to check that it is still responding.
	Ping() error

	// Options returns the settable options of the engine.
	Options() map[string]UciEngineOption
}

// Pv holds the details of a principal variation from an engine search.
type Pv struct {
	Moves      []chess.Move // principal variation moves
	Score      int    // centipawns or moves-to-mate; positive is good for white
	Mate       bool   // if yes then Score is moves-to-mate
	Upperbound bool   // Score is a upperbound
	Lowerbound bool   // Score is a lowerbound
	Rank       int    // 0-based rank of the pv in a MultiPV search
}

// Stats holds statistics from an engine search.
type Stats struct {
	Depth    int           // depth in plies
	SelDepth int           // selective depth
	Nodes    int           // number of nodes searched so far
	Time     time.Duration // amount of time searched so far
}

// UciEngineInfo represents a generic information "event" sent over an Info channel
// while an engine search is in progress.
type UciEngineInfo interface {
	// Err returns an error if one occured. This should be the first thing
	// to check, before calling the other methods.
	Err() error

	// BestMove returns the best move found by the engine. It returns !ok
	// if no best move has been found yet.
	BestMove() (move chess.Move, ok bool)

	// Pv returns the principal variation of this Info. It can be nil if no
	// pv information is available.
	Pv() *Pv

	// Stats returns statistics of the search so far. Any of the Stats
	// fields can be zero, meaning the information is not present (or
	// actually zero).
	Stats() *Stats
}

// Engine options

// UciEngineOption represents a generic settable engine option.
type UciEngineOption interface {
	String() string        // current value
	StringDefault() string // default value
	Set(value string)      // change the value
}

// StringOption represents string option.
type StringOption interface {
	UciEngineOption
}

// IntOption represents a number option, possibly with a limited range.
type IntOption interface {
	Int() int     // current value
	SetInt(int)   // change value
	Default() int // default value
	Min() int     // minimum value
	Max() int     // maximum value (0 = no maximum)
}

// BoolOption represents a togglable option.
type BoolOption interface {
	Bool() bool    // current value
	SetBool(bool)  // change value
	Default() bool // default value
}
