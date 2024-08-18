package algorithm316

import "strings"

func removeDuplicateLetters(s string) string {
	chCnt, visited, queue := [26]int{}, [26]bool{}, make([]int32, 0)
	for _, ch := range s {
		chCnt[ch-'a']++
	}
	for _, ch := range s {
		idx := ch - 'a'
		chCnt[idx]--
		if visited[idx] {
			continue
		}
		for len(queue) > 0 && chCnt[queue[len(queue)-1]-'a'] > 0 && ch < queue[len(queue)-1] {
			c := queue[len(queue)-1]
			queue, visited[c-'a'] = queue[:len(queue)-1], false
		}
		queue, visited[idx] = append(queue, ch), true
	}
	builder := strings.Builder{}
	for _, ch := range queue {
		builder.WriteRune(ch)
	}
	return builder.String()
}
