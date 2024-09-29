package algorithm452

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	rst, end := 1, points[0][1]
	for _, p := range points {
		if p[0] > end {
			rst, end = rst+1, p[1]
		} else if p[1] < end {
			end = p[1]
		}
	}
	return rst
}
