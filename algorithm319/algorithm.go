package algorithm319

func bulbSwitch(n int) int {
	ans := 0
	for i := 1; i*i <= n; i++ {
		ans++
	}
	return ans
}
