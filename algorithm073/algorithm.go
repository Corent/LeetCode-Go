package algorithm073

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	firstRow0, firstCol0 := false, false
	for j := 0; j < n; j++ {
		if firstRow0 = firstRow0 || matrix[0][j] == 0; firstRow0 {
			break
		}
	}
	for i := 0; i < m; i++ {
		if firstCol0 = firstCol0 || matrix[i][0] == 0; firstCol0 {
			break
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j], matrix[i][0] = 0, 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if firstRow0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if firstCol0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
