package player

import (
	"os/exec"
)

// Play plays through omxplayer utility
func (p *LocalPlayer) Play(path string) error {
	c := exec.Command("omxplayer", path)
	return c.Start()
}
