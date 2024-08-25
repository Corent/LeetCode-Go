package algorithm365

func canMeasureWater(x int, y int, target int) bool {
	if target > x+y {
		return false
	}
	if target == 0 || target == x || target == y || target == x+y {
		return true
	}
	if x == 0 {
		return y == target
	}
	if y == 0 {
		return x == target
	}
	return target%gcd(x, y) == 0
}

func gcd(x, y int) int {
	if y > x {
		return gcd(y, x)
	}
	for tmp := x % y; tmp != 0; {
		x, y = y, tmp
		tmp = x % y
	}
	return y
}
