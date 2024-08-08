package algorithm189

func rotate(nums []int, k int) {
	if k < 1 || len(nums) < 2 {
		return
	}
	k %= len(nums)
	reverse(nums, 0, len(nums)-k-1)
	reverse(nums, len(nums)-k, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
}

func reverse(nums []int, from, to int) {
	for ; from < to; from, to = from+1, to-1 {
		nums[from], nums[to] = nums[to], nums[from]
	}
}
