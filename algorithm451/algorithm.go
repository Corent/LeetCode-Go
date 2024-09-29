package algorithm451

import (
	"sort"
	"strings"
)

func frequencySort(s string) string {
	if len(s) < 2 {
		return s
	}
	chs := [256]int{}
	for _, ch := range s {
		chs[ch]++
	}
	arr := make([][2]int, 0)
	for i, cnt := range chs {
		if cnt > 0 {
			arr = append(arr, [2]int{i, cnt})
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})
	builder := &strings.Builder{}
	for _, pair := range arr {
		for i := 0; i < pair[1]; i++ {
			builder.WriteByte(byte(pair[0]))
		}
	}
	return builder.String()
}
