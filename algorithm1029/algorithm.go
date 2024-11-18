package algorithm1029

func twoCitySchedCost(costs [][]int) int {
	left, right, n := 0, len(costs)-1, len(costs)>>1
	for left < right {
		i, j := left, right
		for i < j {
			for i < j && compare(costs[i], costs[j]) <= 0 {
				j--
			}
			costs[i], costs[j] = costs[j], costs[i]
			for i < j && compare(costs[i], costs[j]) <= 0 {
				i++
			}
			costs[i], costs[j] = costs[j], costs[i]
		}
		if i == n {
			break
		}
		if i < n {
			left = i + 1
		} else {
			right = i - 1
		}
	}
	rst := 0
	for i, cost := range costs {
		if i < n {
			rst += cost[0]
		} else {
			rst += cost[1]
		}
	}
	return rst
}

func compare(a, b []int) int {
	return (a[0] - a[1]) - (b[0] - b[1])
}
