package algorithm1024

import "sort"

func videoStitching(clips [][]int, time int) int {
	if len(clips) == 0 {
		return -1
	}
	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] != clips[j][0] {
			return clips[i][0] < clips[j][0]
		}
		return clips[i][1] < clips[j][1]
	})
	if clips[0][0] != 0 {
		return -1
	}
	n, cnt := len(clips), 0
	maxEnd := 0   //记录每次找到的最长片段的end值，下一次要找的片段需要 start <= maxEnd && maxEnd < end
	curIndex := 0 //防止重复遍历之前用过的片段
	for maxEnd < time {
		end := 0
		for i := curIndex; i < n; i++ {
			if clips[i][0] <= maxEnd && clips[i][1] > maxEnd {
				end = maxInt(end, clips[i][1])
				curIndex = i
			}
		}
		if end == 0 {
			return -1
		}
		cnt, maxEnd = cnt+1, end
	}
	return cnt
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}
