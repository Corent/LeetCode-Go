package algorithm263

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n&1 == 0 {
		n >>= 1
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	return n == 1
}
