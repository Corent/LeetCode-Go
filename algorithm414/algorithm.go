package algorithm414

import "math"

func thirdMax(nums []int) int {
	num1, num2, num3 := math.MinInt32-1, math.MinInt32-1, math.MinInt32-1
	for _, num := range nums { // nums1 < nums2 < nums3
		if num > num3 {
			num1 = num2
			num2 = num3
			num3 = num
		} else if num < num3 && num > num2 {
			num1 = num2
			num2 = num
		} else if num < num2 && num > num1 {
			num1 = num
		}
	}
	if num1 != math.MinInt32-1 {
		return num1
	}
	return num3
}
