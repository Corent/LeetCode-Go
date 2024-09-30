package algorithm475

import (
	"math"
	"slices"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	rst := -1
	for _, house := range houses {
		idx, _ := slices.BinarySearch(heaters, house)
		if idx < 0 {
			idx = -(idx + 1)
		}
		distL, distR := math.MaxInt, math.MaxInt
		if idx-1 >= 0 {
			distL = house - heaters[idx-1]
		}
		if idx < len(heaters) {
			distR = heaters[idx] - house
		}
		rst = maxInt(rst, minInt(distR, distL))
	}

	return rst
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func minInt(m, n int) int {
	if m > n {
		return n
	}
	return m
}
