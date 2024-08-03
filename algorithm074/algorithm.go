package algorithm074

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	low, high := 0, m*n-1
	for low <= high {
		mid := low + (high-low)>>1
		i, j := mid/n, mid%n
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
