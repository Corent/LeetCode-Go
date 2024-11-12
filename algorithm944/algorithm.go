package algorithm944

func minDeletionSize(strs []string) int {
	m, n, rst := len(strs), len(strs[0]), 0
	for j := 0; j < n; j++ {
		pre := strs[0][j]
		for i := 1; i < m; i++ {
			if ch := strs[i][j]; ch >= pre {
				pre = ch
				continue
			}
			rst++
			break
		}
	}
	return rst
}
