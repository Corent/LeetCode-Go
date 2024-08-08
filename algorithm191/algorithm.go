package algorithm191

func hammingWeight(n int) int {
	cnt := 0
	for n != 0 {
		n = n & (n - 1)
		cnt++
	}
	return cnt
}
