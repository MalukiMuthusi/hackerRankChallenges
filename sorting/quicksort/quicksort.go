package quicksort

func Quicksort(A []int, p int, r int) {
	if p < r {
		pv := RandomizePivot(A, p, r)
		q := Partion(A, p, r, pv)
		Quicksort(A, p, q-1)
		Quicksort(A, q+1, r)
	}
}
