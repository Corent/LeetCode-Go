package algorithm304

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sums := make([][]int, len(matrix))
	for i := range sums {
		sums[i] = make([]int, len(matrix[0]))
		for j := range sums[i] {
			if i == 0 && j == 0 {
				sums[i][j] = matrix[i][j]
			} else if i == 0 {
				sums[i][j] = sums[i][j-1] + matrix[i][j]
			} else if j == 0 {
				sums[i][j] = sums[i-1][j] + matrix[i][j]
			} else {
				sums[i][j] = sums[i-1][j] + sums[i][j-1] - sums[i-1][j-1] + matrix[i][j]
			}
		}
	}
	return NumMatrix{sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 {
		return this.sums[row2][col2]
	}
	if row1 == 0 {
		return this.sums[row2][col2] - this.sums[row2][col1-1]
	}
	if col1 == 0 {
		return this.sums[row2][col2] - this.sums[row1-1][col2]
	}
	return this.sums[row2][col2] - this.sums[row1-1][col2] - this.sums[row2][col1-1] + this.sums[row1-1][col1-1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
