package algorithm329

var (
	m, n    int
	dp, mtx [][]int
)

func longestIncreasingPath(matrix [][]int) int {
	m, n, mtx = len(matrix), len(matrix[0]), matrix
	dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = maxInt(ans, dfs(i, j))
		}
	}
	return ans
}

func dfs(x, y int) int {
	if dp[x][y] > 0 {
		return dp[x][y]
	}
	nextLen := 0
	if x > 0 && mtx[x-1][y] > mtx[x][y] {
		nextLen = maxInt(nextLen, dfs(x-1, y))
	}
	if x < m-1 && mtx[x+1][y] > mtx[x][y] {
		nextLen = maxInt(nextLen, dfs(x+1, y))
	}
	if y > 0 && mtx[x][y-1] > mtx[x][y] {
		nextLen = maxInt(nextLen, dfs(x, y-1))
	}
	if y < n-1 && mtx[x][y+1] > mtx[x][y] {
		nextLen = maxInt(nextLen, dfs(x, y+1))
	}
	dp[x][y] = nextLen + 1
	return dp[x][y]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
