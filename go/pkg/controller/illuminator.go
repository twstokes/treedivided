package controller

import (
	tree "github.com/twstokes/treedivided"
)

// ShowFanfareForTeam sends the command to show fanfare
func (c *Controller) ShowFanfareForTeam(teamID tree.TeamID) error {
	data := []byte{byte(teamID)}
	return c.sendCommand(showFanfareForTeam, data)
}

// SetGameScore sends the command to update the game score
func (c *Controller) SetGameScore(g *tree.GameState) error {
	data := []byte{byte(g.TeamAScore), byte(g.TeamBScore)}
	return c.sendCommand(updateScoreStates, data)
}

// SetWinner sends the command to set the winner
func (c *Controller) SetWinner(teamID tree.TeamID) error {
	data := []byte{byte(teamID)}
	return c.sendCommand(setWinner, data)
}

// SetTeamColors sends the command to set the team colors
func (c *Controller) SetTeamColors(t tree.Team) error {
	data := []byte{byte(t.ID)}
	data = append(data, t.Colors.Primary.Bytes()...)
	data = append(data, t.Colors.Secondary.Bytes()...)
	return c.sendCommand(setColorsForTeam, data)
}
