package algorithm135

func candy(ratings []int) int {
	n, ans := len(ratings), 0
	candies := make([]int, n)

	candies[0] = 1
	for i := 1; i < n; i++ {
		if candies[i] = 1; ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	ans += candies[n-1]
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if tmp := candies[i+1] + 1; tmp > candies[i] {
				candies[i] = tmp
			}
		}
		ans += candies[i]
	}
	return ans
}
