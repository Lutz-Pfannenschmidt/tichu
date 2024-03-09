package websocket_test

import (
	"fmt"
	"github.com/Lutz-Pfannenschmidt/tichu/internal/tichu"
	"github.com/Lutz-Pfannenschmidt/tichu/internal/websocket"
	"testing"
)

func TestCombCompressing(t *testing.T) {
	cards := []*tichu.Card{
		&tichu.Card{Type: tichu.EIGHT, Color: tichu.Red},
		&tichu.Card{Type: tichu.ACE, Color: tichu.Red},
		&tichu.Card{Type: tichu.ACE, Color: tichu.Blue},
		&tichu.Card{Type: tichu.TWO, Color: tichu.Black},
		&tichu.Card{Type: tichu.THREE, Color: tichu.Red},
		&tichu.Card{Type: tichu.DRAGON, Color: tichu.Special},
		&tichu.Card{Type: tichu.DOG, Color: tichu.Special},
		&tichu.Card{Type: tichu.JACK, Color: tichu.Green},
		&tichu.Card{Type: tichu.EIGHT, Color: tichu.Green},
	}
	for _, card := range cards {
		fmt.Print(card)
		fmt.Print(" ")
	}
	fmt.Println()
	fmt.Println(websocket.CompressCombination(&cards))
	buffer := []*tichu.Card{
		&tichu.Card{Type: tichu.NINE, Color: tichu.Red},
		&tichu.Card{Type: tichu.PHEONIX, Color: tichu.Special},
		&tichu.Card{Type: tichu.ACE, Color: tichu.Green},
		&tichu.Card{Type: tichu.THREE, Color: tichu.Blue},
	}
	websocket.DecompressCombination(&buffer, websocket.CompressCombination(&cards))
	for _, card := range buffer {
		fmt.Print(card)
		fmt.Print(" ")
	}
	fmt.Println()
}

func TestIdSystem(t *testing.T) {
	cards := []*tichu.Card{
		&tichu.Card{Type: tichu.EIGHT, Color: tichu.Red},
		&tichu.Card{Type: tichu.ACE, Color: tichu.Red},
		&tichu.Card{Type: tichu.ACE, Color: tichu.Blue},
		&tichu.Card{Type: tichu.TWO, Color: tichu.Black},
		&tichu.Card{Type: tichu.THREE, Color: tichu.Red},
		&tichu.Card{Type: tichu.DRAGON, Color: tichu.Special},
		&tichu.Card{Type: tichu.DOG, Color: tichu.Special},
		&tichu.Card{Type: tichu.JACK, Color: tichu.Green},
		&tichu.Card{Type: tichu.EIGHT, Color: tichu.Green},
	}
	for _, card := range cards {
		fmt.Println(card)
		fmt.Println(websocket.CardToId(card))
		fmt.Println(websocket.IdToCard(websocket.CardToId(card)))
		fmt.Println()
	}
}
