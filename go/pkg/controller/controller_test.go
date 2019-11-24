package controller

import (
	"bytes"
	"testing"

	"github.com/twstokes/treedivided/pkg/mock"
)

func TestSendCommand(t *testing.T) {
	m := &mock.MockReadWriteCloser{}

	c := NewController(m)
	d := []byte{100}

	err := c.sendCommand(showFanfareForTeam, d)
	if err != nil {
		t.Error(err)
	}

	expected := []byte{byte(showFanfareForTeam), 100, 0, 0, 0, 0, 0, 0}

	if bytes.Compare(m.B.Bytes(), expected) != 0 {
		t.Error("failed to write expected payload")
	}
}

func TestSendCommandTooLargePayload(t *testing.T) {
	m := &mock.MockReadWriteCloser{}

	c := NewController(m)
	d := []byte{0, 0, 0, 0, 0, 0, 0, 0}

	err := c.sendCommand(showFanfareForTeam, d)
	if err == nil {
		t.Error("expected payload size error")
	}
}
