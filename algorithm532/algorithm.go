package algorithm532

import "sort"

func findPairs(nums []int, k int) int {
	rst, n := 0, len(nums)
	sort.Ints(nums)
	if k == 0 {
		for i := 0; i < n; i++ {
			if i+1 < n && nums[i] == nums[i+1] {
				rst++
				j := i + 1
				for j < n && nums[j] == nums[i] {
					j++
				}
				i = j - 1
			}
		}
		return rst
	}
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}
	for i, num := range nums {
		if i+1 < n && num == nums[i+1] {
			continue
		}
		if _, ok := set[num+k]; ok {
			rst++
		}
		delete(set, num)
	}
	return rst
}
