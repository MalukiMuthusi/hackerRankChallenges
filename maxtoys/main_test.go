package maxtoys_test

import (
	"testing"

	maxtoys "github.com/MalukiMuthusi/hackerRankChallenges/max_toys"
	"github.com/stretchr/testify/assert"
)

func TestMaximumToys(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		k      int
		output int
	}{
		// {
		// 	name:   "happy case",
		// 	prices: []int{1, 12, 5, 111, 200, 1000, 10},
		// 	k:      50,
		// 	output: 4,
		// },
		// {
		// 	name:   "small numbers",
		// 	prices: []int{1, 2, 3, 4},
		// 	k:      7,
		// 	output: 3,
		// },
		// {
		// 	name:   "small numbers random",
		// 	prices: []int{3, 7, 2, 9, 4},
		// 	k:      15,
		// 	output: 3,
		// },

		{
			name:   "large values",
			prices: []int{33324560, 77661073, 31948330, 21522343, 97176507, 5724692, 24699815, 12079402, 6479353, 28430129, 42427721, 57127004, 26256001, 29446837, 65107604, 9809008, 65846182, 8470661, 13597655, 360},
			k:      100000,
			output: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := maxtoys.MaximumToys(tt.prices, tt.k)
			assert.Equal(t, tt.output, ans)
		})
	}
}
