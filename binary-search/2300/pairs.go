package pairs

import (
	"slices"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	//fmt.Printf("%v\n", potions)

	// precalculate the number that will be successful + store in a map
	potionsCountMap := map[int]int{}
	for i, potion := range potions {
		if _, seen := potionsCountMap[potion]; seen {
			continue
		}
		potionsCountMap[potion] = len(potions) - i
	}

	//fmt.Printf("%v\n", potionsCountMap)

	for i, spell := range spells {
		searchPotion := int(success) / spell
		if int(success)%spell != 0 {
			searchPotion += 1
		}

		//fmt.Printf("success = %d, spell=%d, need minimum potion = %d\n", success, spell, searchPotion)
		idx, exists := slices.BinarySearch(potions, searchPotion)

		if exists {
			//fmt.Printf("found minimum potion = %d idx=%d, matches %d spells\n", potions[idx], idx, potionsCountMap[potions[idx]])
			spells[i] = potionsCountMap[potions[idx]]
			continue
		}

		if idx >= len(potions) {
			//fmt.Printf("no matches\n")
			spells[i] = 0
			continue
		}

		//fmt.Printf("closest match = %d idx=%d, matches %d spells\n", potions[idx], idx, potionsCountMap[potions[idx]])
		spells[i] = potionsCountMap[potions[idx]]
	}
	return spells
}
