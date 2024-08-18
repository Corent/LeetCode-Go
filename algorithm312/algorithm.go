package algorithm312

func maxCoins(nums []int) int {
	n := len(nums) + 2
	arr := make([]int, n)
	for i := 1; i < n-1; i++ {
		arr[i] = nums[i-1]
	}
	arr[0], arr[n-1] = 1, 1
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for gap := 2; gap < n; gap++ {
		for left := 0; left < n-gap; left++ {
			right := left + gap
			for i := left + 1; i < right; i++ {
				dp[left][right] = maxInt(dp[left][right], dp[left][i]+dp[i][right]+arr[left]*arr[i]*arr[right])
			}
		}
	}
	return dp[0][n-1]
}

func maxInt(m, n int) int {
	if m < n {
		return n
	}
	return m
}
