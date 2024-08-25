package algorithm363

import "math"

func maxSumSubmatrix(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	for i := 1; i <= m; i++ {
		sum[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			sum[i][j] = matrix[i-1][j-1] + sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1]
		}
	}
	rst := math.MinInt
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for g := 0; g < i; g++ {
				for h := 0; h < j; h++ {
					if one := sum[i][j] - (sum[i][h] + sum[g][j] - sum[g][h]); one <= k && one > rst {
						rst = one
					}
				}
			}
		}
	}
	return rst
}
