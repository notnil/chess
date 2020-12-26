package opening

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/notnil/chess"
)

// BookECO represents the Encyclopaedia of Chess Openings https://en.wikipedia.org/wiki/Encyclopaedia_of_Chess_Openings
// BookECO is safe for concurrent use.
type BookECO struct {
	root             *node
	startingPosition *chess.Position
}

// NewBookECO returns a new BookECO.  This operation has to parse 2k rows of CSV data and insert it into a graph
// so it can take some time.
func NewBookECO() *BookECO {
	startingPosition := &chess.Position{}
	if err := startingPosition.UnmarshalText([]byte("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")); err != nil {
		panic(err)
	}
	b := &BookECO{
		root: &node{
			children: map[string]*node{},
			pos:      chess.NewGame().Position(),
			label:    label(),
		},
		startingPosition: startingPosition,
	}
	r := csv.NewReader(bytes.NewBuffer(ecoData))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for i, row := range records {
		if i == 0 {
			continue
		}
		o := &Opening{code: row[0], title: row[1], pgn: row[2]}
		b.insert(o)
	}
	return b
}

// Find implements the Book interface
func (b *BookECO) Find(moves []*chess.Move) *Opening {
	for n := b.followPath(b.root, moves); n != nil; n = n.parent {
		if n.opening != nil {
			return n.opening
		}
	}
	return nil
}

// Possible implements the Book interface
func (b *BookECO) Possible(moves []*chess.Move) []*Opening {
	n := b.followPath(b.root, moves)
	openings := []*Opening{}
	for _, n := range b.nodeList(n) {
		if n.opening != nil {
			openings = append(openings, n.opening)
		}
	}
	return openings
}

func (b *BookECO) followPath(n *node, moves []*chess.Move) *node {
	if len(moves) == 0 {
		return n
	}
	c, ok := n.children[moves[0].String()]
	if !ok {
		return n
	}
	return b.followPath(c, moves[1:])
}

func (b *BookECO) insert(o *Opening) error {
	posList := []*chess.Position{b.startingPosition}
	moves := []*chess.Move{}
	for _, s := range parseMoveList(o.pgn) {
		pos := posList[len(posList)-1]
		m, err := chess.LongAlgebraicNotation{}.Decode(pos, s)
		if err != nil {
			panic(err)
		}
		moves = append(moves, m)
		posList = append(posList, pos.Update(m))
	}
	n := b.root
	b.ins(n, o, posList[1:], moves)
	return nil
}

func (b *BookECO) ins(n *node, o *Opening, posList []*chess.Position, moves []*chess.Move) {
	pos := posList[0]
	move := moves[0]
	moveStr := move.String()
	var child *node
	for mv, c := range n.children {
		if mv == moveStr {
			child = c
			break
		}
	}
	if child == nil {
		child = &node{
			parent:   n,
			children: map[string]*node{},
			pos:      pos,
			label:    label(),
		}
		n.children[moveStr] = child
	}
	if len(posList) == 1 {
		child.opening = o
		return
	}
	b.ins(child, o, posList[1:], moves[1:])
}

type node struct {
	parent   *node
	children map[string]*node
	opening  *Opening
	pos      *chess.Position
	label    string
}

func (b *BookECO) draw(w io.Writer) error {
	s := "digraph g {\n"
	for _, n := range b.nodeList(b.root) {
		title := ""
		if n.opening != nil {
			title = n.opening.title
		}
		s += fmt.Sprintf(`%s [label="%s"];`+"\n", n.label, title)
		for m, c := range n.children {
			s += fmt.Sprintf(`%s -> %s [label="%s"];`+"\n", n.label, c.label, m)
		}
	}
	s += "}"
	_, err := w.Write([]byte(s))
	return err
}

func (b *BookECO) nodes(root *node, ch chan *node) {
	ch <- root
	for _, c := range root.children {
		b.nodes(c, ch)
	}
}

func (b *BookECO) nodeList(root *node) []*node {
	ch := make(chan *node)
	go func() {
		b.nodes(root, ch)
		close(ch)
	}()
	nodes := []*node{}
	for n := range ch {
		nodes = append(nodes, n)
	}
	return nodes
}

var (
	labelCount = 0
	alphabet   = "abcdefghijklmnopqrstuvwxyz"
)

func label() string {
	s := "a" + fmt.Sprint(labelCount)
	labelCount++
	return s
}

// 1.b2b4 e7e5 2.c1b2 f7f6 3.e2e4 f8b4 4.f1c4 b8c6 5.f2f4 d8e7 6.f4f5 g7g6
func parseMoveList(pgn string) []string {
	strs := strings.Split(pgn, " ")
	cp := []string{}
	for _, s := range strs {
		i := strings.Index(s, ".")
		if i == -1 {
			cp = append(cp, s)
		} else {
			cp = append(cp, s[i+1:])
		}
	}
	return cp
}
