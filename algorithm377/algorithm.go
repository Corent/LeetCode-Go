package algorithm377

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1) // dp[i]表示和为i的组合个数
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
