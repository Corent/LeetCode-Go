package algorithm127

func ladderLength(beginWord string, endWord string, wordList []string) int {
	n, ans, queue, wordSet := len(beginWord), 1, make([]string, 0), make(map[string]struct{})
	for _, word := range wordList {
		wordSet[word] = struct{}{}
	}
	if _, ok := wordSet[endWord]; !ok {
		return 0
	}
	delete(wordSet, beginWord)
	queue = append(queue, beginWord)
	for len(queue) > 0 {
		size := len(queue)
		for cnt := 0; cnt < size; cnt++ {
			word := queue[0]
			queue = queue[1:]
			chs := []rune(word)
			for j := 0; j < n; j++ {
				origin := chs[j]
				for c := 'a'; c <= 'z'; c++ {
					if c == origin {
						continue
					}
					chs[j] = c
					newWord := string(chs)
					if newWord == endWord {
						return ans + 1
					}
					if _, ok := wordSet[newWord]; ok {
						queue = append(queue, newWord)
						delete(wordSet, newWord)
					}
				}
				chs[j] = origin
			}
		}
		ans++
	}
	return 0
}
