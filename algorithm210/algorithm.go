package algorithm210

var (
	NumCourses   int
	nexts, prevs []map[int]struct{}
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	NumCourses, nexts, prevs = numCourses, make([]map[int]struct{}, numCourses), make([]map[int]struct{}, numCourses)
	for i := 0; i < numCourses; i++ {
		nexts[i], prevs[i] = make(map[int]struct{}), make(map[int]struct{})
	}
	for _, prerequisite := range prerequisites {
		e, s := prerequisite[0], prerequisite[1]
		nexts[s][e], prevs[e][s] = struct{}{}, struct{}{}
	}
	ans := make([]int, 0, numCourses)
	for zeroPrev := findNext(); zeroPrev != -1; zeroPrev = findNext() {
		ans = append(ans, zeroPrev)
		for n, _ := range nexts[zeroPrev] {
			delete(prevs[n], zeroPrev)
		}
		prevs[zeroPrev] = nil
	}
	if len(ans) != numCourses {
		ans = make([]int, 0)
	}
	return ans
}

func findNext() int {
	for i := 0; i < NumCourses; i++ {
		if prevs[i] != nil && len(prevs[i]) == 0 {
			return i
		}
	}
	return -1
}
