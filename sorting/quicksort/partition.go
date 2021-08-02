package quicksort

import (
	"math/rand"
	"time"
)

func Partion(a []int, p int, r int) int {

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(r-p) + p
	a[n], a[r] = a[r], a[n]

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
