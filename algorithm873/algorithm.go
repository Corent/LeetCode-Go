package algorithm873

func lenLongestFibSubseq(arr []int) (ans int) {
	n := len(arr)
	indices := make(map[int]int, n)
	for i, x := range arr {
		indices[x] = i
	}
	dp := make([][]int, n) // dp[i][j]表示以i j结尾的斐波那契数列的最大长度
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i, x := range arr {
		for j := n - 1; j >= 0 && arr[j]*2 > x; j-- {
			if k, ok := indices[x-arr[j]]; ok {
				dp[j][i] = maxInt(dp[k][j]+1, 3)
				ans = maxInt(ans, dp[j][i])
			}
		}
	}
	return
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}
