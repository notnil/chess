package main

import (
	"fmt"

	"github.com/loganjspears/chess"
	"github.com/loganjspears/chess/ai"
)

type inputPlayer struct{}

func (p inputPlayer) Move(gs *chess.GameState) string {
	fmt.Println("Enter Move in Algebraic Notation (ex. 'e4'):")
	input := ""
	fmt.Scanln(&input)
	return input
}

func main() {
	g := chess.NewGame()
	p := inputPlayer{}
	cpu := ai.New(2)
	count := 0
	for g.Outcome() == chess.NoOutcome {
		fmt.Println(g.State().Board().Draw())
		fmt.Println(g.State().String())
		if count%2 == 0 {
			alg := p.Move(g.State())
			if err := g.MoveAlg(alg); err != nil {
				fmt.Println("invalid move")
				continue
			}
		} else {
			move := cpu.Move(g.State())
			g.Move(move)
			fmt.Println("CPU moved ", move.String())
		}
		count++
	}
	fmt.Println(g.State().Board().Draw())
	fmt.Printf("Game completed. %s by %s.\n", g.Outcome(), g.Method())
	fmt.Println(g.String())
}
