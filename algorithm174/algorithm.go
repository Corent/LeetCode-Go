package algorithm174

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m+1)
	for i := m - 1; i >= 0; i-- {
		dp[i] = make([]int, n+1)
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[i][j] = maxInt(1, 1-dungeon[i][j])
			} else if i != m-1 && j == n-1 {
				dp[i][j] = maxInt(1, dp[i+1][j]-dungeon[i][j])
			} else if i == m-1 && j != n-1 {
				dp[i][j] = maxInt(1, dp[i][j+1]-dungeon[i][j])
			} else {
				dp[i][j] = maxInt(1, minInt(dp[i+1][j], dp[i][j+1])-dungeon[i][j])
			}
		}
	}
	return dp[0][0]
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
