package algorithm204

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	del := make([]bool, n)
	del[0], del[1], del[2] = true, true, false
	for i := 2; i < n; i++ {
		if del[i] {
			continue
		}
		if i*i > n {
			break
		}
		for j := 2; i*j < n; j++ {
			del[i*j] = true
		}
	}
	cnt := 0
	for _, deleted := range del {
		if !deleted {
			cnt++
		}
	}
	return cnt
}
