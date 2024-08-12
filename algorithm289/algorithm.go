package algorithm289

var m, n int

func gameOfLife(board [][]int) {
	m, n = len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			liveCnt := liveNeighborCnt(board, i, j)
			if board[i][j] == 1 && liveCnt >= 2 && liveCnt <= 3 {
				board[i][j] = 3
			}
			if board[i][j] == 0 && liveCnt == 3 {
				board[i][j] = 2
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] >>= 1
		}
	}
}

func liveNeighborCnt(board [][]int, i, j int) int {
	liveCnt := 0
	xStart, xEnd := maxInt(i-1, 0), minInt(i+1, m-1)
	yStart, yEnd := maxInt(j-1, 0), minInt(j+1, n-1)
	for x := xStart; x <= xEnd; x++ {
		for y := yStart; y <= yEnd; y++ {
			liveCnt += board[x][y] & 1
		}
	}
	liveCnt -= board[i][j] & 1
	return liveCnt
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func maxInt(m, n int) int {
	if m < n {
		return n
	}
	return m
}
