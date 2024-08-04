package algorithm096

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for leftCnt, rightCnt := 0, i-1; leftCnt <= rightCnt; leftCnt, rightCnt = leftCnt+1, rightCnt-1 {
			times := 2
			if leftCnt == rightCnt {
				times = 1
			}
			dp[i] += dp[leftCnt] * dp[rightCnt] * times
		}
	}
	return dp[n]
}

// 超时
/*func numTrees(n int) int {
	return numTreesCore(1, n)
}

func numTreesCore(from, to int) int {
	if from > to {
		return 1
	}
	cnt := 0
	for root := from; root <= to; root++ {
		leftCnt, rightCnt := numTreesCore(from, root-1), numTreesCore(root+1, to)
		cnt += leftCnt * rightCnt
	}
	return cnt
}*/
