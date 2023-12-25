package flowers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		flowerbed []int
		n         int
		want      bool
	}{
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			want:      true,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			want:      false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			assert.Equalf(t, tt.want, canPlaceFlowers(tt.flowerbed, tt.n),
				"got wrong answer for canPlaceFlowers(%v, %d)", tt.flowerbed, tt.n,
			)
		})
	}
}
