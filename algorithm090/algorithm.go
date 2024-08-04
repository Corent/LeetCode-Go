package algorithm090

import (
	"slices"
	"sort"
)

var (
	NUMS, current []int
	ans           [][]int
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	NUMS, current, ans = nums, make([]int, 0, len(nums)), make([][]int, 0)
	for k := 0; k <= len(nums); k++ {
		pick(0, k)
	}
	return ans
}

func pick(from, k int) {
	if len(NUMS)-from < k {
		return
	}
	if k == 0 {
		ans = append(ans, slices.Clone[[]int](current))
		return
	}
	for i := from; i < len(NUMS); i++ {
		if i != from && NUMS[i-1] == NUMS[i] {
			continue
		}
		current = append(current, NUMS[i])
		pick(i+1, k-1)
		current = current[:len(current)-1]
	}
}
