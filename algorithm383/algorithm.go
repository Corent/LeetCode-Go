package algorithm383

func canConstruct(ransomNote string, magazine string) bool {
	hash := [26]int{}
	for _, r := range magazine {
		hash[r-'a']++
	}
	for _, r := range ransomNote {
		hash[r-'a']--
		if hash[r-'a'] < 0 {
			return false
		}
	}
	return true
}
