package algorithm368

import "sort"

func largestDivisibleSubset(nums []int) []int {
	dp, prevs := make([]int, len(nums)), make([]int, len(nums))
	sort.Ints(nums)
	idx, maxLen := -1, 0
	for i := range nums {
		dp[i], prevs[i] = 1, -1
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && dp[i] < dp[j]+1 {
				dp[i], prevs[i] = dp[j]+1, j
			}
		}
		if dp[i] > maxLen {
			maxLen, idx = dp[i], i
		}
	}
	rst := make([]int, maxLen)
	for i := idx; i >= 0; i = prevs[i] {
		rst = append([]int{nums[i]}, rst...)
	}
	return rst
}
