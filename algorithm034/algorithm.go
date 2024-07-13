package algorithm034

import "fmt"

func searchRange(nums []int, target int) []int {
	fmt.Print(len(nums))
	ans := []int{-1, -1}
	if len(nums) == 0 || target < nums[0] || target > nums[len(nums)-1] {
		return ans
	}
	if target == nums[0] {
		ans[0] = 0
	}
	if target == nums[len(nums)-1] {
		ans[1] = len(nums) - 1
	}
	if ans[0] == -1 {
		i, j := 0, len(nums)-1
		for i < j {
			if j == i+1 {
				if nums[j] == target {
					ans[0] = j
				}
				break
			}
			mid := (i + j) / 2
			if nums[mid] < target {
				i = mid
				continue
			}
			if nums[mid] > target {
				j = mid - 1
				continue
			}
			j = mid
		}
	}
	if ans[0] == -1 {
		return ans
	}
	if ans[1] == -1 {
		i, j := 0, len(nums)-1
		for i < j {
			if j == i+1 {
				if nums[i] == target {
					ans[1] = i
				}
				break
			}
			mid := (i + j) / 2
			if nums[mid] < target {
				i = mid + 1
				continue
			}
			if nums[mid] > target {
				j = mid
				continue
			}
			i = mid
		}
	}
	return ans
}
