package algorithm1201

import "math"

func nthUglyNumber(n int, a int, b int, c int) int {
	ab := a / gcd(a, b) * b
	ac := a / gcd(a, c) * c
	bc := b / gcd(b, c) * c
	abc := ab / gcd(ab, c) * c
	left, right := 1, math.MaxInt
	for left < right {
		mid := left + (right-left)>>1
		cnt := mid/a + mid/b + mid/c - mid/ab - mid/ac - mid/bc + mid/abc
		if cnt >= n {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func gcd(x, y int) int {
	if y > x {
		return gcd(y, x)
	}
	for z := x % y; z != 0; z = x % y {
		x, y = y, z
	}
	return y
}
