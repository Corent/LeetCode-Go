package algorithm373

import (
	"container/heap"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	m, n := len(nums1), len(nums2)
	mth := make(maxTopHeap, minInt(m, k))
	for i := range mth {
		mth[i] = []int{i, 0, nums1[i] + nums2[0]}
	}
	rst := make([][]int, 0, k)
	for len(rst) < k {
		peek := heap.Pop(&mth).([]int)
		i, j := peek[0], peek[1]
		rst = append(rst, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&mth, []int{i, j + 1, nums1[i] + nums2[j+1]})
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
