package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKidsWithCandies(t *testing.T) {
	for i, example := range []struct {
		candies      []int
		extraCandies int
		result       []bool
	}{
		{
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			result:       []bool{true, true, true, false, true},
		},
		{
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			result:       []bool{true, false, false, false, false},
		},
		{
			candies:      []int{12, 1, 12},
			extraCandies: 10,
			result:       []bool{true, false, true},
		},
	} {
		t.Run(fmt.Sprintf("example %d", i), func(t *testing.T) {
			t.Logf("kidsWithCandies(%v, %d)", example.candies, example.extraCandies)

			assert.Equal(t, example.result, kidsWithCandies(example.candies, example.extraCandies))
		})
	}
}
