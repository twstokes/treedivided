package player

import (
	"os/exec"
)

// Play plays through afplay utility
func (p *LocalPlayer) Play(path string) error {
	c := exec.Command("afplay", path)
	return c.Start()
}
