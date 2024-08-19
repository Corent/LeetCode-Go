package algorithm343

import "math"

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1
	for i := 3; i <= n; i++ {
		dp[i] = -1
		for j := 1; j <= i; j++ {
			dp[i] = maxInt(j*dp[i-j], maxInt(dp[i], j*(i-j)))
		}
	}
	return dp[n]
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func integerBreak2(n int) int {
	switch n {
	case 1, 2:
		return 1
	case 3:
		return 2
	}
	switch n % 3 {
	case 0:
		return int(math.Pow(3, float64(n/3)))
	case 1:
		return 2 * 2 * int(math.Pow(3, float64((n-4)/3)))
	}
	return 2 * int(math.Pow(3, float64(n/3)))
}
