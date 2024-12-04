package algorithm1266

func minTimeToVisitAllPoints(points [][]int) int {
	rst, x0, x1 := 0, points[0][0], points[0][1]
	for i := 1; i < len(points); i++ {
		y0, y1 := points[i][0], points[i][1]
		rst += maxInt(absInt(x0-y0), absInt(x1-y1))
		x0, x1 = y0, y1
	}
	return rst
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
