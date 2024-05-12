package algorithm006

import "strings"

func convert(s string, numRows int) string {
	if len(s) <= 1 || numRows <= 1 {
		return s
	}
	numCols := len(s) / (numRows*2 - 2)
	if len(s)%(numRows*2-2) != 0 {
		numCols++
	}
	m := make([][]uint8, numRows)
	for i := 0; i < numRows; i++ {
		m[i] = make([]uint8, numCols)
	}
	x, y, xState, yState := 0, 0, 1, 0
	for k := 0; k < len(s); k++ {
		m[x][y] = s[k]
		if x == numRows-1 && xState == 1 {
			xState, yState = -1, 1
		} else if x == 0 && xState == -1 {
			xState, yState = 1, 0
		}
		x, y = x+xState, y+yState
	}
	builder := strings.Builder{}
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if m[i][j] == 0 {
				continue
			}
			_ = builder.WriteByte(m[i][j])
		}
	}
	return builder.String()
}
