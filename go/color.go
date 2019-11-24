package tree

import "image/color"

type Color struct {
	color.RGBA
}

func NewColor(r uint8, g uint8, b uint8) Color {
	c := color.RGBA{R: r, G: g, B: b}
	return Color{c}
}

// Bytes converts team colors to a slice of bytes
func (c *Color) Bytes() []byte {
	return []byte{
		c.R,
		c.G,
		c.B,
	}
}
