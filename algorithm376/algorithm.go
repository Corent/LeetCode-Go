package algorithm376

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	rst := 1
	dp1 := make([]int, n) // dp1[i]表示以nums[i]为结尾的摆动子序列的最大长度，并且对于i > 0, dp1[i - 1] > dp1[i]
	dp2 := make([]int, n) // dp1[i]表示以nums[i]为结尾的摆动子序列的最大长度，并且对于i > 0, dp1[i - 1] < dp1[i]
	for i := 0; i < n; i++ {
		dp1[i], dp2[i] = 1, 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[i] {
				dp1[i] = maxInt(dp1[i], dp2[j]+1)
			} else if nums[j] < nums[i] {
				dp2[i] = maxInt(dp2[i], dp1[j]+1)
			}
		}
		rst = maxInt(rst, maxInt(dp1[i], dp2[i]))
	}
	return rst
}

func maxInt(m, n int) int {
	if m < n {
		return n
	}
	return m
}
