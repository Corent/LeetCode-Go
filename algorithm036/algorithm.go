package algorithm036

func isValidSudoku(board [][]byte) bool {
	if len(board) != 9 || len(board[0]) != 9 {
		return false
	}
	for i := 0; i < 9; i++ {
		s1, s2, s3 := [9]int{}, [9]int{}, [9]int{}
		row, cul := (i/3)*3, (i%3)*3
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if s1[board[i][j]-'1']++; s1[board[i][j]-'1'] > 1 {
					return false
				}
			}
			if board[j][i] != '.' {
				if s2[board[j][i]-'1']++; s2[board[j][i]-'1'] > 1 {
					return false
				}
			}
			r, c := row+(j/3), cul+(j%3)
			if board[r][c] != '.' {
				if s3[board[r][c]-'1']++; s3[board[r][c]-'1'] > 1 {
					return false
				}
			}
		}
	}
	return true
}
