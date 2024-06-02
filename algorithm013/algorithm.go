package algorithm013

var m = map[uint8]int{
	'I': 1, 'i': 1,
	'V': 5, 'v': 5,
	'X': 10, 'x': 10,
	'L': 50, 'l': 50,
	'C': 100, 'c': 100,
	'D': 500, 'd': 500,
	'M': 1000, 'm': 1000,
}

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans := m[s[0]]
	for i := 1; i < len(s); i++ {
		pre, now := m[s[i-1]], m[s[i]]
		if pre < now {
			ans += now - pre*2 // - pre + now - pre
		} else {
			ans += now
		}
	}
	return ans
}
