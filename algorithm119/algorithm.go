package algorithm119

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		for j := i; j >= 0; j-- {
			if i == j {
				result[j] = 1
			} else if j > 0 {
				result[j] = result[j-1] + result[j]
			}
		}
	}
	return result
}
