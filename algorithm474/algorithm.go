package algorithm474

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		cnt0, cnt1 := 0, 0
		for _, c := range str {
			if c == '0' {
				cnt0++
			} else {
				cnt1++
			}
		}
		if cnt0 > m || cnt1 > n {
			continue
		}
		for i := m; i >= cnt0; i-- {
			for j := n; j >= cnt1; j-- {
				if cnt := dp[i-cnt0][j-cnt1] + 1; cnt > dp[i][j] {
					dp[i][j] = cnt
				}
			}
		}
	}
	return dp[m][n]
}
