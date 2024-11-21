package algorithm1035

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	dp[0] = make([]int, len(nums2)+1)
	for i := 1; i <= len(nums1); i++ {
		dp[i] = make([]int, len(nums2)+1)
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else if dp[i][j] = dp[i-1][j]; dp[i][j-1] > dp[i-1][j] {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(nums1)][len(nums2)]
}
