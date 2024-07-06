package algorithm030

import "maps"

func findSubstring(s string, words []string) []int {
	if s == "" || len(words) == 0 || len(s) < len(words[0])*len(words) {
		return make([]int, 0)
	}
	ans, oMap := make([]int, 0), make(map[string]int)
	for _, word := range words {
		oMap[word]++
	}
	wMap, wLen, wsLen := make(map[string]int), len(words[0]), len(words[0])*len(words)
	maps.Copy(wMap, oMap)
	for i := 0; i < len(s)-wsLen+1; i++ {
		ok := true
		for j := i; j < i+wsLen; j += wLen {
			tmp := s[j : j+wLen]
			if cnt, exist := wMap[tmp]; exist && cnt > 0 {
				wMap[tmp] = cnt - 1
			} else {
				ok = false
				break
			}
		}
		noneLeft := true
		for _, v := range wMap {
			if noneLeft = noneLeft && (v == 0); !noneLeft {
				break
			}
		}
		if ok && noneLeft {
			ans = append(ans, i)
		}
		maps.Copy(wMap, oMap)
	}
	return ans
}
