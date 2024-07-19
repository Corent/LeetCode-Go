package algorithm044

// https://www.iteye.com/blog/shmilyaw-hotmail-com-2154716

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1) // dp[i][j]表示串s[0...i], p[0...j]这两个串是否匹配
	dp[0] = make([]bool, n+1)
	dp[0][0] = true
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] && p[j-1] == '*'
	}
	for i := 1; i <= m; i++ {
		dp[i] = make([]bool, n+1)
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' { // p[j - 1] != '*'
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j] // p[0...j]以*结尾，结尾的*匹配0个，或者多个
			}
		}
	}
	return dp[m][n]
}
