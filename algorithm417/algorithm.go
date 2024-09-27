package algorithm417

var (
	m, n int
	h    [][]int
	dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
)

func pacificAtlantic(heights [][]int) [][]int {
	m, n, h = len(heights), len(heights[0]), heights
	toPacific, toAtlantic := make([][]bool, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		toPacific[i], toAtlantic[i] = make([]bool, n), make([]bool, n)
	}
	for i := 0; i < m; i++ {
		dfs(i, 0, toPacific)
		dfs(i, n-1, toAtlantic)
	}
	for j := 0; j < n; j++ {
		dfs(0, j, toPacific)
		dfs(m-1, j, toAtlantic)
	}
	rst := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if toPacific[i][j] && toAtlantic[i][j] {
				rst = append(rst, []int{i, j})
			}
		}
	}
	return rst
}

func dfs(i, j int, canReach [][]bool) {
	if canReach[i][j] {
		return
	}
	canReach[i][j] = true
	for _, d := range dirs {
		nextI, nextJ := i+d[0], j+d[1]
		if nextI < 0 || nextI >= m || nextJ < 0 || nextJ >= n || h[i][j] > h[nextI][nextJ] {
			continue
		}
		dfs(nextI, nextJ, canReach)
	}
}
