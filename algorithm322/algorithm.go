package algorithm322

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if idx := i - coin; idx >= 0 && dp[idx] >= 0 {
				if v := dp[idx] + 1; dp[i] == -1 || v < dp[i] {
					dp[i] = v
				}
			}
		}
	}
	return dp[amount]
}
