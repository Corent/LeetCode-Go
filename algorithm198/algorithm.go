package algorithm198

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) > 2 {
		nums[2] += nums[0]
	}
	for i := 3; i < len(nums); i++ {
		nums[i] += maxInt(nums[i-2], nums[i-3])
	}
	if length := len(nums); length == 1 {
		return nums[0]
	}
	return maxInt(nums[len(nums)-1], nums[len(nums)-2])
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}
