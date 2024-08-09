package algorithm202

func isHappy(n int) bool {
	set := map[int]struct{}{n: {}}
	for n != 1 {
		next := 0
		for n > 0 {
			m := n % 10
			n /= 10
			next += m * m
		}
		n = next
		if _, ok := set[n]; ok {
			break
		}
		set[n] = struct{}{}
	}
	return n == 1
}
