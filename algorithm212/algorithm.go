package algorithm212

var (
	m, n  int
	Board [][]byte
	ans   []string
)

func findWords(board [][]byte, words []string) []string {
	m, n, Board, ans = len(board), len(board[0]), board, make([]string, 0)
	starts := make(map[byte]map[string]struct{})
	for _, word := range words {
		wordSet, ok := starts[word[0]]
		if !ok || wordSet == nil {
			wordSet = make(map[string]struct{})
		}
		wordSet[word] = struct{}{}
		starts[word[0]] = wordSet
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			wordSet, ok := starts[board[i][j]]
			if !ok || wordSet == nil {
				continue
			}
			toDeleted := make([]string, 0)
			for word, _ := range wordSet {
				if find(i, j, 0, word) {
					ans = append(ans, word)
					toDeleted = append(toDeleted, word)
				}
			}
			for _, word := range toDeleted {
				delete(wordSet, word)
			}
			starts[board[i][j]] = wordSet
		}
	}
	return ans
}

func find(i, j, from int, word string) bool {
	if Board[i][j] != word[from] {
		return false
	}
	if from == len(word)-1 {
		return true
	}
	Board[i][j] = '#'
	found := false
	if i > 0 && Board[i-1][j] == word[from+1] {
		found = find(i-1, j, from+1, word)
	}
	if !found && i < m-1 && Board[i+1][j] == word[from+1] {
		found = find(i+1, j, from+1, word)
	}
	if !found && j > 0 && Board[i][j-1] == word[from+1] {
		found = find(i, j-1, from+1, word)
	}
	if !found && j < n-1 && Board[i][j+1] == word[from+1] {
		found = find(i, j+1, from+1, word)
	}
	Board[i][j] = word[from]
	return found
}
