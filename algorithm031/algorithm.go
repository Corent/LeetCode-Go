package algorithm031

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	idx := len(nums) - 2
	for ; idx >= 0 && nums[idx] >= nums[idx+1]; idx-- {
	}
	if idx == -1 {
		reverse(nums, 0, len(nums)-1)
		return
	}

	target := idx + 1
	for ; target < len(nums) && nums[target] > nums[idx]; target++ {
	}
	nums[idx], nums[target-1] = nums[target-1], nums[idx]
	reverse(nums, idx+1, len(nums)-1)
}

func reverse(nums []int, from, to int) {
	if len(nums) < 2 {
		return
	}
	for i, j := from, to; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
