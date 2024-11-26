package algorithm1385

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	rst := 0
	for _, m := range arr1 {
		idx := 0
		for ; idx < len(arr2) && abs(m-arr2[idx]) > d; idx++ {
		}
		if idx == len(arr2) {
			rst++
		}
	}
	return rst
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
