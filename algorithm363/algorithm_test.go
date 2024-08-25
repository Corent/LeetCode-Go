package algorithm363

import "testing"

func TestMaxSumSubmatrix(t *testing.T) {
	matrix := [][]int{{1, 0, 1}, {0, -2, 3}}
	t.Log(maxSumSubmatrix(matrix, 2))
}
