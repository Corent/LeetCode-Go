package algorithm1092

import "strings"

func shortestCommonSupersequence(str1 string, str2 string) string {
	n, m := len(str1), len(str2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		dp[i][m] = n - i
	}
	for j := 0; j < m; j++ {
		dp[n][j] = m - j
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if str1[i] == str2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = minInt(dp[i+1][j], dp[i][j+1]) + 1
			}
		}
	}
	builder, i, j := &strings.Builder{}, 0, 0
	for i < n && j < m {
		if str1[i] == str2[j] {
			builder.WriteByte(str1[i])
			i, j = i+1, j+1
		} else if dp[i+1][j] == dp[i][j]-1 {
			builder.WriteByte(str1[i])
			i++
		} else if dp[i][j+1] == dp[i][j]-1 {
			builder.WriteByte(str2[j])
			j++
		}
	}
	if i < n {
		builder.WriteString(str1[i:])
	} else if j < m {
		builder.WriteString(str2[j:])
	}
	return builder.String()
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}
