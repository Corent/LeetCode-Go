package algorithm416

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum&1 != 0 {
		return false
	}
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, v := range nums {
		for i := sum >> 1; i >= v; i-- {
			dp[i] = dp[i] || dp[i-v]
		}
		if dp[sum>>1] {
			return true
		}
	}
	return false
}
