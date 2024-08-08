package algorithm154

func findMin(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	mid := 0
	for i, j := 0, n-1; nums[i] >= nums[j]; {
		if j-i == 1 {
			mid = j
			break
		}
		mid = (i + j) >> 1
		if nums[mid] > nums[j] {
			i = mid
		} else if nums[mid] < nums[j] {
			j = mid
		} else {
			j--
		}
	}
	return nums[mid]
}
