package algorithm054

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, m*n)
	xTop, xBottom, yLeft, yRight := 1, m-1, 0, n-1
	x, y, xInc, yInc := 0, 0, 0, 1
	for i := 0; i < m*n; i++ {
		ans[i] = matrix[x][y]
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
	return ans
}
