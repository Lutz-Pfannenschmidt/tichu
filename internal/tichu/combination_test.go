package tichu

import (
	"fmt"
	"testing"
)

func TestPheonix(t *testing.T) {
	fmt.Println(NewCombination([]*Card{&Card{Type: PHOENIX, Color: Special}}))
}

func TestDog(t *testing.T) {
	fmt.Println(NewCombination([]*Card{&Card{Type: DOG, Color: Special}}))
}
