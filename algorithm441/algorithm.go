package algorithm441

func arrangeCoins(n int) int {
	m, total := 1, 1
	for total < n {
		m++
		total += m
	}
	if total == n {
		return m
	}
	return m - 1
}
