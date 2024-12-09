package algorithm1238

func circularPermutation(n int, start int) []int {
	rst := make([]int, 0)
	for i := 0; i < 1<<n; i++ {
		rst = append(rst, (i>>1)^i^start)
	}
	return rst
}
