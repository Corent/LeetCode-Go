package algorithm375

import "math"

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := i - 1; j > 0; j-- {
			if j+1 == i {
				dp[j][i] = j
				continue
			}
			dp[j][i] = math.MaxInt
			for k := j + 1; k < i; k++ {
				dp[j][i] = minInt(dp[j][i], k+maxInt(dp[j][k-1], dp[k+1][i]))
			}
		}
	}
	return dp[1][n]
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
