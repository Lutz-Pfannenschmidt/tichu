package tichu

import "image/color"

type CardColor int

const (
	Red CardColor = iota
	Blue
	Black
	Green
	Special
)

func (c CardColor) RGBA() color.Color {
	switch c {
	case Red:
		return color.RGBA{R: 255, G: 0, B: 0, A: 255}
	case Blue:
		return color.RGBA{R: 0, G: 0, B: 255, A: 255}
	case Black:
		return color.RGBA{R: 0, G: 0, B: 0, A: 255}
	case Green:
		return color.RGBA{R: 0, G: 255, B: 0, A: 255}
	case Special:
		return color.RGBA{R: 255, G: 255, B: 0, A: 255}
	}
	return color.RGBA{}
}
