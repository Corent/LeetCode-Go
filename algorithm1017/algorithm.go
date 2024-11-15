package algorithm1017

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	bytes := make([]byte, 0)
	for n != 0 {
		m := n % -2
		n /= -2
		if m < 0 {
			m, n = m+2, n+1
		}
		bytes = append(bytes, byte(m)+'0')
	}
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}
