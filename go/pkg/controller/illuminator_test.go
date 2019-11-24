package controller

import (
	"bytes"
	"testing"

	tree "github.com/twstokes/treedivided"

	"github.com/twstokes/treedivided/pkg/mock"
)

func TestShowFanfareForTeam(t *testing.T) {
	m := &mock.MockReadWriteCloser{}

	c := NewController(m)
	c.ShowFanfareForTeam(tree.TeamA)

	expected := []byte{byte(showFanfareForTeam), byte(tree.TeamA), 0, 0, 0, 0, 0, 0}

	if bytes.Compare(m.B.Bytes(), expected) != 0 {
		t.Error("payload didn't match expectation")
	}
}

func TestSetGameScore(t *testing.T) {
	m := &mock.MockReadWriteCloser{}

	c := NewController(m)
	g := tree.GameState{TeamAScore: 10, TeamBScore: 20, GameOver: false}
	c.SetGameScore(&g)

	expected := []byte{byte(updateScoreStates), 10, 20, 0, 0, 0, 0, 0}

	if bytes.Compare(m.B.Bytes(), expected) != 0 {
		t.Error("payload didn't match expectation")
	}
}

func TestSetWinner(t *testing.T) {
	m := &mock.MockReadWriteCloser{}

	c := NewController(m)
	c.SetWinner(tree.TeamA)

	expected := []byte{byte(setWinner), byte(tree.TeamA), 0, 0, 0, 0, 0, 0}

	if bytes.Compare(m.B.Bytes(), expected) != 0 {
		t.Error("payload didn't match expectation")
	}
}

func TestSetTeamColors(t *testing.T) {
	m := &mock.MockReadWriteCloser{}

	c := NewController(m)

	clemson := tree.Team{
		ID:   tree.TeamA,
		Name: "Clemson Tigers",
		Colors: tree.TeamColors{
			Primary:   tree.NewColor(246, 103, 51),
			Secondary: tree.NewColor(82, 45, 128),
		},
		SongPath: "test.m4a",
	}

	c.SetTeamColors(clemson)

	expected := []byte{byte(setColorsForTeam), byte(tree.TeamA), 246, 103, 51, 82, 45, 128}

	if bytes.Compare(m.B.Bytes(), expected) != 0 {
		t.Error("payload didn't match expectation")
	}
}
