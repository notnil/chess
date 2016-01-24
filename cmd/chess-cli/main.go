package main

import (
	"fmt"

	"github.com/loganjspears/chess"
	"github.com/loganjspears/chess/ai"
)

type inputPlayer struct{}

func (p inputPlayer) Move(gs *chess.Position) string {
	fmt.Println("Enter Move in Algebraic Notation (ex. 'e4'):")
	input := ""
	fmt.Scanln(&input)
	return input
}

func main() {
	g := chess.NewGame()
	p := inputPlayer{}
	cpu := ai.New(3)
	count := 0
	for g.Outcome() == chess.NoOutcome {
		fmt.Println(g.Position().Board().Draw())
		fmt.Println(g.Position().String())
		if count%2 == 0 {
			alg := p.Move(g.Position())
			if err := g.MoveAlg(alg); err != nil {
				fmt.Println("invalid move")
				continue
			}
		} else {
			move := cpu.Move(g.Position())
			g.Move(move)
			fmt.Println("CPU moved ", move.String())
		}
		count++
	}
	fmt.Println(g.Position().Board().Draw())
	fmt.Printf("Game completed. %s by %s.\n", g.Outcome(), g.Method())
	fmt.Println(g.String())
}
