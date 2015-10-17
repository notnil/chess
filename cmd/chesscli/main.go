package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/loganjspears/chess"
)

func main() {
	app := cli.NewApp()
	app.Name = "chess"
	app.Usage = "command line tool for chess functions"
	app.Commands = []cli.Command{recordGame()}
	app.Run(os.Args)
}

func recordGame() cli.Command {
	return cli.Command{
		Name:  "record",
		Usage: "Enter moves in algebraic notation.  Enter 'back' to go back a move or 'pgn' to print the PGN.",
		Action: func(c *cli.Context) {
			g := chess.NewGame()
			for g.Outcome() == chess.NoOutcome {
				fmt.Println(g.State().Board().Draw())
				fmt.Println(g.State().String())
				fmt.Println("Enter Move in Algebraic Notation (ex. 'e4'):")
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
		},
	}
}
