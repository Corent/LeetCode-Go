package algorithm462

import "sort"

func minMoves2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	target, cnt := nums[len(nums)/2], 0
	for _, num := range nums {
		cnt += absInt(num - target)
	}
	return cnt
}

func absInt(m int) int {
	if m < 0 {
		return -m
	}
	return m
}
