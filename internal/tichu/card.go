// Package tichu provides the necessary types and functions for a game of Tichu.
package tichu

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

// CardType represents the type of a card in the game of Tichu.
type CardType int

// String returns a string representation of the CardType.
func (c CardType) String() string {
	return cardStrings[c]
}

// Value returns the value of the CardType.
func (c CardType) Value() int {
	return cardValues[c]
}

// IsSpecial checks if the CardType is a special card in the game of Tichu.
func (c CardType) IsSpecial() bool {
	return c == DOG || c == MAHJONG || c == PHOENIX || c == DRAGON
}

// Points returns the points associated with the CardType.
func (c CardType) Points() int {
	switch c {
	case FIVE:
		return 5
	case TEN:
	case KING:
		return 10
	case DRAGON:
		return 25
	case PHOENIX:
		return -25
	}
	return 0
}

// Card represents a card in the game of Tichu.
type Card struct {
	Type  CardType
	Color CardColor
}

// Render draws the Card on a GraphicContext at the specified x and y coordinates.
func (c *Card) Render(gc *draw2dimg.GraphicContext, x, y float64) {
	gc.SetFillColor(c.Color.RGBA())
	gc.SetStrokeColor(c.Color.RGBA())
	gc.BeginPath()
	draw2dkit.RoundedRectangle(gc, x, y, x+30, y+30, 5, 5)
	gc.FillStroke()
	gc.Close()
}

// cardStrings maps CardType values to their string representations.
var cardStrings = map[CardType]string{
	DOG:     "DOG",
	MAHJONG: "MAHJONG",
	TWO:     "2",
	THREE:   "3",
	FOUR:    "4",
	FIVE:    "5",
	SIX:     "6",
	SEVEN:   "7",
	EIGHT:   "8",
	NINE:    "9",
	TEN:     "10",
	JACK:    "J",
	QUEEN:   "Q",
	KING:    "K",
	ACE:     "A",
	PHOENIX: "PHOENIX",
	DRAGON:  "DRAGON",
}

// cardValues maps CardType values to their integer values.
var cardValues = map[CardType]int{
	DOG:     -2,
	MAHJONG: 1,
	TWO:     2,
	THREE:   3,
	FOUR:    4,
	FIVE:    5,
	SIX:     6,
	SEVEN:   7,
	EIGHT:   8,
	NINE:    9,
	TEN:     10,
	JACK:    11,
	QUEEN:   12,
	KING:    13,
	ACE:     14,
	PHOENIX: -1,
	DRAGON:  25,
}

// Constants representing the different types of cards in the game of Tichu.
const (
	DOG CardType = iota
	MAHJONG
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE
	PHOENIX
	DRAGON
)
