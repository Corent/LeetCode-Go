package algorithm287

import "math"

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		n := int(math.Abs(float64(nums[i])))
		if nums[n-1] < 0 {
			return n
		}
		nums[n-1] = -nums[n-1]
	}
	return -1
}
