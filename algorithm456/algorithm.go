package algorithm456

import "math"

func find132pattern(nums []int) bool {
	// stack保存从后往前遍历的升序序列，栈顶为最大值
	// third为顺序在栈顶元素之后，并且小于栈顶元素的最大值
	stack, third := make([]int, 0), math.MinInt
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < third {
			return true
		}
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			third = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}
