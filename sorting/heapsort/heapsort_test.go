package heapsort_test

import (
	"testing"

	"github.com/MalukiMuthusi/hackerRankChallenges/sorting/heapsort"
	"github.com/stretchr/testify/assert"
)

func TestLeft(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output int
		err    bool
	}{
		{
			name:   "happy case",
			input:  3,
			output: 6,
		},
		{
			name:   "zero input",
			input:  0,
			output: 0,
		},
		{
			name:   "negative input",
			input:  -2,
			output: 0,
			err:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := heapsort.Left(tt.input)
			if tt.err {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.output, res)
		})
	}
}

func TestRight(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output int
		err    bool
	}{
		{
			name:   "happy case",
			input:  3,
			output: 7,
		},
		{
			name:   "zero input",
			input:  0,
			output: 1,
		},
		{
			name:   "negative input",
			input:  -3,
			output: 0,
			err:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := heapsort.Right(tt.input)
			if tt.err {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.output, res)
		})
	}
}

func TestParent(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output int
		err    bool
	}{
		{
			name:   "happy case",
			input:  6,
			output: 3,
		},
		{
			name:   "odd number",
			input:  7,
			output: 3,
		},
		{
			name:   "negative number",
			input:  -6,
			output: 0,
			err:    true,
		},
		{
			name:   "very large number",
			input:  11298192,
			output: 5649096,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := heapsort.Parent(tt.input)
			if tt.err {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.output, res)
		})
	}
}

func TestMaxHeapify(t *testing.T) {
	type Input struct {
		a []int
		i int
	}
	tests := []struct {
		name   string
		input  Input
		output []int
	}{
		{
			name: "happy case",
			input: Input{
				a: []int{0, 16, 4, 10, 14, 7, 9, 3, 2, 8, 1},
				i: 2,
			},
			output: []int{0, 16, 14, 10, 8, 7, 9, 3, 2, 4, 1},
		},
	}

	for _, tt := range tests {
		heapsort.MaxHeapify(tt.input.a, tt.input.i)
		assert.Equal(t, tt.output, tt.input.a, tt.name)
	}
}

func TestBuildMaxHeap(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "happy case",
			input:  []int{0, 2, 4, 10, 14, 7, 9, 3, 16, 8, 1},
			output: []int{0, 16, 14, 10, 8, 7, 9, 3, 4, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := heapsort.BuildMaxHeap(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
