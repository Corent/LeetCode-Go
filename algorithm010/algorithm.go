package algorithm010

func isMatch(s string, p string) bool {
	return isMatchCore(s, 0, p, 0)
}

func isMatchCore(s string, idx1 int, p string, idx2 int) bool {
	if idx2 == len(p) {
		return idx1 == len(s)
	}
	if idx1 == len(s) {
		if idx2+1 < len(p) && p[idx2+1] == '*' {
			return isMatchCore(s, idx1, p, idx2+2)
		}
		return idx2 == len(p)
	}
	if idx2+1 == len(p) || p[idx2+1] != '*' { // p[idx2 + 1]不是*的情况
		return (s[idx1] == p[idx2] || p[idx2] == '.') && isMatchCore(s, idx1+1, p, idx2+1)
	}
	// p[idx2 + 1] = *
	if p[idx2] == '.' {
		for i := idx1; i <= len(s); i++ {
			if isMatchCore(s, i, p, idx2+2) {
				return true
			}
		}
		return false
	}

	for i, c := idx1, p[idx2]; i <= len(s) && (i == idx1 || s[i-1] == c); i++ {
		if isMatchCore(s, i, p, idx2+2) {
			return true
		}
	}
	return false
}

func isMatchCore2(s string, idx1 int, p string, idx2 int) bool {
	if idx2 == len(p) {
		return idx1 == len(s)
	}
	// p[idx2 + 1]不是*的情况
	if idx2+1 == len(p) || p[idx2+1] != '*' {
		if idx1 == len(s) || s[idx1] != p[idx2] && p[idx2] != '.' {
			return false
		}
		return isMatchCore2(s, idx1+1, p, idx2+1)
	}
	// p[idx2 + 1] = *
	for ; idx1 < len(s) && (s[idx1] == p[idx2] || p[idx2] == '.'); idx1++ {
		if isMatchCore2(s, idx1, p, idx2+2) {
			return true
		}
	}
	return isMatchCore2(s, idx1, p, idx2+2) // .*或a*匹配0个字符
}
