package algorithm378

import "container/heap"

func kthSmallest(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	mth := make(maxTopHeap, minInt(m, k))
	for i := range mth {
		mth[i] = []int{i, 0, matrix[i][0]}
	}
	var rst int
	for ; k > 0; k-- {
		peek := heap.Pop(&mth).([]int)
		i, j := peek[0], peek[1]
		rst = matrix[i][j]
		if j+1 < n {
			heap.Push(&mth, []int{i, j + 1, matrix[i][j+1]})
		}
	}
	return rst
}

type maxTopHeap [][]int

func (m maxTopHeap) Len() int {
	return len(m)
}
func (m maxTopHeap) Less(i, j int) bool {
	return m[i][2] < m[j][2]
}
func (m maxTopHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxTopHeap) Push(x interface{}) {
	*m = append(*m, x.([]int))
}

func (m *maxTopHeap) Pop() interface{} {
	rst := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return rst
}

func minInt(m, n int) int {
	if m > n {
		return n
	}
	return m
}

var (
	n, K int
	m    [][]int
)

func kthSmallest2(matrix [][]int, k int) int {
	m, n, K = matrix, len(matrix), k
	rst := 0
	for L, R := matrix[0][0], matrix[n-1][n-1]; L <= R; {
		mid := (L + R) >> 1
		if guess(mid) {
			rst, L = mid, mid+1
		} else {
			R = mid - 1
		}
	}
	return rst
}

func guess(g int) bool {
	sum := 0
	for i := 0; i < n; i++ {
		cnt := 0
		for L, R := 0, n-1; L <= R; {
			mid := (L + R) >> 1
			if m[i][mid] < g {
				L, cnt = mid+1, mid+1
			} else {
				R = mid - 1
			}
		}
		sum += cnt
	}
	return K > sum
}
