package heapsort

import "fmt"

// Left returns the index of the left node
func Left(i int) (int, error) {
	if i < 0 {
		return 0, fmt.Errorf("%d should be a positive number", i)
	}
	return i << 1, nil // i << 1 == i*2
}

// Right returns the index of the right node
func Right(i int) (int, error) {
	if i < 0 {
		return 0, fmt.Errorf("%d should be a positive number", i)
	}
	return i<<1 + 1, nil // i << 1 == i*2
}

// Parent returns the index of the parent node
func Parent(i int) (int, error) {
	if i < 0 {
		return 0, fmt.Errorf("%d should be a positive number", i)
	}
	return i >> 1, nil // i >> 1 == i / 2
}

// MaxHeapify maintains the max-heap property
//	It assumes that the binary trees rooted at Left(i) and Right(i) are
//	maxHeaps but a[i] might be smaller than its children thus violating maxHeap
//	property
//	it lets the value at a[i] float down in the max heap so that the subtree
//	rooted at index i obeys the maxHeap property
func MaxHeapify(a []int, i int) {
	l, err := Left(i)
	if err != nil {
		return
	}

	r, _ := Right(i)
	if err != nil {
		return
	}

	ln := len(a)
	var largest int = i

	if l < ln && a[l] > a[i] {
		largest = l
	}

	if r < ln && a[r] > a[largest] {
		largest = r
	}

	if largest != i {
		a[largest], a[i] = a[i], a[largest]
		MaxHeapify(a, largest)
	}

}

// BuildMaxHeap, builds a max heap from an array
func BuildMaxHeap(a []int) []int {
	for i := len(a) / 2; i > 0; i-- {
		MaxHeapify(a, i)
	}
	return a
}

func HeapSort(a []int) []int {
	BuildMaxHeap(a)
	sortedArray := make([]int, 1)
	for i := len(a) - 1; i > 0; i-- {
		a[1], a[i] = a[i], a[1]
		sortedArray = append(sortedArray, a[i])
		a = a[:len(a)-1]
		MaxHeapify(a, 1)
	}
	return sortedArray
}
