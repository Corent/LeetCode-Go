package algorithm342

func isPowerOfFour(n int) bool {
	if n < 1 || n&(n-1) != 0 {
		return false
	}
	cnt := -1
	for n != 0 {
		n, cnt = n>>1, cnt+1
	}
	return cnt&1 == 0
}
