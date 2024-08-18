package algorithm309

import "math"

func maxProfit(prices []int) int {
	s0, s1, s2 := 0, -prices[0], math.MinInt
	for i := 1; i < len(prices); i++ {
		pre0, pre1, pre2 := s0, s1, s2
		s0 = maxInt(pre0, pre2)
		s1 = maxInt(pre0-prices[i], pre1)
		s2 = pre1 + prices[i]
	}
	return maxInt(s0, s2)
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}
