package algorithm435

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 || len(intervals[0]) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	var rst int
	var prev []int
	for _, interval := range intervals {
		if prev != nil && prev[1] > interval[0] {
			if rst++; prev[1] > interval[1] {
				prev = interval
			}
		} else {
			prev = interval
		}
	}
	return rst
}
