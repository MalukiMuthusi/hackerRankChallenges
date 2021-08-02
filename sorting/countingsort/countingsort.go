package countingsort

func CountingSort(A []int, k int) []int {
	l := len(A)           // size of the array
	B := make([]int, l)   // holds the result
	C := make([]int, k+1) // temporary array

	// make sure C is initialized.
	//THIS IS UNNECESSARY IN Go

	// count the number of element's x
	for j := 0; j < l; j++ {
		x := A[j]
		C[x] += 1 // count the number of element's x
	}

	for i := 1; i <= k; i++ {
		C[i] = C[i] + C[i-1]
		// C[i] now contains the number of elements less than or equal to i
	}

	for j := l - 1; j >= 0; j-- {
		x := A[j] // value of the element in the original array
		n := C[x] // number of elements less than or equal to x

		B[n-1] = x // place x in its correct index

		C[x] = C[x] - 1 // decrement the counter of x in C
	}

	return B
}
