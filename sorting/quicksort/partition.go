package quicksort

import (
	"math/rand"
	"time"
)

// Partition sorts the array and partitions it into two parts and returns the index
func Partion(a []int, p int, r int, pv int) int {
	a[pv], a[r] = a[r], a[pv]

	x := a[r]
	i := p - 1

	for j := p; j < r; j++ {
		if a[j] <= x {
			i += 1
			a[i], a[j] = a[j], a[i]
		}
	}

	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func RandomizePivot(A []int, p int, r int) int {
	if (r - p) <= 0 {
		return p
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(r-p) + p
	return n
}
