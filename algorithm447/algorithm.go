package algorithm447

func numberOfBoomerangs(points [][]int) int {
	rst := 0
	m, n := make(map[int]int), len(points)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			disVal := dis(points[i], points[j])
			m[disVal] = m[disVal] + 1
		}
		for k, v := range m {
			if v >= 2 {
				rst += v * (v - 1)
			}
			delete(m, k)
		}
	}
	return rst
}

func dis(a, b []int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}
