package algorithm692

import "container/heap"

func topKFrequent(words []string, k int) []string {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	h := &minTopHeap{}
	for w, c := range cnt {
		heap.Push(h, pair{w, c})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).(pair).w
	}
	return ans
}

type pair struct {
	w string
	c int
}
type minTopHeap []pair

func (h minTopHeap) Len() int { return len(h) }
func (h minTopHeap) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.c < b.c || a.c == b.c && a.w > b.w
}
func (h minTopHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minTopHeap) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *minTopHeap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
