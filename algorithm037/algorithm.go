package algorithm037

func solveSudoku(board [][]byte) {
	solveSudokuHelper(0, board)
}

func solveSudokuHelper(idx int, board [][]byte) bool {
	if idx == 81 {
		return true
	}
	i, j := idx/9, idx%9
	if board[i][j] != '.' {
		return solveSudokuHelper(idx+1, board)
	}
	candidates := findCandidates(idx, board)
	for _, candidate := range candidates {
		board[i][j] = candidate
		if solveSudokuHelper(idx+1, board) {
			return true
		}
	}
	board[i][j] = '.'
	return false
}

func findCandidates(idx int, board [][]byte) []byte {
	i, j := idx/9, idx%9
	if board[i][j] != '.' {
		return []byte{}
	}
	s, row, col := [9]int{}, (i/3)*3, (j/3)*3
	for k := 0; k < 9; k++ {
		if board[i][k] != '.' {
			s[board[i][k]-'1']++
		}
		if board[k][j] != '.' {
			s[board[k][j]-'1']++
		}
		r, c := row+k/3, col+k%3
		if board[r][c] != '.' {
			s[board[r][c]-'1']++
		}
	}
	candidates := make([]byte, 0)
	for k := 0; k < 9; k++ {
		if s[k] == 0 {
			candidates = append(candidates, byte(k)+'1')
		}
	}
	return candidates
}
