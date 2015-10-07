package chess

import "testing"

func TestGame(t *testing.T) {
	g := NewGame()
	g.Move(E2, E4, nil)
}
