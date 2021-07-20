package insertionsort_test

import (
	"testing"

	insertionsort "github.com/MalukiMuthusi/hackerRankChallenges/sorting/insertionSort"
	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T){
	tests:= []struct{
		name string
		input []int
		output []int
	}{
		{
			name: "happy case",
			input: []int{5,2,4,6,1,3},
			output: []int{1,2,3,4,5,6},
		},
		{
			name: "already sorted",
			input: []int{1,2,3,4,5,6},
			output: []int{1,2,3,4,5,6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := insertionsort.InsertionSort(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}