package algorithm1335

import "math"

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}
	dp := make([][]int, d+1)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}
	maxVal := 0
	for i := range jobDifficulty {
		maxVal = maxInt(maxVal, jobDifficulty[i])
		dp[0][i] = maxVal
	}
	for i := 1; i < d; i++ {
		for j := i; j < n; j++ {
			maxVal = 0
			for k := j; k >= i; k-- {
				maxVal = maxInt(maxVal, jobDifficulty[k])
				dp[i][j] = minInt(dp[i][j], maxVal+dp[i-1][k-1])
			}
		}
	}
	return dp[d-1][n-1]
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func minInt(m, n int) int {
	if m > n {
		return n
	}
	return m
}
