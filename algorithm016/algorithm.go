package algorithm016

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	ans, diff := 0, math.MaxInt
	for i := 0; i < len(nums)-2; i++ {
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if newDiff := abs(sum - target); newDiff < diff {
				diff, ans = newDiff, sum
			}
			if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
