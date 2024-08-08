package algorithm153

func findMin(nums []int) int {
	mid := 0
	for down, up := 0, len(nums)-1; down < up; {
		if nums[down] < nums[up] {
			mid = down
			break
		}
		if up-mid == 1 {
			mid = up
			break
		}

		mid = (down + up) >> 1
		if nums[mid] >= nums[down] {
			down = mid
		} else if nums[mid] <= nums[up] {
			up = mid
		}
	}
	return nums[mid]
}
