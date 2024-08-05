package algorithm118

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := range result {
		result[i] = make([]int, i+1)
		for j := range result[i] {
			if j == 0 || j == i {
				result[i][j] = 1
			} else if i > 0 {
				result[i][j] = result[i-1][j] + result[i-1][j-1]
			}
		}
	}
	return result
}
