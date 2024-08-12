package algorithm279

import "math"

func numSquares(n int) int {
	squareNums := make([]int, int(math.Sqrt(float64(n))+1))
	for i := range squareNums {
		squareNums[i] = (i + 1) * (i + 1)
	}
	level, queue := 0, []int{n}
	for len(queue) > 0 {
		level++
		nextQueue := make([]int, 0)
		for _, m := range queue {
			for _, sn := range squareNums {
				if m == sn {
					return level
				} else if m < sn {
					break
				} else {
					nextQueue = append(nextQueue, m-sn)
				}
			}
		}
		queue = nextQueue
	}
	return level
}
