package algorithm650

func minSteps(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = i
		for k, j := 2, i/2; j >= 1; k, j = k+1, i/(k+1) {
			if i%j == 0 {
				dp[i] = minInt(dp[i], dp[j]+i/j)
			}
		}
	}
	return dp[n]
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}
