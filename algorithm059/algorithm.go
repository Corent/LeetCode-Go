package algorithm059

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	xTop, xBottom, yLeft, yRight := 1, n-1, 0, n-1
	x, y, xInc, yInc := 0, 0, 0, 1
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < n*n; i++ {
		matrix[x][y] = i + 1
		if xInc == 0 && yInc == 1 {
			if y == yRight {
				xInc, yInc, yRight = 1, 0, yRight-1
			}
		} else if xInc == 1 && yInc == 0 {
			if x == xBottom {
				xInc, yInc, xBottom = 0, -1, xBottom-1
			}
		} else if xInc == 0 && yInc == -1 {
			if y == yLeft {
				xInc, yInc, yLeft = -1, 0, yLeft+1
			}
		} else if xInc == -1 && yInc == 0 {
			if x == xTop {
				xInc, yInc, xTop = 0, 1, xTop+1
			}
		}
		x, y = x+xInc, y+yInc
	}
	return matrix
}
