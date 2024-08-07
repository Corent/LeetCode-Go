package algorithm132

func minCut(s string) int {
	memo, n := make([][]bool, len(s)), len(s)
	for i := n - 1; i >= 0; i-- {
		memo[i] = make([]bool, n)
		for j := i; j < n; j++ {
			if j-i < 2 {
				memo[i][j] = s[i] == s[j]
			} else {
				memo[i][j] = memo[i+1][j-1] && s[i] == s[j]
			}
		}
	}
	dp := make([]int, n+1) // dp[i]表示s[0...i]的最少切割次数
	dp[0] = -1
	for i := 1; i <= n; i++ {
		dp[i] = i - 1
		for from := 0; from < i; from++ {
			if memo[from][i-1] && dp[from]+1 < dp[i] {
				dp[i] = dp[from] + 1
			}
		}
	}
	return dp[n]
}

/*var (
	dp  [][]bool
	ans int
)

// 超时
func minCut(s string) int {
	dp, ans = make([][]bool, len(s)), math.MaxInt
	for i := len(s) - 1; i >= 0; i-- {
		dp[i] = make([]bool, len(s))
		for j := i; j < len(s); j++ {
			if j-i < 2 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
		}
	}
	minCutCore(s, 0, 0)
	return ans
}

func minCutCore(s string, from, cnt int) {
	if from == len(s) {
		if cnt-1 < ans {
			ans = cnt - 1
		}
		return
	}
	for i := from; i < len(s); i++ {
		if !dp[from][i] {
			continue
		}
		minCutCore(s, i+1, cnt+1)
	}
}
*/
