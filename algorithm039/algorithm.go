package algorithm039

import (
	"slices"
)

var (
	ans  [][]int
	sums []int
)

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return ans
	}
	ans, sums = [][]int{}, []int{}
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
		sums = append(sums, candidates[i])
		combinationSumHelper(candidates, i, target-candidates[i])
		sums = sums[:len(sums)-1]
	}
}
