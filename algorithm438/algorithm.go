package algorithm438

func findAnagrams(s string, p string) []int {
	hash, m, n := [26]int{}, len(s), len(p)
	for i := 0; i < n; i++ {
		hash[p[i]-'a']++
	}
	rst := make([]int, 0)
	for i, j := 0, 0; j < m; {
		hash[s[j]-'a']--
		for i <= j && hash[s[j]-'a'] < 0 {
			hash[s[i]-'a']++
			i++
		}
		if j-i+1 == n {
			rst = append(rst, i)
		}
		j++
	}
	return rst
}
