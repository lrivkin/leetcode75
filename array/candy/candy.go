package main

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {

	result := make([]bool, len(candies))

	maxCandies := slices.Max(candies)

	for i, numCandies := range candies {
		if numCandies+extraCandies >= maxCandies {
			result[i] = true
		}
	}

	return result
}
