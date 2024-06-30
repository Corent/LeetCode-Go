package algorithm029

import "math"

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return math.MinInt
	}
	if dividend == math.MinInt && divisor == -1 {
		return math.MaxInt
	}
	if divisor == 1 || dividend == 0 {
		return dividend
	}
	ans, isNegative := 0, (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0)
	dend, sor := int64(dividend), int64(divisor)
	if dividend < 0 {
		dend = -dend
	}
	if divisor < 0 {
		sor = -sor
	}
	for dend >= sor {
		digit, tmp := 1, sor<<1
		for dend > tmp {
			digit <<= 1
			tmp <<= 1
		}
		ans += digit
		tmp >>= 1
		dend -= tmp
	}
	if isNegative {
		ans = -ans
	}
	if ans > math.MaxInt32 {
		ans = math.MaxInt32
	}
	if ans < math.MinInt32 {
		ans = math.MinInt32
	}
	return ans
}
