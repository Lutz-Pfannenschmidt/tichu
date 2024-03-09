package tichu

import "math/rand"

type Deck struct {
	Cards []*Card
}

func NewDeck() *Deck {
	deck := &Deck{}
	deck.generate()
	deck.Shuffle()
	return deck
}

var genericCards = []CardType{
	TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING, ACE,
}

var specialCards = []CardType{
	DOG, MAHJONG, DRAGON, PHEONIX,
}

// generate creates a new deck of cards
func (d *Deck) generate() {
	d.Cards = make([]*Card, 56)
	// Generate the generic cards
	for i, card := range genericCards {
		for j := 0; j < 4; j++ {
			d.Cards[i*4+j] = &Card{Type: card, Color: CardColor(j)}
		}
	}

	// Generate the special cards
	for i, card := range specialCards {
		d.Cards[52+i] = &Card{Type: card, Color: Special}
	}
}

// Shuffle the deck
func (d *Deck) Shuffle() {
	for i := range d.Cards {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

func (d *Deck) Draw() *Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}
