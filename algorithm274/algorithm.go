package algorithm274

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	for i, n := len(citations)-1, len(citations); i >= 0; i-- {
		if citations[i] <= n-i-1 {
			return n - i - 1
		}
	}
	return len(citations)
}
