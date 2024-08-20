package algorithm357

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	rst, tmp := 10, 9 // n > 1 时，最高位9种选择，即1-9
	for i := 2; i < n+1; i++ {
		tmp *= 9 - (i - 2)
		rst += tmp
	}
	return rst
}
