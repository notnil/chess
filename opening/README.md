# opening

## Datasource

The [Encyclopaedia of Chess Openings](https://en.wikipedia.org/wiki/Encyclopaedia_of_Chess_Openings) (ECO) functions as the datasource for this package.  A consise list of openings with PGNs can be found [here](http://www.webcitation.org/query?url=http://www.geocities.com/siliconvalley/lab/7378/eco.htm&date=2010-02-20+10:14:24).

## Visual

Advance Variation subtree of the French Defense:

![subtree](https://github.com/notnil/opening/raw/master/test.png)

## Example

```go   
package main

import (
    "fmt"

    "github.com/notnil/chess"
    "github.com/notnil/chess/opening"
)

func main(){
    g := chess.NewGame()
	g.MoveStr("e4")
	g.MoveStr("e6")

	// print French Defense
	book := opening.NewBookECO()
	o := book.Find(g.Moves())
	fmt.Println(o.Title())
}
```