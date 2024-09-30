package algorithm481

func magicalString(n int) int {
	if n == 0 {
		return 0
	}
	if n >= 1 && n <= 3 {
		return 1
	}
	cnt, num, head, tail := 1, 1, 2, 3
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 2, 2
	for tail < n {
		for i := 0; i < dp[head]; i++ {
			dp[tail] = num
			if tail < n && num == 1 {
				cnt++
			}
			tail++
		}
		num, head = num^3, head+1
	}
	return cnt
}
