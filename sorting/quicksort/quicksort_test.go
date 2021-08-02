package quicksort_test

import (
	"testing"

	"github.com/MalukiMuthusi/hackerRankChallenges/sorting/quicksort"
	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name   string
		A      []int
		p      int
		r      int
		output []int
	}{
		{
			name:   "happy case",
			A:      []int{2, 8, 7, 1, 3, 5, 6, 4},
			p:      0,
			r:      7,
			output: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},

		{
			name:   "random list of 8 values",
			A:      []int{5, 8, 4, 1, 3, 2, 6, 7},
			p:      0,
			r:      7,
			output: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quicksort.Quicksort(tt.A, tt.p, tt.r)
			assert.Equal(t, tt.output, tt.A)
		})
	}
}
