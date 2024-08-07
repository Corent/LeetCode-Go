package algorithm130

var (
	m, n  int
	Board [][]byte
)

func solve(board [][]byte) {
	m, n, Board = len(board), len(board[0]), board
	for i := 0; i < m; i++ {
		if Board[i][0] == 'O' {
			process(i, 0)
		}
		if Board[i][n-1] == 'O' {
			process(i, n-1)
		}
	}
	for j := 0; j < n; j++ {
		if Board[0][j] == 'O' {
			process(0, j)
		}
		if Board[m-1][j] == 'O' {
			process(m-1, j)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if Board[i][j] == 'O' {
				Board[i][j] = 'X'
			} else if Board[i][j] == 'E' {
				Board[i][j] = 'O'
			}
		}
	}
}

func process(i, j int) {
	queue := make([][]int, 0)
	queue, Board[i][j] = append(queue, []int{i, j}), 'E'
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		x, y := p[0], p[1]
		if x > 0 && Board[x-1][y] == 'O' {
			queue = append(queue, []int{x - 1, y})
			Board[x-1][y] = 'E'
		}
		if x < m-1 && Board[x+1][y] == 'O' {
			queue = append(queue, []int{x + 1, y})
			Board[x+1][y] = 'E'
		}
		if y > 0 && Board[x][y-1] == 'O' {
			queue = append(queue, []int{x, y - 1})
			Board[x][y-1] = 'E'
		}
		if y < n-1 && Board[x][y+1] == 'O' {
			queue = append(queue, []int{x, y + 1})
			Board[x][y+1] = 'E'
		}
	}
}
