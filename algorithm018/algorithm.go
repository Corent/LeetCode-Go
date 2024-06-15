package algorithm018

import "sort"

func fourSum(nums []int, target int) [][]int {
	rst := make([][]int, 0)
	if len(nums) < 4 {
		return rst
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		threeSums := threeSum(nums, i+1, target-nums[i])
		if len(threeSums) == 0 {
			continue
		}
		for _, t := range threeSums {
			rst = append(rst, append([]int{nums[i]}, t...))
		}
	}
	return rst
}

func threeSum(nums []int, from, target int) [][]int {
	rst := make([][]int, 0)
	for i := from; i < len(nums)-2; i++ {
		if i > from && nums[i] == nums[i-1] {
			continue
		}
		twoSums := twoSum(nums, i+1, target-nums[i])
		if len(twoSums) == 0 {
			continue
		}
		for _, t := range twoSums {
			rst = append(rst, append([]int{nums[i]}, t...))
		}
	}
	return rst
}

func twoSum(nums []int, from, target int) [][]int {
	rst := make([][]int, 0)
	if len(nums) < 2 {
		return rst
	}
	for i, j := from, len(nums)-1; i < j; {
		if sum := nums[i] + nums[j]; sum == target {
			rst = append(rst, []int{nums[i], nums[j]})
			for i < j && nums[i] == nums[i+1] {
				i++
			}
			for i < j && nums[j] == nums[j-1] {
				j--
			}
			i, j = i+1, j-1
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return rst
}
