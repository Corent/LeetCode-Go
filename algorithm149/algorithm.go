package algorithm149

func maxPoints(points [][]int) int {
	n, idx1, idx2, cnt, maxCnt := len(points), 0, 1, 0, 0
	if n <= 2 {
		return n
	}
	for idx1 < n && idx2 < n {
		if points[idx1][0] == points[idx2][0] {
			for i := 0; i < n; i++ {
				if points[i][0] == points[idx2][0] {
					cnt++
				}
			}
		} else {
			for i := 0; i < n; i++ {
				//if (((long)(points[index1][1] - points[index2][1]) * (points[index1][0] - points[i][0])) == ((long)(points[index1][1] - points[i][1]) * (points[index1][0] - points[index2][0]))){
				if int64(points[idx1][1]-points[idx2][1])*int64(points[idx1][0]-points[i][0]) == int64(points[idx1][1]-points[i][1])*int64(points[idx1][0]-points[idx2][0]) {
					cnt++
				}
			}
		}
		if cnt == n {
			return n
		}
		if maxCnt < cnt {
			maxCnt = cnt
		}
		cnt = 0
		if idx2 < n-1 {
			idx2++
		} else {
			idx1++
			idx2 = idx1 + 1
		}
	}
	return maxCnt
}
