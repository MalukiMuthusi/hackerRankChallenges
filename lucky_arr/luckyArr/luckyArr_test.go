package luckyArr_test

import (
	"testing"

	"github.com/MalukiMuthusi/hackerRankChallenges/lucky_arr/luckyArr"

	"github.com/stretchr/testify/assert"
)

func TestLuckyArr(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output bool
	}{
		{
			name:   "happy case",
			input:  []int{1, 3, 4, 2, 2},
			output: true,
		},
		{
			name:   "long array list",
			input:  []int{-1, -2, 4, 21, 67, -101, 0, 55, 32, 103, -2, 0, 78, 13},
			output: true,
		},
		{
			name:   "array that is not lucky",
			input:  []int{-1, 0, 4, 21, 67, -101, 0, 45, -1, 32, 103, -2, 0, 78, 13},
			output: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := luckyArr.LuckyArr(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
