package algorithm007

import "math"

func reverse(x int) int {
	xInt64 := int64(x)
	n, nums, isNegative := 0, make([]int64, 64), false
	if xInt64 < 0 {
		xInt64, isNegative = -xInt64, true
	}
	for xInt64 != 0 {
		nums[n] = xInt64 % 10
		n, xInt64 = n+1, xInt64/10
	}
	ans := int64(0)
	for i := 0; i < n; i++ {
		t := ans
		ans = t*10 + nums[i]
		if t != ans/10 || t*10 > math.MaxInt32 {
			return 0
		}
	}
	if isNegative {
		ans = -ans
	}
	return int(ans)
}
