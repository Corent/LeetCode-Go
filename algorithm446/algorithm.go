package algorithm446

func numberOfArithmeticSlices(nums []int) int {
	n, rst := len(nums), 0
	dp, m := make([][]int, n), make(map[int][]int)
	for i, num := range nums {
		dp[i] = make([]int, n)
		if ids, ok := m[num]; !ok {
			m[num] = []int{i}
		} else {
			m[num] = append(ids, i)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			target := 2*nums[j] - nums[i]
			if ids, ok := m[target]; ok {
				for _, id := range ids {
					if id < j {
						dp[i][j] += dp[j][id] + 1
					}
				}
			}
			rst += dp[i][j]
		}
	}
	return rst
}
