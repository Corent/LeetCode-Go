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

/**
 * O(nlog(n))
 * https://www.bilibili.com/video/BV1Wf4y1y7ou
 * dp[i]表示长度为i+1的子序列，尾数的最优解
 */

func lengthOfLISv2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		lastIdx := len(dp) - 1
		if lastIdx < 0 || nums[i] > dp[lastIdx] {
			dp = append(dp, nums[i])
		} else if nums[i] < dp[0] {
			dp[0] = nums[i]
		} else {
			dp[binarySearch(dp, nums[i])] = nums[i]
		}
	}
	return len(dp)
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
