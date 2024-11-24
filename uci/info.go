package uci

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/notnil/chess"
)

var missingWdlErr = errors.New("uci: wdl unavailable; this is mostly likely because UCI_ShowWDL has not been set")

// SearchResults is the result from the most recent CmdGo invocation.  It includes
// data such as the following:
// info depth 21 seldepth 31 multipv 1 score cp 39 nodes 862438 nps 860716 hashfull 409 tbhits 0 time 1002 pv e2e4
// bestmove e2e4 ponder c7c5
type SearchResults struct {
	BestMove *chess.Move
	Ponder   *chess.Move
	Info     Info
}

// Info corresponds to the "info" engine output:
// the engine wants to send infos to the GUI. This should be done whenever one of the info has changed.
// The engine can send only selected infos and multiple infos can be send with one info command,
// e.g. "info currmove e2e4 currmovenumber 1" or
// 	 "info depth 12 nodes 123456 nps 100000".
// Also all infos belonging to the pv should be sent together
// e.g. "info depth 2 score cp 214 time 1242 nodes 2124 nps 34928 pv e2e4 e7e5 g1f3"
// I suggest to start sending "currmove", "currmovenumber", "currline" and "refutation" only after one second
// to avoid too much traffic.
// Additional info:
// * depth
// 	search depth in plies
// * seldepth
// 	selective search depth in plies,
// 	if the engine sends seldepth there must also a "depth" be present in the same string.
// * time
// 	the time searched in ms, this should be sent together with the pv.
// * nodes
// 	x nodes searched, the engine should send this info regularly
// * pv  ...
// 	the best line found
// * multipv
// 	this for the multi pv mode.
// 	for the best move/pv add "multipv 1" in the string when you send the pv.
// 	in k-best mode always send all k variants in k strings together.
// * score
// 	* cp
// 		the score from the engine's point of view in centipawns.
// 	* mate
// 		mate in y moves, not plies.
// 		If the engine is getting mated use negativ values for y.
// 	* lowerbound
// 	  the score is just a lower bound.
// 	* upperbound
// 	   the score is just an upper bound.
// * currmove
// 	currently searching this move
// * currmovenumber
// 	currently searching move number x, for the first move x should be 1 not 0.
// * hashfull
// 	the hash is x permill full, the engine should send this info regularly
// * nps
// 	x nodes per second searched, the engine should send this info regularly
// * tbhits
// 	x positions where found in the endgame table bases
// * cpuload
// 	the cpu usage of the engine is x permill.
// * string
// 	any string str which will be displayed be the engine,
// 	if there is a string command the rest of the line will be interpreted as .
// * refutation   ...
//    move  is refuted by the line  ... , i can be any number >= 1.
//    Example: after move d1h5 is searched, the engine can send
//    "info refutation d1h5 g6h5"
//    if g6h5 is the best answer after d1h5 or if g6h5 refutes the move d1h5.
//    if there is norefutation for d1h5 found, the engine should just send
//    "info refutation d1h5"
// 	The engine should only send this if the option "UCI_ShowRefutations" is set to true.
// * currline   ...
//    this is the current line the engine is calculating.  is the number of the cpu if
//    the engine is running on more than one cpu.  = 1,2,3....
//    if the engine is just using one cpu,  can be omitted.
//    If  is greater than 1, always send all k lines in k strings together.
// 	The engine should only send this if the option "UCI_ShowCurrLine" is set to true.
type Info struct {
	Depth             int
	Seldepth          int
	PV                []*chess.Move
	Multipv           int
	Time              time.Duration
	Nodes             int
	Score             Score
	CurrentMove       *chess.Move
	CurrentMoveNumber int
	Hashfull          int
	NPS               int
	TBHits            int
	CPULoad           int
}

// Score corresponds to the "info"'s score engine output:
// * score
// * cp
// 	the score from the engine's point of view in centipawns.
// * mate
// 	mate in y moves, not plies.
// 	If the engine is getting mated use negativ values for y.
// * lowerbound
//   the score is just a lower bound.
// * upperbound
//    the score is just an upper bound.
// * win
//    the probability value of a win from the engine's point of view
// * draw
//    the probability value of a draw from the engine's point of view
// * loss
//    the probability value of a loss from the engine's point of view

type Score struct {
	CP         int
	Mate       int
	LowerBound bool
	UpperBound bool
	Win        int
	Draw       int
	Loss       int
}

// WinPct returns the probability a given position can be converted into a
// win from the engine's perspective on a scale from 0.0->1.0. The engine
// must support the UCI_ShowWDL option and it must be set in order for this
// function to work.
func (score *Score) WinPct() (float32, error) {
	total := score.Win + score.Draw + score.Loss
	if total == 0 {
		return 0.0, missingWdlErr
	}
	return float32(score.Win) / float32(total), nil
}

// DrawPct returns the probability a given position can be converted into a
// draw from the engine's perspective on a scale from 0.0->1.0. The engine
// must support the UCI_ShowWDL option and it must be set in order for this
// function to work.
func (score *Score) DrawPct() (float32, error) {
	total := score.Win + score.Draw + score.Loss
	if total == 0 {
		return 0.0, missingWdlErr
	}
	return float32(score.Draw) / float32(total), nil
}

// LossPct returns the probability a given position can be converted into a
// draw from the engine's perspective on a scale from 0.0->1.0. The engine
// must support the UCI_ShowWDL option and it must be set in order for this
// function to work.
func (score *Score) LossPct() (float32, error) {
	total := score.Win + score.Draw + score.Loss
	if total == 0 {
		return 0.0, missingWdlErr
	}
	return float32(score.Loss) / float32(total), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface and parses
// data like the following:
// info depth 24 seldepth 32 multipv 1 score cp 29 wdl 791 209 0 nodes 5130101 nps 819897 hashfull 967 tbhits 0 time 6257 pv d2d4
func (info *Info) UnmarshalText(text []byte) error {
	parts := strings.Split(string(text), " ")
	if len(parts) == 0 {
		return errors.New("uci: invalid info line " + string(text))
	}
	if parts[0] != "info" {
		return errors.New("uci: invalid info line " + string(text))
	}
	ref := ""
	for i := 1; i < len(parts); i++ {
		s := parts[i]
		switch s {
		case "score":
			continue
		case "lowerbound":
			info.Score.LowerBound = true
			continue
		case "upperbound":
			info.Score.UpperBound = true
			continue
		}
		if ref == "" {
			ref = s
			continue
		}
		switch ref {
		case "depth":
			v, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			info.Depth = v
		case "seldepth":
			v, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			info.Seldepth = v
		case "multipv":
			v, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			info.Multipv = v
		case "cp":
			v, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			info.Score.CP = v
		case "wdl":
			var err error
			info.Score.Win, err = strconv.Atoi(parts[i])
			if err != nil {
				return err
			}
			info.Score.Draw, err = strconv.Atoi(parts[i+1])
			if err != nil {
				return err
			}
			info.Score.Loss, err = strconv.Atoi(parts[i+2])
			if err != nil {
				return err
			}
			i += 2
		case "nodes":
			v, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			info.Nodes = v
		case "mate":
			v, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			info.Score.Mate = v
		case "currmovenumber":
			v, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			info.CurrentMoveNumber = v
		case "currmove":
			m, err := chess.UCINotation{}.Decode(nil, s)
			if err != nil {
				return err
			}
			info.CurrentMove = m
		case "hashfull":
			v, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			info.Hashfull = v
		case "tbhits":
			v, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			info.TBHits = v
		case "time":
			v, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			info.Time = time.Millisecond * time.Duration(v)
		case "nps":
			v, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			info.NPS = v
		case "cpuload":
			v, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			info.CPULoad = v
		case "pv":
			m, err := chess.UCINotation{}.Decode(nil, s)
			if err != nil {
				return err
			}
			info.PV = append(info.PV, m)
		}
		if ref != "pv" {
			ref = ""
		}
	}
	return nil
}
