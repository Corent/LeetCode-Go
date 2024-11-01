package algorithm1297

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	m := make(map[string]int)
	for i := 0; i <= len(s)-minSize; i++ {
		str := s[i : i+minSize]
		count := 0
		for j := 0; j < len(str); j++ {
			if 1<<int(str[j]-'a')&count == 0 {
				count = count | 1<<int(str[j]-'a')

			}
		}
		sum := 0
		for j := 0; j < 32; j++ {
			if (1<<j)&count != 0 {
				sum++
			}
		}

		if sum <= maxLetters {
			_, ok := m[str]
			if !ok {
				m[str] = 1
			} else {
				m[str] = m[str] + 1
			}
		}
	}
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}
