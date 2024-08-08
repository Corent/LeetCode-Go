package algorithm200

func numIslands(grid [][]byte) int {
	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				bfs(grid, i, j)
				cnt++
			}
		}
	}
	return cnt
}

func bfs(grid [][]byte, i, j int) {
	queue, m, n := [][2]int{[2]int{i, j}}, len(grid), len(grid[0])
	grid[i][j] = '2'
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if p[0] > 0 && grid[p[0]-1][p[1]] == '1' {
			queue = append(queue, [2]int{p[0] - 1, p[1]})
			grid[p[0]-1][p[1]] = '2'
		}
		if p[0] < m-1 && grid[p[0]+1][p[1]] == '1' {
			queue = append(queue, [2]int{p[0] + 1, p[1]})
			grid[p[0]+1][p[1]] = '2'
		}
		if p[1] > 0 && grid[p[0]][p[1]-1] == '1' {
			queue = append(queue, [2]int{p[0], p[1] - 1})
			grid[p[0]][p[1]-1] = '2'
		}
		if p[1] < n-1 && grid[p[0]][p[1]+1] == '1' {
			queue = append(queue, [2]int{p[0], p[1] + 1})
			grid[p[0]][p[1]+1] = '2'
		}
	}
}
