package algorithm347

import (
	"container/heap"
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	mth := make(maxTopHeap, 0)
	for n, cnt := range m {
		heap.Push(&mth, []int{n, cnt})
	}
	fmt.Println(mth)
	rst := make([]int, k)
	for i := 0; i < k; i++ {
		rst[i] = heap.Pop(&mth).([]int)[0]
	}
	return rst
}

type maxTopHeap [][]int

func (m maxTopHeap) Len() int {
	return len(m)
}
func (m maxTopHeap) Less(i, j int) bool {
	return m[i][1] >= m[j][1]
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
