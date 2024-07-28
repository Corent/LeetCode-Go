package algorithm057

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) == 0 {
		return intervals
	}
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	return merge(intervals)
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	intervalList := make([][]int, 0)
	var current []int
	for _, interval := range intervals {
		if current == nil {
			current = interval
			continue
		}
		if current[1] >= interval[0] {
			if interval[1] > current[1] {
				current[1] = interval[1]
			}
		} else {
			intervalList = append(intervalList, current)
			current = interval
		}
	}
	if current != nil {
		intervalList = append(intervalList, current)
	}
	return intervalList
}
