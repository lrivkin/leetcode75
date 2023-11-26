package guess

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGuessNumber(t *testing.T) {

	type guessTest struct {
		n      int
		target int
	}

	for _, tt := range []guessTest{
		{10, 6},
		{1, 1},
		{2, 1},
		{1000, 50},
	} {
		t.Run(fmt.Sprintf("target = %d n = %d", tt.target, tt.n), func(t *testing.T) {
			assert.Equal(t, tt.target, guessNumber(guesser{tt.target}, tt.n))
		})
	}
}
