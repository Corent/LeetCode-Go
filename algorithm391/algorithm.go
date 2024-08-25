package algorithm391

import (
	"fmt"
	"math"
)

func isRectangleCover(rectangles [][]int) bool {
	left, right, top, bottom := math.MaxInt, math.MinInt, math.MinInt, math.MaxInt
	totalArea, flags := 0, make(map[string]struct{})
	for _, rectangle := range rectangles {
		left, bottom, right, top = minInt(left, rectangle[0]), minInt(bottom, rectangle[1]), maxInt(right, rectangle[2]), maxInt(top, rectangle[3])
		totalArea += (rectangle[3] - rectangle[1]) * (rectangle[2] - rectangle[0])
		pointLT := fmt.Sprintf("%d %d", rectangle[0], rectangle[3])
		pointLB := fmt.Sprintf("%d %d", rectangle[0], rectangle[1])
		pointRT := fmt.Sprintf("%d %d", rectangle[2], rectangle[3])
		pointRB := fmt.Sprintf("%d %d", rectangle[2], rectangle[1])
		putOnce(flags, pointLT)
		putOnce(flags, pointLB)
		putOnce(flags, pointRT)
		putOnce(flags, pointRB)
	}
	pointLT := fmt.Sprintf("%d %d", left, top)
	pointLB := fmt.Sprintf("%d %d", left, bottom)
	pointRT := fmt.Sprintf("%d %d", right, top)
	pointRB := fmt.Sprintf("%d %d", right, bottom)
	if len(flags) == 4 && mapContains(flags, pointLT) && mapContains(flags, pointLB) && mapContains(flags, pointRT) && mapContains(flags, pointRB) {
		return totalArea == (right-left)*(top-bottom)
	}
	return false
}

func putOnce(m map[string]struct{}, k string) {
	if _, ok := m[k]; ok {
		delete(m, k)
	} else {
		m[k] = struct{}{}
	}
}

func mapContains(m map[string]struct{}, k string) bool {
	_, ok := m[k]
	return ok
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func maxInt(m, n int) int {
	if m < n {
		return n
	}
	return m
}
