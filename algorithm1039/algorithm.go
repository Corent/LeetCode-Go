package algorithm1039

import "math"

func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for length := 3; length <= n; length++ {
		for i := 0; i < n; i++ {
			j, tmp := i+length-1, math.MaxInt
			for k := i + 1; k < j; k++ {
				tmp = minInt(tmp, values[i]*values[j%n]*values[k%n]+dp[i][k%n]+dp[k%n][j%n])
			}
			dp[i][j%n] = tmp
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
