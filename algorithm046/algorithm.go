package algorithm046

import "slices"

var ans [][]int

func permute(nums []int) [][]int {
	ans = make([][]int, 0)
	if len(nums) == 0 {
		return ans
	}
	permuteCore(nums, 0)
	return ans
}

func permuteCore(nums []int, form int) {
	if form == len(nums)-1 {
		ans = append(ans, slices.Clone(nums))
		return
	}
	for i := form; i < len(nums); i++ {
		swap(nums, form, i)
		permuteCore(nums, form+1)
		swap(nums, form, i)
	}
}

func swap(nums []int, i, j int) {
	if i != j {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
