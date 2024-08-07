package algorithm139

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	m, wordSet := make([][]bool, n), make(map[string]struct{})
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}
	for i := 0; i < n; i++ {
		m[i] = make([]bool, n)
		for j := i; j < n; j++ {
			_, m[i][j] = wordSet[s[i:j+1]]
		}
	}

	dp := make([]bool, n)
	for j := 0; j < n; j++ {
		if j == 0 {
			dp[j] = m[0][0]
			continue
		}
		if m[0][j] {
			dp[j] = true
			continue
		}
		for i := 0; i < j; i++ {
			if dp[i] && m[i+1][j] {
				dp[j] = true
				break
			}
		}
	}
	return dp[n-1]
}
