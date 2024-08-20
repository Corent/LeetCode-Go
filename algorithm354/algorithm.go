package algorithm354

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	dp := make([][]int, 0, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		lastIdx := len(dp) - 1
		if lastIdx < 0 || dp[lastIdx][1] < envelopes[i][1] {
			dp = append(dp, envelopes[i])
		} else if envelopes[i][1] < dp[0][1] {
			dp[0] = envelopes[i]
		} else {
			dp[binarySearch(dp, envelopes[i])] = envelopes[i]
		}
	}
	return len(dp)
}

func binarySearch(nums [][]int, target []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid][1] == target[1] {
			return mid
		}
		if nums[mid][1] < target[1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// 超时
/*func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] < envelopes[j][1]
	})
	rst, dp := 1, make([]int, len(envelopes))
	dp[0] = 1
	for i := 1; i < len(envelopes); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}
		rst = maxInt(rst, dp[i])
	}
	return rst
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}*/
