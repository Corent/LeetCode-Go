package algorithm079

var (
	m, n int
	w    string
)

func exist(board [][]byte, word string) bool {
	m, n, w = len(board), len(board[0]), word
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == w[0] && find(board, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func find(board [][]byte, i, j, idx int) bool {
	if idx == len(w) {
		return true
	}
	if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != w[idx] {
		return false
	}
	board[i][j] = '#'
	found := find(board, i-1, j, idx+1) ||
		find(board, i, j-1, idx+1) ||
		find(board, i+1, j, idx+1) ||
		find(board, i, j+1, idx+1)
	board[i][j] = w[idx]
	return found
}
