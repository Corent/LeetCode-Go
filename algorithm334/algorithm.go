package algorithm334

import "math"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	minVal, secMinVal := math.MaxInt, math.MaxInt
	for _, v := range nums {
		if v <= minVal {
			minVal = v
		} else if v <= secMinVal {
			secMinVal = v
		} else {
			return true
		}
	}
	return false
}
