package algorithm393

func validUtf8(data []int) bool {
	for i := 0; i < len(data); {
		if data[i] > 247 {
			return false
		}
		cnt := prefix1Cnt(data[i])
		if cnt < 0 || cnt == 1 {
			return false
		}
		if cnt == 0 {
			i++
			continue
		}
		// 2 3 4
		if i+cnt-1 >= len(data) {
			return false
		}
		for j := i + 1; j <= i+cnt-1; j++ {
			if prefix1Cnt(data[j]) != 1 {
				return false
			}
		}
		i = i + cnt
	}
	return true
}

func prefix1Cnt(n int) int {
	if n>>7 == 0 {
		return 0
	}
	if n>>6 == 2 {
		return 1
	}
	if n>>5 == 6 {
		return 2
	}
	if n>>4 == 14 {
		return 3
	}
	if n>>3 == 30 {
		return 4
	}
	return -1
}
