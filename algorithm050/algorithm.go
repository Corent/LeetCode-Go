package algorithm050

import "math"

func myPow(x float64, n int) float64 {
	if isFloat64Equal(x, 0.0) && n < 0 {
		return 0.0
	}
	absn := n
	if absn < 0 {
		absn = -absn
	}
	rst := myPowWithUnsignedExponent(x, absn)
	if n < 0 {
		return 1.0 / rst
	}
	return rst
}

func myPowWithUnsignedExponent(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	rst := myPowWithUnsignedExponent(base, exponent>>1)
	rst *= rst
	if exponent&1 == 1 {
		rst *= base
	}
	return rst
}

func isFloat64Equal(num1, num2 float64) bool {
	return math.Abs(num1-num2) < 0.0000001
}
