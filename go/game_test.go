package tree

import "testing"

func TestGetWinner(t *testing.T) {
	g := GameState{
		1, 0, true,
	}

	w, _ := g.getWinner()
	if w != TeamA {
		t.Error("expected team A to win")
	}
}

func TestGetWinnerTiedGame(t *testing.T) {
	g := GameState{
		0, 0, false,
	}

	w, err := g.getWinner()
	if err == nil {
		t.Error("expected error")
	}

	if w != 0 {
		t.Error("expected invalid winner id")
	}
}

func TestGetWinnerGameNotOver(t *testing.T) {
	g := GameState{
		1, 0, false,
	}

	_, err := g.getWinner()
	if err == nil {
		t.Error("expected error")
	}
}
