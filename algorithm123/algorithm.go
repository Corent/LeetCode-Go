package algorithm123

func maxProfit(prices []int) int {
	dp1, dp2, minMax, n, ans := make([]int, len(prices)), make([]int, len(prices)), 0, len(prices), 0
	minMax = prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minMax {
			minMax = prices[i]
		}
		if dp1[i] = dp1[i-1]; prices[i]-minMax > dp1[i] {
			dp1[i] = prices[i] - minMax
		}
	}
	minMax = prices[n-1]
	for i := n - 2; i >= 0; i-- {
		if prices[i] > minMax {
			minMax = prices[i]
		}
		if dp2[i] = dp2[i+1]; minMax-prices[i] > dp2[i] {
			dp2[i] = minMax - prices[i]
		}
		if tmp := dp1[i] + dp2[i]; tmp > ans {
			ans = tmp
		}
	}
	return ans
}
