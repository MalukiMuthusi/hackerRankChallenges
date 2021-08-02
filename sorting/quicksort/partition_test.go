package quicksort_test

import (
	"testing"

	"github.com/MalukiMuthusi/hackerRankChallenges/sorting/quicksort"
	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		name   string
		A      []int
		p      int
		r      int
		output int
	}{
		{
			name:   "happy case",
			A:      []int{2, 8, 7, 1, 3, 5, 6, 4},
			p:      0,
			r:      7,
			output: 3,
		},

		{
			name:   "sorted items",
			A:      []int{1, 2, 3, 8, 5, 6, 7, 4},
			p:      0,
			r:      7,
			output: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := quicksort.Partion(tt.A, tt.p, tt.r)
			assert.Equal(t, tt.output, n)
		})
	}
}
