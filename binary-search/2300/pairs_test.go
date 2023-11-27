package pairs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccessfulPairs(t *testing.T) {
	type testCase struct {
		spells  []int
		potions []int
		success int64
		want    []int
	}

	for i, tt := range []testCase{
		{
			spells:  []int{5, 1, 3},
			potions: []int{1, 2, 3, 4, 5},
			success: 7,
			want:    []int{4, 0, 3},
		},
		{
			spells:  []int{3, 1, 2},
			potions: []int{8, 5, 8},
			success: 16,
			want:    []int{2, 0, 2},
		},
	} {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, tt.want, successfulPairs(tt.spells, tt.potions, tt.success))
		})
	}
}
