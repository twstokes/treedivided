package controller

import (
	"errors"
	"io"
)

// the maximum payload our controller expects for commands
const maxPayload = 8

// a controller command is defined by the controller software
type command uint8

const (
	showFanfareForTeam command = 1
	updateScoreStates  command = 2
	setWinner          command = 3
	setColorsForTeam   command = 4
)

// Controller sends commands with a ReadWriteCloser
type Controller struct {
	io.ReadWriteCloser
}

// NewController creates an Illuminator driven by a Writable
func NewController(r io.ReadWriteCloser) *Controller {
	return &Controller{r}
}

func (c *Controller) sendCommand(com command, d []byte) error {
	// combine our command and data
	data := append([]byte{byte(com)}, d...)

	if len(data) > maxPayload {
		return errors.New("payload too large")
	}

	// make a payload buffer to the expected size
	payload := make([]byte, maxPayload)
	copy(payload, data)

	_, err := c.Write(payload)
	if err != nil {
		return err
	}

	return nil
}
