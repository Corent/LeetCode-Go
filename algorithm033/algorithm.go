package algorithm033

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		if nums[left] < nums[right] {
			if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < nums[right] {
				if nums[mid] > target {
					right = mid - 1
				} else {
					left++
				}
			} else {
				if nums[mid] < target {
					left = mid + 1
				} else {
					right--
				}
			}
		}
	}
	return -1
}
