package algorithm040

import (
	"slices"
	"sort"
)

var (
	ans  [][]int
	sums []int
)

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return ans
	}
	ans, sums = [][]int{}, []int{}
	sort.Ints(candidates)
	combinationSumHelper(candidates, 0, target)
	return ans
}

func combinationSumHelper(candidates []int, idx, target int) {
	if target <= 0 {
		if target == 0 {
			ans = append(ans, slices.Clone(sums))
		}
		return
	}
	for i := idx; i < len(candidates); i++ {
		if i > idx && candidates[i] == candidates[i-1] {
			continue
		}
		sums = append(sums, candidates[i])
		combinationSumHelper(candidates, i+1, target-candidates[i])
		sums = sums[:len(sums)-1]
	}
}
