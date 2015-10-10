package chess

import (
	"log"
	"regexp"
	"strings"
)

/*
[Event "Troll Masters"]
[Site "Gausdal NOR"]
[Date "2001.01.05"]
[Round "1"]
[White "Edvardsen,R"]
[Black "Carlsen,Magnus"]
[Result "1/2-1/2"]
[WhiteElo "2055"]
[BlackElo ""]
[ECO "D12"]

1.d4 Nf6 2.Nf3 d5 3.e3 Bf5 4.c4 c6 5.Nc3 e6 6.Bd3 Bxd3 7.Qxd3 Nbd7 8.b3 Bd6
9.O-O O-O 10.Bb2 Qe7 11.Rad1 Rad8 12.Rfe1 dxc4 13.bxc4 e5 14.dxe5 Nxe5 15.Nxe5 Bxe5
16.Qe2 Rxd1 17.Rxd1 Rd8 18.Rxd8+ Qxd8 19.Qd1 Qxd1+ 20.Nxd1 Bxb2 21.Nxb2 b5
22.f3 Kf8 23.Kf2 Ke7  1/2-1/2
*/

func decodePGN(pgn string) (*Game, error) {
	tagPairs := getTagPairs(pgn)
	moveStrs, outcome := moveList(pgn)
	g := NewGame(TagPairs(tagPairs))
	for _, alg := range moveStrs {
		if err := g.MoveAlg(alg); err != nil {
			log.Println(moveStrs)
			return nil, err
		}
	}
	g.outcome = outcome
	return g, nil
}

func encodePGN(g *Game) string {
	return ""
}

var (
	tagPairRegex = regexp.MustCompile(`[(.*)\w"(.*)"]`)
)

func getTagPairs(pgn string) map[string]string {
	tagPairs := map[string]string{}
	matches := tagPairRegex.FindAllString(pgn, -1)
	for _, m := range matches {
		results := tagPairRegex.FindStringSubmatch(m)
		if len(results) == 3 {
			tagPairs[results[1]] = results[2]
		}
	}
	return tagPairs
}

var (
	moveNumRegex = regexp.MustCompile(`(?:\d+\.+)?(.*)`)
)

func moveList(pgn string) ([]string, Outcome) {
	// remove comments
	text := removeSection("{", "}", pgn)
	// remove variations
	text = removeSection(`\(`, `\)`, text)
	// remove tag pairs
	text = removeSection(`\[`, `\]`, text)
	// remove line breaks
	text = strings.Replace(text, "\n", " ", -1)

	list := strings.Split(text, " ")
	filtered := []string{}
	var outcome Outcome
	for _, move := range list {
		move = strings.TrimSpace(move)
		switch move {
		case string(NoOutcome), string(WhiteWon), string(BlackWon), string(Draw):
			outcome = Outcome(move)
		case "":
		default:
			results := moveNumRegex.FindStringSubmatch(move)
			if len(results) == 2 && results[1] != "" {
				filtered = append(filtered, results[1])
			}
		}
	}
	return filtered, outcome
}

func removeSection(leftChar, rightChar, s string) string {
	r := regexp.MustCompile(leftChar + ".*?" + rightChar)
	for {
		i := r.FindStringIndex(s)
		if i == nil {
			return s
		}
		s = s[0:i[0]] + s[i[1]:len(s)]
	}
	return s
}
