package algorithm473

import "sort"

func makesquare(matchsticks []int) bool {
	if len(matchsticks) < 4 {
		return false
	}
	sum := 0
	for _, stick := range matchsticks {
		sum += stick
	}
	if sum%4 != 0 {
		return false
	}
	sums := [4]int{}
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks))) // 减少搜索量
	return helper(matchsticks, sums, 0, sum/4)
}

func helper(nums []int, sums [4]int, pos, target int) bool {
	if pos == len(nums) {
		return true
	}
	for i := 0; i < 4; i++ {
		if sums[i]+nums[pos] > target {
			continue
		}
		sums[i] += nums[pos]
		if helper(nums, sums, pos+1, target) {
			return true
		}
		sums[i] -= nums[pos]
	}
	return false
}
