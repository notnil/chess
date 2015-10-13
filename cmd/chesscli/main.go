package main

import (
	"fmt"

	"github.com/loganjspears/chess"
)

func main() {
	g := chess.NewGame()
	for g.Outcome() == chess.NoOutcome {
		fmt.Println(g.State().Board().Draw())
		fmt.Println(g.State().String())
		input := ""
		fmt.Scanln(&input)
		if input == "back" {
			g = g.TakeBack(1)
			continue
		} else if input == "pgn" {
			fmt.Println(g.String())
			continue
		}
		if err := g.MoveAlg(input); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Printf("Game completed. %s by %s.\n", g.Outcome(), g.Method())
	fmt.Println(g.String())
}
