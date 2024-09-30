package algorithm463

func islandPerimeter(grid [][]int) int {
	m, n, maxX, maxY := 0, 0, len(grid), len(grid[0])
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				m++
				if i < maxX-1 && grid[i+1][j] == 1 {
					n++
				}
				if j < maxY-1 && grid[i][j+1] == 1 {
					n++
				}
			}
		}
	}
	return m*4 - n*2
}
