package algorithm693

func hasAlternatingBits(n int) bool {
	bit := n & 1
	for n >>= 1; n > 0; n >>= 1 {
		if n&1 != (1 - bit) {
			return false
		}
		bit = 1 - bit
	}
	return true
}
