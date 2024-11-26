package algorithm1234

var (
	avg          int
	sHash, wHash [4]int
)

func balancedString(s string) int {
	avg, sHash, wHash = len(s)/4, [4]int{}, [4]int{}
	for i := range s {
		sHash[idx(s[i])]++
	}
	rst := len(s)
	for left, right := 0, 0; right < len(s); right++ {
		wHash[idx(s[right])]++
		if !check() {
			continue
		}
		rst = minInt(rst, right-left+1)
		for left <= right {
			wHash[idx(s[left])]--
			left++
			if !check() {
				break
			}
			rst = minInt(rst, right-left+1)
		}
	}
	return rst
}

func check() bool {
	for i := 0; i < 4; i++ {
		if sHash[i]-wHash[i] > avg {
			return false
		}
	}
	return true
}

func idx(ch byte) int {
	switch ch {
	case 'Q':
		return 0
	case 'W':
		return 1
	case 'E':
		return 2
	case 'R':
		return 3
	default:
		break
	}
	return -1
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}
