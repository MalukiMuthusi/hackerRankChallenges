package medians

import "github.com/MalukiMuthusi/hackerRankChallenges/sorting/quicksort"

func RandomizedPartition(A []int, p int, r int) int {
	pv := quicksort.RandomizePivot(A, p, r)
	q := quicksort.Partion(A, p, r, pv)
	return q
}
