package algorithm121

func maxProfit(prices []int) int {
	low, high := make([]int, len(prices)), make([]int, len(prices))
	for i, price := range prices {
		if i == 0 || low[i-1] >= price {
			low[i] = price
		} else {
			low[i] = low[i-1]
		}
	}
	ans, n := 0, len(prices)
	for i := n - 1; i >= 0; i-- {
		if i == n-1 || high[i+1] <= prices[i] {
			high[i] = prices[i]
		} else {
			high[i] = high[i+1]
		}
		diff := high[i] - low[i]
		if diff > ans {
			ans = diff
		}
	}
	return ans
}
