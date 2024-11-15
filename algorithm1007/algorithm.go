package algorithm1007

func minDominoRotations(tops []int, bottoms []int) int {
	numCnt, topCnt, bottomCnt := [6]int{}, [6]int{}, [6]int{}
	for i := 0; i < len(tops); i++ {
		if tops[i] == bottoms[i] {
			numCnt[tops[i]-1]++
			continue
		}
		numCnt[tops[i]-1], numCnt[bottoms[i]-1] = numCnt[tops[i]-1]+1, numCnt[bottoms[i]-1]+1
		topCnt[bottoms[i]-1], bottomCnt[tops[i]-1] = topCnt[bottoms[i]-1]+1, bottomCnt[tops[i]-1]+1
	}
	for i := 0; i < 6; i++ {
		if numCnt[i] == len(tops) {
			return minInt(topCnt[i], bottomCnt[i])
		}
	}
	return 0
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}
