package countingsort_test

import (
	"testing"

	"github.com/MalukiMuthusi/hackerRankChallenges/sorting/countingsort"
	"github.com/stretchr/testify/assert"
)

func TestCountingSort(t *testing.T) {
	tests := []struct {
		name   string
		A      []int
		k      int
		output []int
	}{
		{
			name:   "happy case",
			A:      []int{2, 5, 0, 3, 2, 3, 0, 5},
			k:      5,
			output: []int{0, 0, 2, 2, 3, 3, 5, 5},
		},
		{
			name:   "random numbers 1-9",
			A:      []int{3, 6, 1, 2, 9, 4, 8, 7, 0, 5},
			k:      9,
			output: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:   "student's marks in science subject",
			A:      []int{76, 31, 92, 56, 59, 81, 46, 27, 99, 66},
			k:      99,
			output: []int{27, 31, 46, 56, 59, 66, 76, 81, 92, 99},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := countingsort.CountingSort(tt.A, tt.k)
			assert.Equal(t, tt.output, ans)
		})
	}

}
