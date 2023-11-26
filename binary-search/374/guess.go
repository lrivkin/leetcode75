package guess

type guesser struct {
	targetNum int
}

func (g guesser) guess(num int) int {
	if num == g.targetNum {
		return 0
	} else if num > g.targetNum {
		return -1
	}
	return 1
}

func guessNumber(g guesser, n int) int {
	minN := 0

	myGuess := n / 2
	for {
		g := g.guess(myGuess)
		//fmt.Printf("g: %v, guess: %d, minN: %d, maxN: %d\n", g, myGuess, minN, n)
		if g > 0 {
			// 1: Your guess is lower than the number I picked (i.e. num < pick).
			minN = myGuess + 1
		} else if g < 0 {
			// -1: Your guess is higher than the number I picked (i.e. num > pick).
			n = myGuess - 1
		} else {
			// 0: your guess is equal to the number I picked (i.e. num == pick).
			return myGuess
		}
		myGuess = minN + (n-minN)/2
	}

	return -1
}
