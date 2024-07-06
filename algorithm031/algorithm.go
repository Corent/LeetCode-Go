package algorithm031

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	idx := len(nums) - 2
	for ; idx >= 0 && nums[idx] >= nums[idx+1]; idx-- { // 从后往前找到第一个递减的数字，位置记为p
	}
	if idx == -1 { // 找不到，说明已经完全逆序
		reverse(nums, 0, len(nums)-1)
		return
	}

	target := idx + 1
	for ; target < len(nums) && nums[target] > nums[idx]; target++ { // 从往后找到位置q，使得nums[q+1]是第一个比nums[p]小的数字
	}
	nums[idx], nums[target-1] = nums[target-1], nums[idx] // 交换nums[p] nums[q]
	reverse(nums, idx+1, len(nums)-1)                     // reverse p之后的数组
}

func reverse(nums []int, from, to int) {
	if len(nums) < 2 {
		return
	}
	for i, j := from, to; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
