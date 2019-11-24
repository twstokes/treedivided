package tree

import (
	"errors"
	"reflect"
)

// GameState contains the state of the game
type GameState struct {
	TeamAScore int
	TeamBScore int
	GameOver   bool
}

func gameStateChanged(old *GameState, new *GameState) bool {
	return !reflect.DeepEqual(old, new)
}

func (g *GameState) getWinner() (TeamID, error) {
	if !g.GameOver {
		return 0, errors.New("game has to be over to get winner")
	}

	if g.TeamAScore == g.TeamBScore {
		return 0, errors.New("teams cannot be tied to get winner")
	}

	if g.TeamAScore > g.TeamBScore {
		return TeamA, nil
	}

	return TeamB, nil
}
