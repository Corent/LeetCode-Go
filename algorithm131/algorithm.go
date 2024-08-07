package algorithm131

import (
	"fmt"
	"slices"
)

var (
	dp  [][]bool
	ans [][]string
	cur []string
)

func partition(s string) [][]string {
	dp, ans, cur = make([][]bool, len(s)), make([][]string, 0), make([]string, 0)
	for i := len(s) - 1; i >= 0; i-- {
		dp[i] = make([]bool, len(s))
		for j := i; j < len(s); j++ {
			if j-i < 2 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
		}
	}
	fmt.Println(dp)
	partitionCore(s, 0)
	return ans
}

func partitionCore(s string, from int) {
	if from == len(s) {
		ans = append(ans, slices.Clone[[]string](cur))
		return
	}
	for i := from; i < len(s); i++ {
		if !dp[from][i] {
			continue
		}
		cur = append(cur, s[from:i+1])
		partitionCore(s, i+1)
		cur = cur[:len(cur)-1]
	}
}
