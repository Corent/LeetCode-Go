package algorithm056

import "sort"

type Segments [][]int

func (s Segments) Len() int {
	return len(s)
}
func (s Segments) Less(i, j int) bool {
	if s[i][0] != s[j][0] {
		return s[i][0] < s[j][0]
	}
	return s[i][1] < s[j][1]
}
func (s Segments) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	intervalList := make([][]int, 0)
	sort.Sort(Segments(intervals))
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
