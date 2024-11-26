package algorithm1388

import "math"

func maxSizeSlices(slices []int) int {
	return maxInt(maxSizeSlicesCore(slices[1:]), maxSizeSlicesCore(slices[:len(slices)-1]))
}

func maxSizeSlicesCore(slices []int) int {
	N, n := len(slices), (len(slices)+1)/3
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = math.MinInt
		}
	}
	dp[0][0], dp[0][1], dp[1][0], dp[1][1] = 0, slices[0], 0, maxInt(slices[0], slices[1])
	for i := 2; i < N; i++ {
		dp[i][0] = 0
		for j := 1; j <= n; j++ {
			dp[i][j] = maxInt(dp[i-1][j], dp[i-2][j-1]+slices[i])
		}
	}
	return dp[N-1][n]
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}
