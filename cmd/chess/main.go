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
	app.Commands = []cli.Command{recordGame(), fen(), pgn()}
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

func fen() cli.Command {
	return cli.Command{
		Name:  "fen",
		Usage: "Enter FEN using standard input.",
		Action: func(c *cli.Context) {
			fen, err := chess.FEN(os.Stdin)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				return
			}
			g := chess.NewGame(fen)
			fmt.Fprintln(os.Stdout, g.State().String())
			fmt.Fprintln(os.Stdout, g.State().Board().Draw())
		},
	}
}

func pgn() cli.Command {
	return cli.Command{
		Name:  "pgn",
		Usage: "Enter PGN using standard input.  Use (k/j) to view the next or previous move, or quit to stop.",
		Action: func(c *cli.Context) {
			pgn, err := chess.PGN(os.Stdin)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				return
			}
			g := chess.NewGame(pgn)
			fmt.Println("Viewing PGN:")
			fmt.Println(g.String())
			states := g.States()
			count := 0
			for {
				state := states[count]
				fmt.Println(state.Board().Draw())
				fmt.Println(state.String())
				fmt.Println("Use (s/d) to view the previous or next move, or quit to stop.")
				input := ""
				fmt.Scanln(&input)
				switch input {
				case "d":
					if count == len(states)-1 {
						fmt.Printf("Game completed. %s by %s.\n", g.Outcome(), g.Method())
					} else {
						count++
					}
				case "s":
					if count != 0 {
						count--
					}
				case "quit":
					return
				}
			}
		},
	}
}
