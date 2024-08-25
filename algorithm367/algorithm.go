package algorithm367

func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left < right {
		mid := (left + right) >> 1
		if m := mid * mid; m == num {
			return true
		} else if m < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left*left == num
}
