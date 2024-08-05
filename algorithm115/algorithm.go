package algorithm115

// https://blog.csdn.net/laplacepoisson/article/details/124290860

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	dp := make([][]int, m+1) // dp[i][j]表示s[0...i]和t[0...j]的结果
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			if j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j]
				if s[i-1] == t[j-1] {
					dp[i][j] += dp[i-1][j-1]
				}
			}
		}
	}
	return dp[m][n]
}
