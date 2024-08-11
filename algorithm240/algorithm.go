package algorithm240

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for i, j := 0, len(matrix[0])-1; i < len(matrix) && j >= 0; {
		if matrix[i][j] == target {
			return true
		}
		if target < matrix[i][j] {
			j--
		} else {
			i++
		}
	}
	return false
}
