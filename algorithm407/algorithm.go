package algorithm407

import "container/heap"

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}
	m, n, ans, mth := len(heightMap), len(heightMap[0]), 0, make(minTopHeap, 0)
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
		visited[i][0], visited[i][n-1] = true, true
		heap.Push(&mth, []int{i, 0, heightMap[i][0]})
		heap.Push(&mth, []int{i, n - 1, heightMap[i][n-1]})
	}
	for i := 0; i < n; i++ {
		visited[0][i], visited[m-1][i] = true, true
		heap.Push(&mth, []int{0, i, heightMap[0][i]})
		heap.Push(&mth, []int{m - 1, i, heightMap[m-1][i]})
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(mth) > 0 {
		cell := heap.Pop(&mth).([]int)
		for _, dir := range dirs {
			row, col := cell[0]+dir[0], cell[1]+dir[1]
			if row >= 0 && row < m && col >= 0 && col < n && !visited[row][col] {
				visited[row][col] = true
				ans += maxInt(0, cell[2]-heightMap[row][col])
				heap.Push(&mth, []int{row, col, maxInt(heightMap[row][col], cell[2])})
			}
		}
	}
	return 0
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}

type minTopHeap [][]int

func (m minTopHeap) Len() int {
	return len(m)
}
func (m minTopHeap) Less(i, j int) bool {
	return m[i][2] < m[j][2]
}
func (m minTopHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minTopHeap) Push(x interface{}) {
	*m = append(*m, x.([]int))
}

func (m *minTopHeap) Pop() interface{} {
	rst := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return rst
}
