package algorithm464

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger >= desiredTotal {
		return true
	}
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	return canWin(maxChoosableInteger, desiredTotal, 0, make(map[int]bool))
}

func canWin(length, total, used int, m map[int]bool) bool {
	rst, ok := m[used]
	if ok {
		return rst
	}
	for i := 0; i < length; i++ {
		cur := 1 << i
		if cur&used == 0 {
			if total <= i+1 || !canWin(length, total-(i+1), cur|used, m) {
				m[used] = true
				return true
			}
		}
	}
	m[used] = false
	return false
}
