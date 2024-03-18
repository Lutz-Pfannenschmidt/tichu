package tichu

import "sort"

type Combination struct {
	// the slice of cards that make up the combination
	Cards []*Card
	// the lowest value of the combination
	Base int8
	// the amount of different values in the combination
	length int8
	// the maximum amount of cards with the same value
	width int8
}

// NewCombination creates a new Combination from a slice of cards.
func NewCombination(cards []*Card) *Combination {
	containsPhoenix := false

	// sort the cards by value
	sort.SliceStable(cards, func(i, j int) bool {
		return cards[i].Type.Value() < cards[j].Type.Value()
	})

	// new combination
	res := &Combination{Cards: cards}

	// if no cards, return empty (the player passes)
	if len(cards) == 0 {
		return res
	}

	// Special case DOG (if Dog is only Card)
	if len(cards) == 1 {
		if cards[0].Type == DOG {
			res.length = -1
			res.width = -1
			return res
		}
	}

	// count the amount of each value
	amounts := map[int]int{}

	for i := range cards {
		// only add the value if it is not a special card
		if !cards[i].Type.IsSpecial() {
			amounts[cards[i].Type.Value()]++
		} else {
			// if it is special card, set the flag

			// if it also is a phoenix, set the flag
			if cards[i].Type == PHOENIX {
				containsPhoenix = true
			}
		}
	}

	// Define start and end
	start := 25
	end := 0

	for k := range amounts {
		if k < start {
			start = k
		}
		if k > end {
			end = k
		}
	}

	// If len=5 -> if a triple and a pair => full house
	if len(amounts) == 2 && res.length == 5 {
		//check whether the amounts match the full house pattern
		if amounts[start] == 3 && amounts[end] == 2 {
			return res
		}
	}

	// Check if the Phoenix card is present in the combination
	if containsPhoenix {
		// If Phoenix is the only card, set the base value of the combination to -1 and return the combination
		if len(cards) == 1 {
			res.Base = -1
			res.length = 1
			res.width = 1
			return res
		}

		// Initialize variables to track the minimum and maximum card values, and the key associated with the minimum value
		min := 5
		minKey := -1
		max := 0
		// Iterate over the range of card values
		for i := start; i <= end; i++ {
			// Update the minimum value and associated key if a lower value is found
			if amounts[i] < min {
				min = amounts[i]
				minKey = i
			}
			// Update the maximum value if a higher value is found
			if amounts[i] > max {
				max = amounts[i]
			}
		}

		if min == max && start != end {
			amounts[end+1]++
		} else {
			amounts[minKey]++
		}
	}

	// Set the length of the combination to the number of unique card values
	res.length = int8(len(amounts))

	// Set the width of the combination to the number of cards with the same value as the start value
	res.width = int8(amounts[start])
	// Iterate over the values in the amounts map
	for _, v := range amounts {
		// If any value is not equal to the width, return nil as the combination is not valid
		if v != int(res.width) {
			return nil
		}
	}

	if res.length < 5 && res.width < 2 && res.length != 1 {
		return nil
	}

	return res
}
