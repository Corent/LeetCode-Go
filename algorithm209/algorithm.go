package algorithm209

import "math"

func minSubArrayLen(target int, nums []int) int {
	ans, left, sum := math.MaxInt, 0, 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for left <= right && sum >= target {
			ans = minInt(ans, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}
