package tree

import (
	"testing"
)

func TestColorToByteSlice(t *testing.T) {
	c := NewColor(20, 30, 40)

	s := c.Bytes()
	if s[0] != c.R || s[1] != c.G || s[2] != c.B {
		t.Error("bytes didn't match expected color values")
	}
}
