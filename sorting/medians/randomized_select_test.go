package medians_test

import (
	"testing"

	"github.com/MalukiMuthusi/hackerRankChallenges/sorting/medians"
	"github.com/stretchr/testify/assert"
)

func TestRandomizedSelect(t *testing.T) {

	type input struct {
		A []int
		p int
		r int
		i int
	}

	tests := []struct {
		name   string
		input  input
		output int
	}{
		{
			name: "happy case",
			input: input{
				A: []int{8, 2, 5, 1, 7, 4, 6, 0, 9, 3},
				p: 0,
				r: 9,
				i: 5,
			},
			output: 4,
		},
		{
			name: "smallest value",
			input: input{
				A: []int{8, 2, 5, 1, 7, 4, 6, 0, 9, 3},
				p: 0,
				r: 9,
				i: 1,
			},
			output: 0,
		},
		{
			name: "largest value",
			input: input{
				A: []int{8, 2, 5, 1, 7, 4, 6, 0, 9, 3},
				p: 0,
				r: 9,
				i: 10,
			},
			output: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := medians.RandomizedSelect(tt.input.A, tt.input.p, tt.input.r, tt.input.i)

			assert.Equal(t, tt.output, n)
		})
	}
}
