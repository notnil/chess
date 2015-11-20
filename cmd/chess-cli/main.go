package main

import "github.com/loganjspears/chess"

type Player interface {
	Move(gs *chess.GameState) *chess.Move
}

type InputPlayer struct{}

func (p InputPlayer) Move(gs *chess.GameState) *chess.Move {

}

func main() {

}
