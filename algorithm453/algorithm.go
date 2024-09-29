package algorithm453

import "math"

func minMoves(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	sum, minVal := 0, math.MaxInt
	for _, num := range nums {
		if sum += num; num < minVal {
			minVal = num
		}
	}
	return sum - minVal*len(nums)
}
