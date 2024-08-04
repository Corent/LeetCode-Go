package algorithm089

import "math"

func grayCode(n int) []int {
	cnt := int(math.Pow(2, float64(n)))
	ans := make([]int, cnt)
	for i := range ans {
		ans[i] = (i >> 1) ^ i
	}
	return ans
}
