package tree

import (
	"errors"
	"log"
	"time"
)

// Config stores the configuration for the tree
type Config struct {
	TeamA       Team
	TeamB       Team
	Illuminator Illuminator
	Fetcher     Fetcher
	Player      Player
}

// Fetcher fetches data about the game
type Fetcher interface {
	GetLatestState() (*GameState, error)
	WaitTime() time.Duration
}

// Player plays audio on the system
type Player interface {
	Play(path string) error
}

// Illuminator changes the lights on the tree
type Illuminator interface {
	SetTeamColors(teams Team) error
	ShowFanfareForTeam(id TeamID) error
	SetGameScore(*GameState) error
	SetWinner(id TeamID) error
}

// Tree is a 🎄 tree
type Tree struct {
	c     *Config
	state *GameState
}

// NewTree returns a new Tree
func NewTree(c *Config) *Tree {
	return &Tree{c: c}
}

// Run polls for new data using a fetcher
func (t *Tree) Run(s chan bool) {
	timer := time.NewTimer(0)

	for {
		select {
		case <-timer.C:
			timer.Reset(t.c.Fetcher.WaitTime())

			// get latest from fetcher
			log.Println("Fetching...")
			latestState, err := t.c.Fetcher.GetLatestState()
			if err != nil {
				log.Println("Error getting latest from fetcher!", err)
				break
			}

			prevState := t.state
			t.state = latestState

			if prevState == nil {
				log.Println("First run - updating the illuminator.")
				t.c.Illuminator.SetTeamColors(t.c.TeamA)
				t.c.Illuminator.SetTeamColors(t.c.TeamB)
				t.c.Illuminator.SetGameScore(t.state)
				break
			}

			if !gameStateChanged(prevState, latestState) {
				log.Println("Game state didn't change since last fetch.")
				break
			}

			if latestState.GameOver {
				log.Println("Game over!")
				t.gameOver()
				break
			}

			team, err := t.getTeamThatScored(prevState)
			if err != nil {
				log.Println("Score analysis error:", err)

				switch err.Error() {
				case "both changed", "reverse score":
					// we still want to update the MCU, but with no fanfare
					log.Println("Setting the game score.")
					t.c.Illuminator.SetGameScore(t.state)
				}

				break
			}

			log.Println(team.Name, "scored.")
			t.c.Player.Play(team.SongPath)
			t.c.Illuminator.ShowFanfareForTeam(team.ID)
			t.c.Illuminator.SetGameScore(t.state)
		case <-s:
			// stop fetching
			log.Println("Stopping fetcher.")
			return
		}
	}
}

func (t *Tree) gameOver() {
	winner, err := t.state.getWinner()
	if err != nil {
		log.Println(err)
		return
	}

	var team Team

	switch winner {
	case TeamA:
		team = t.c.TeamA
	case TeamB:
		team = t.c.TeamB
	}

	log.Println(team.Name, "won!")
	t.c.Illuminator.SetWinner(team.ID)
	t.c.Player.Play(team.SongPath)
}

func (t *Tree) getTeamThatScored(p *GameState) (*Team, error) {
	teamADelta := t.state.TeamAScore - p.TeamAScore
	teamBDelta := t.state.TeamBScore - p.TeamBScore

	if teamADelta > 0 && teamBDelta > 0 {
		// both scores changed
		return nil, errors.New("both changed")
	}

	if teamADelta == 0 && teamBDelta == 0 {
		// neither scores changed
		return nil, errors.New("neither changed")
	}

	if teamADelta < 0 || teamBDelta < 0 {
		// a score went into reverse, probably a correction
		return nil, errors.New("reverse score")
	}

	if teamADelta > 0 {
		return &t.c.TeamA, nil
	}

	return &t.c.TeamB, nil
}
