package tichu

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

type CardType int

func (c CardType) String() string {
	return cardStrings[c]
}

func (c CardType) Value() int {
	return cardValues[c]
}

func (c CardType) IsSpecial() bool {
	return c == DOG || c == MAHJONG || c == PHEONIX || c == DRAGON
}

func (c CardType) Points() int {
	switch c {
	case FIVE:
		return 5
	case TEN:
	case KING:
		return 10
	case DRAGON:
		return 25
	case PHEONIX:
		return -25
	}
	return 0
}

type Card struct {
	Type  CardType
	Color CardColor
}

func (c *Card) Render(gc *draw2dimg.GraphicContext, x, y float64) {
	gc.SetFillColor(c.Color.RGBA())
	gc.SetStrokeColor(c.Color.RGBA())
	gc.BeginPath()
	draw2dkit.RoundedRectangle(gc, x, y, x+30, y+30, 5, 5)
	img, err := draw2dimg.LoadFromPngFile("")
	if err != nil {
		panic(err)
	}
	draw2dimg.DrawImage()
	gc.FillStroke()
	gc.Close()
}

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
	PHEONIX: "PHEONIX",
	DRAGON:  "DRAGON",
}

var cardValues = map[CardType]int{
	DOG:     -1,
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
	PHEONIX: -1,
	DRAGON:  25,
}

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
	PHEONIX
	DRAGON
)
