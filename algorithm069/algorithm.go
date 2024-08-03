package algorithm069

func mySqrt(x int) int {
	s, e := int64(1), int64(x)
	for s+1 < e {
		m := (s + e) >> 1
		m2 := m * m
		if m2 > int64(x) {
			e = m
		} else if m2 < int64(x) {
			s = m
		} else {
			return int(m)
		}
	}
	if e*e > int64(x) {
		return int(s)
	}
	return int(e)
}
