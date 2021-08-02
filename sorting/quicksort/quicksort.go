package quicksort

func Quicksort(A []int, p int, r int) {
	if p < r {
		q := Partion(A, p, r)
		Quicksort(A, p, q-1)
		Quicksort(A, q+1, r)
	}
}
