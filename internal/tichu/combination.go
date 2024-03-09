package tichu

import "sort"

type Combination struct {
	Cards           []*Card
	Base            int8
	length          int8
	level           int8
	containsPhoenix bool
}

// NewCombination creates a new Combination from a slice of cards.
func NewCombination(cards []*Card) *Combination {
	isLegal := true
	containsSpecial := false

	// sort the cards by value
	sort.SliceStable(cards, func(i, j int) bool {
		return cards[i].Type.Value() < cards[j].Type.Value()
	})

	// new combination
	res := &Combination{Cards: cards}

	// create the combination
	res.length = int8(len(cards))
	res.Base = int8(cards[0].Type.Value())

	amounts := map[int]int{}

	for i := int8(0); i < res.length; i++ {
		if cards[i].Type != PHOENIX {
			amounts[cards[i].Type.Value()]++
		}
		if cards[i].Type.IsSpecial() {
			containsSpecial = true
		}
	}

	if containsSpecial && !res.containsPhoenix {
		if len(cards) != 1 {
			isLegal = false
		}
	}

	if isLegal {
		return res
	}
	return nil
}
