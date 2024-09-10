package algorithm403

func canCross(stones []int) bool {
	if len(stones) == 0 {
		return false
	}
	m := make(map[int]map[int]struct{})
	for _, stone := range stones {
		m[stone] = make(map[int]struct{})
	}
	m[stones[0]][0] = struct{}{}
	for _, stone := range stones {
		lastSteps := m[stone]
		for lastStep := range lastSteps {
			for nextStep := lastStep - 1; nextStep <= lastStep+1; nextStep++ {
				if nextStep <= 0 {
					continue
				}
				if s, ok := m[stone+nextStep]; ok {
					s[nextStep] = struct{}{}
					m[stone+nextStep] = s
				}
			}
		}
	}
	return len(m[stones[len(stones)-1]]) > 0
}
