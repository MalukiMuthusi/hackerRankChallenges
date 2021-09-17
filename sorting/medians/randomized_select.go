package medians

// return the ith smallest value of the array
func RandomizedSelect(A []int, p int, r int, i int) int {

	// when A has one element return it
	if p == r {
		return A[p]
	}

	// return a random index from the array
	q := RandomizedPartition(A, p, r)

	// k has the number of items p-q
	k := q - p + 1
	if i == k {
		return A[q]
	}

	if i < k {
		// i lies in the index between p and q
		return RandomizedSelect(A, p, q-1, i)
	}

	// i lies in the index between q and r
	return RandomizedSelect(A, q+1, r, i-k)
}
