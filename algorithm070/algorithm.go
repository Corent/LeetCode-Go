package algorithm070

func climbStairs(n int) int {
	if n < 4 {
		return n
	}
	pre1, pre2 := 3, 5
	for i := 4; i < n; i++ {
		next := pre1 + pre2
		pre1, pre2 = pre2, next
	}
	return pre2
}
