package algorithm600

import "fmt"

// http://www.mamicode.com/info-detail-2614808.html
func findIntegers(n int) int {
	if n == 0 {
		return 1
	}
	s := fmt.Sprintf("%b", n)
	dp := make([]int, len(s)+1)
	dp[0], dp[1] = 1, 2
	for i := 2; i <= len(s); i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	sum := 0
	for i, k := 0, len(s); i < len(s); i, k = i+1, k-1 {
		if s[i] == '1' {
			sum += dp[k-1]
			if i > 0 && s[i-1] == '1' {
				return sum
			}
		}
	}
	sum++ // n本身
	return sum
}
