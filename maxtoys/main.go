package maxtoys

func MaximumToys(prices []int, k int) int {
	prices = SortPrices(prices)
	s := 0
	c := 0
	for _, v := range prices {
		if s > k {
			break
		}

		if s+v > k {
			break
		}
		s += v
		c += 1
	}
	return c
}

func SortPrices(prices []int) []int {
	// use insertionSort to sort the prices
	l := len(prices)
	for i := 0; i < l; i++ {
		k := prices[i]
		j := i - 1
		for j > -1 && prices[j] > k {
			prices[j], prices[j+1] = prices[j+1], prices[j]
			j -= 1
		}
		prices[j+1] = k
	}
	return prices
}
