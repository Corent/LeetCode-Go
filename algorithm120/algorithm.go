package algorithm120

import "math"

func minimumTotal(triangle [][]int) int {
	dp, ans := make([][]int, len(triangle)), math.MaxInt
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
		for j := range dp[i] {
			if i == 0 {
				dp[i][j] = triangle[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == i {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = dp[i-1][j-1]
				if dp[i][j] > dp[i-1][j] {
					dp[i][j] = dp[i-1][j]
				}
				dp[i][j] += triangle[i][j]
			}
			if i == len(triangle)-1 && ans > dp[i][j] {
				ans = dp[i][j]
			}
		}
	}
	return ans
}
