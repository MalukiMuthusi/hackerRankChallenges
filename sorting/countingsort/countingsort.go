package countingsort

// sorting algorithm
// k is the largest digit of the values to be sorted
func CountingSort(A []int, k int) []int {
	l := len(A)           // size of the array
	B := make([]int, l)   // holds the result
	C := make([]int, k+1) // temporary array

	// count the how many times the digits appear in the array A
	for j := 0; j < l; j++ {
		x := A[j]
		C[x] += 1
	}

	// array C[i] has the index of the elements in Array A. This is the new index for elements to be sorted
	for i := 1; i <= k; i++ {
		C[i] = C[i] + C[i-1]
	}

	for j := l - 1; j >= 0; j-- {
		x := A[j] // value of the element in the original array
		n := C[x] // number of elements less than or equal to x

		B[n-1] = x // place x in its correct index

		C[x] = C[x] - 1 // decrement the counter of x in C
	}

	return B
}
