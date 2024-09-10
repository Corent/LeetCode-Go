package algorithm406

import (
	"slices"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] != people[j][0] {
			return people[j][0] < people[i][0]
		}
		return people[i][1] < people[j][1]
	})
	rst := make([][]int, 0)
	for _, p := range people {
		left, right := slices.Clone[[][]int](rst[:p[1]]), slices.Clone[[][]int](rst[p[1]:])
		rst = append(append(left, p), right...)
	}
	return rst
}
