package algorithm827

func largestIsland(grid [][]int) (ans int) {
	dirs := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n, t, tags := len(grid), 0, make([][]int, len(grid)) // t表示岛屿编号 tags[i][j]表示grid[i][j]所属的岛屿
	for i := range tags {
		tags[i] = make([]int, n)
	}
	areas := map[int]int{} // areas[t]表示编号为t的岛屿的面积
	var dfs func(int, int)
	dfs = func(i, j int) {
		tags[i][j] = t
		areas[t]++
		for _, d := range dirs {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < n && 0 <= y && y < n && grid[x][y] > 0 && tags[x][y] == 0 {
				dfs(x, y)
			}
		}
	}
	for i, row := range grid {
		for j, x := range row {
			if x > 0 && tags[i][j] == 0 { // 枚举没有访问过的陆地
				t = i*n + j + 1
				dfs(i, j)
				ans = maxInt(ans, areas[t])
			}
		}
	}

	for i, row := range grid {
		for j, x := range row {
			if x == 0 { // 枚举可以添加陆地的位置
				newArea, conn := 1, map[int]bool{0: true}
				for _, d := range dirs {
					x, y := i+d.x, j+d.y
					if 0 <= x && x < n && 0 <= y && y < n && !conn[tags[x][y]] {
						newArea += areas[tags[x][y]]
						conn[tags[x][y]] = true
					}
				}
				ans = maxInt(ans, newArea)
			}
		}
	}
	return
}

func maxInt(a, b int) int {
	if b > a {
		return b
	}
	return a
}
