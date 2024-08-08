package algorithm162

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		if l == r {
			return l
		}
		mid := l + (r-l)>>1
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return -1
}
