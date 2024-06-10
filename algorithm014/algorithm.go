package algorithm014

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	longestCommonPrefixCore(strs, 0, len(strs)-1)
	return strs[0]
}

func longestCommonPrefixCore(strs []string, from, to int) {
	if from == to {
		return
	}

	mid := (from + to) / 2
	longestCommonPrefixCore(strs, from, mid)
	longestCommonPrefixCore(strs, mid+1, to)

	s := ""
	for i := 0; i < len(strs[from]) && i < len(strs[mid+1]); i++ {
		if strs[from][i] != strs[mid+1][i] {
			break
		}
		s += string(strs[from][i])
	}
	strs[from] = s
}
