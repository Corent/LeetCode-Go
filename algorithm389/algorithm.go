package algorithm389

func findTheDifference(s string, t string) byte {
	chs := [26]int{}
	for i, c := range s {
		chs[c-'a']++
		chs[t[i]-'a']--
	}
	chs[t[len(t)-1]-'a']--
	for i, c := range chs {
		if c < 0 {
			return byte(i + 'a')
		}
	}
	return 0
}
