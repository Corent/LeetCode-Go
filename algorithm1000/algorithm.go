package algorithm1000

import "math"

func mergeStones(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + stones[i-1]
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for size := 2; size <= n; size++ {
		for i := 0; i+size <= n; i++ {
			j := i + size - 1
			dp[i][j] = math.MaxInt
			for m := i; m < j; m += k - 1 {
				dp[i][j] = minInt(dp[i][j], dp[i][m]+dp[m+1][j])
			}
			if (size-1)%(k-1) == 0 {
				dp[i][j] += sum[j+1] - sum[i]
			}
		}
	}
	return dp[0][n-1]
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}
