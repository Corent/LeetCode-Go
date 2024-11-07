package algorithm793

// https://blog.51cto.com/muse/6330102

func preimageSizeFZF(k int) int {
	last := 1
	for last < k {
		last = 5*last + 1
	}
	for last > 1 {
		if last-1 == k {
			return 0
		}
		last = (last - 1) / 5
		k %= last
	}
	return 5
}
