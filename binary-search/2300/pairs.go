package pairs

import (
	"slices"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	slices.Sort(potions)
	numPotions := len(potions)

	for i, spell := range spells {

		left := 0
		right := numPotions

		for left < right {
			mid := left + (right-left)/2
			currValue := int64(potions[mid] * spell)
			if currValue >= success {
				right = mid
			} else {
				left = mid + 1
			}
		}
		spells[i] = numPotions - left
	}
	return spells
}
