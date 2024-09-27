package algorithm410

func splitArray(nums []int, k int) int {
	left, right, rst := 0, 0, 0
	for _, v := range nums {
		right += v
	}
	for left <= right {
		mid := left + (right-left)>>1
		if judge(nums, mid, k) {
			rst, right = mid, mid-1
		} else {
			left = mid + 1
		}
	}
	return rst
}

func judge(nums []int, mid, k int) bool {
	sum := 0
	for _, v := range nums {
		if v > mid {
			return false
		}
		if sum+v > mid {
			sum, k = v, k-1
		} else {
			sum += v
		}
	}
	return k >= 1
}
