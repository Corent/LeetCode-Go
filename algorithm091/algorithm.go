package algorithm091

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	if n == 1 {
		if s[0] == '0' {
			return 0
		}
		return 1
	}

	dp := make([]int, n)
	if s[0] != '0' {
		dp[0] = 1
	}
	if s[0] != '0' && s[1] != '0' {
		dp[1] = 1
	}
	if s[0] != '0' && (s[0]-'0')*10+s[1]-'0' <= 26 {
		dp[1]++
	}
	for i := 2; i < n; i++ {
		if s[i] != '0' {
			dp[i] += dp[i-1]
		}
		if s[i-1] != '0' && (s[i-1]-'0')*10+s[i]-'0' <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n-1]
}
