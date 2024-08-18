package algorithm338

func countBits(n int) []int {
	rst := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		cnt := 0
		for j := i; j != 0; {
			j, cnt = j&(j-1), cnt+1
		}
		rst[i] = cnt
	}
	return rst
}
