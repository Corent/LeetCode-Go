package algorithm695

func maxAreaOfIsland(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	var area func(r, c int) int
	area = func(r, c int) int {
		if r < 0 || r >= m || c < 0 || c >= n {
			return 0
		}
		if grid[r][c] != 1 {
			return 0
		} // != 0 || != 2
		grid[r][c] = 2 // visited
		return 1 + area(r+1, c) + area(r-1, c) + area(r, c+1) + area(r, c-1)
	}
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			ans = max(ans, area(r, c))
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
