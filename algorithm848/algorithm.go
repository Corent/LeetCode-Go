package algorithm848

func shiftingLetters(s string, shifts []int) string {
	n, chs := 0, make([]byte, len(s))
	for i := len(shifts) - 1; i >= 0; i-- {
		n = (n + shifts[i]) % 26
		chs[i] = byte((n+int(s[i]-'a'))%26) + 'a'
	}
	return string(chs)
}
