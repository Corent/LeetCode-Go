package algorithm397

var memo map[int]int

func integerReplacement(n int) int {
	memo = make(map[int]int)
	return integerReplacementCore(n)
}

func integerReplacementCore(n int) (cnt int) {
	if n == 1 {
		return 0
	}
	if c, ok := memo[n]; ok {
		return c
	}
	defer func() {
		memo[n] = cnt
	}()
	if n&1 == 0 {
		return 1 + integerReplacementCore(n/2)
	}
	return 2 + minInt(integerReplacementCore(n/2), integerReplacementCore(n/2+1))
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}
