package algorithm440

func findKthNumber(n int, k int) int {
	k--
	curr := 1
	for k > 0 {
		step := calStep(n, curr, curr+1)
		if k < step {
			curr, k = curr*10, k-1
		} else {
			curr, k = curr+1, k-step
		}
	}
	return curr
}

func calStep(n, n1, n2 int) int {
	rst := 0
	for n1 <= n {
		rst += minInt(n+1, n2) - n1
		n1, n2 = n1*10, n2*10
	}
	return rst
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}
