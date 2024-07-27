package algorithm049

import "sort"

func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}
	m := make(map[string][]string)
	for _, s := range strs {
		key := GetKey(s)
		ss, ok := m[key]
		if !ok {
			ss = []string{s}
		} else {
			ss = append(ss, s)
		}
		m[key] = ss
	}
	ans := make([][]string, 0)
	for _, ss := range m {
		ans = append(ans, ss)
	}
	return ans
}

func GetKey(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
	return string(bytes)
}
