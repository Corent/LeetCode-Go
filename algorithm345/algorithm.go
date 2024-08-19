package algorithm345

var m = map[byte]struct{}{
	'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
	'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {},
}

func reverseVowels(s string) string {
	chs := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		_, ok := m[chs[i]]
		for !ok && i < j {
			i++
			_, ok = m[chs[i]]
		}
		_, ok = m[chs[j]]
		for !ok && i < j {
			j--
			_, ok = m[chs[j]]
		}
		chs[i], chs[j] = chs[j], chs[i]
	}
	return string(chs)
}
