package algorithm390

// https://blog.csdn.net/qq_47875463/article/details/122277421
func lastRemaining(n int) int {
	a1 := 1
	for k, cnt, step := 0, n, 1; cnt > 1; {
		if k&1 == 0 || cnt&1 != 0 {
			a1 += step
		}
		k, cnt, step = k+1, cnt>>1, step<<1
	}
	return a1
}
