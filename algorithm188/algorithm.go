package algorithm188

func maxProfit(k int, prices []int) int {
	if k < 1 || len(prices) == 0 {
		return 0
	}
	n := len(prices)
	if k >= n {
		return maxProfitNoLimit(prices)
	}
	// local[i]表示总交易次数为i，并且在最后一天要做交易的情况下的最大获益
	// global[i]表示总交易次数为i的最大获益
	local, global := make([]int, k+1), make([]int, k+1)
	for i := 1; i < n; i++ {
		diff := prices[i] - prices[i-1]
		for j := k; j > 0; j-- {
			local[j] = maxInt(global[j-1], local[j]+diff)
			global[j] = maxInt(global[j], local[j])
		}
	}
	return global[k]
}

// algorithm122
func maxProfitNoLimit(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}
