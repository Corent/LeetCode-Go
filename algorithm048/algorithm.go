package algorithm048

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			x, y := make([]int, 4), make([]int, 4)
			x[0], y[0] = i, j
			x[1], y[1] = y[0], n-1-x[0]
			x[2], y[2] = y[1], n-1-x[1]
			x[3], y[3] = y[2], n-1-x[2]
			tmp := matrix[x[3]][y[3]]
			for k := 3; k > 0; k-- {
				matrix[x[k]][y[k]] = matrix[x[k-1]][y[k-1]]
			}
			matrix[x[0]][y[0]] = tmp
		}
	}
}
