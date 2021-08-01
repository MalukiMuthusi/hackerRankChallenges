package quicksort

func Partion(a []int, p int, r int) int {
	x := a[r]
	i := p - 1

	for j := p; j < r-1; j++ {
		if a[j] <= x {
			i += 1
			a[i], a[j] = a[j], a[i]
		}
	}

	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}
