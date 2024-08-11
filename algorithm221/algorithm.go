package algorithm221

import "math"

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp, ans := make([][]int, m), math.MinInt
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = int(matrix[i][j] - '0')
			} else if matrix[i][j] == '1' {
				dp[i][j] = minInt(minInt(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			ans = maxInt(ans, dp[i][j])
		}
	}
	return ans * ans
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func maxInt(m, n int) int {
	if m < n {
		return n
	}
	return m
}
