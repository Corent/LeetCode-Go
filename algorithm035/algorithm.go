package algorithm035

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 || target <= nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	i, j := 0, len(nums)-1
	for i < j {
		if j-i == 1 {
			if nums[i] == target {
				return i
			}
			return j // nums[j] >= target
		}
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			j = mid
			continue
		}
		i = mid
	}
	if nums[i] >= target {
		return i
	}
	return i + 1
}
