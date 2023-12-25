package flowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	openPlots := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}
		// left plot is empty
		if (i == 0 || flowerbed[i-1] == 0) &&
			// right plot is empty
			(i+1 == len(flowerbed) || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			openPlots++

			if openPlots == n {
				return true
			}
		}

	}
	return false
}
