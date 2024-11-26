package algorithm1310

func xorQueries(arr []int, queries [][]int) []int {
	xors, rst := make([]int, len(arr)), make([]int, len(queries))
	for i := range arr {
		if xors[i] = arr[i]; i > 0 {
			xors[i] ^= xors[i-1]
		}
	}
	for i := range queries {
		if rst[i] = xors[queries[i][1]]; queries[i][0] > 0 {
			rst[i] ^= xors[queries[i][0]-1]
		}
	}
	return rst
}
