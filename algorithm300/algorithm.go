package algorithm300

import "math"

func lengthOfLIS(nums []int) int {
	dp, ans := make([]int, len(nums)), math.MinInt
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}
		ans = maxInt(ans, dp[i])
	}
	return ans
}

func maxInt(m, n int) int {
	if m < n {
		return n
	}
	return m
}
