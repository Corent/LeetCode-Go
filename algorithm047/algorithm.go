package algorithm047

import (
	"slices"
	"sort"
)

var (
	ans  [][]int
	used []bool
	arrs []int
)

func permuteUnique(nums []int) [][]int {
	ans, used, arrs = make([][]int, 0), make([]bool, len(nums)), make([]int, 0, len(nums))
	if len(nums) == 0 {
		return ans
	}
	sort.Ints(nums)
	permute(nums)
	return ans
}

func permute(nums []int) {
	if len(arrs) == len(nums) {
		ans = append(ans, slices.Clone(arrs))
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		arrs, used[i] = append(arrs, nums[i]), true
		permute(nums)
		arrs, used[i] = arrs[:len(arrs)-1], false
	}
}
