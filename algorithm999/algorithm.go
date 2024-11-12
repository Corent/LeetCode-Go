package algorithm999

func numRookCaptures(board [][]byte) int {
	m, n, rst := len(board), len(board[0]), 0
	ri, rj, idx := -1, -1, -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'R' {
				ri, rj = i, j
				break
			}
		}
	}
	if ri == -1 || rj == -1 {
		return rst
	}
	for idx = ri - 1; idx >= 0 && board[idx][rj] == '.'; {
		idx--
	}
	if idx >= 0 && board[idx][rj] == 'p' {
		rst++
	}

	for idx = ri + 1; idx < m && board[idx][rj] == '.'; {
		idx++
	}
	if idx < m && board[idx][rj] == 'p' {
		rst++
	}

	for idx = rj - 1; idx >= 0 && board[ri][idx] == '.'; {
		idx--
	}
	if idx >= 0 && board[ri][idx] == 'p' {
		rst++
	}

	for idx = rj + 1; idx < n && board[ri][idx] == '.'; {
		idx++
	}
	if idx < n && board[ri][idx] == 'p' {
		rst++
	}
	return rst
}
