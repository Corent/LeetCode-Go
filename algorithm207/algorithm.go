package algorithm207

var (
	NumCourses   int
	nexts, prevs []map[int]struct{}
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	NumCourses, nexts, prevs = numCourses, make([]map[int]struct{}, numCourses), make([]map[int]struct{}, numCourses)
	for i := 0; i < numCourses; i++ {
		nexts[i], prevs[i] = make(map[int]struct{}), make(map[int]struct{})
	}
	for _, prerequisite := range prerequisites {
		s, e := prerequisite[0], prerequisite[1]
		nexts[s][e], prevs[e][s] = struct{}{}, struct{}{}
	}
	cnt := 0
	for zeroPrev := findNext(); zeroPrev != -1; zeroPrev = findNext() {
		for n, _ := range nexts[zeroPrev] {
			delete(prevs[n], zeroPrev)
		}
		prevs[zeroPrev] = nil
		cnt++
	}
	return cnt == numCourses
}

func findNext() int {
	for i := 0; i < NumCourses; i++ {
		if prevs[i] != nil && len(prevs[i]) == 0 {
			return i
		}
	}
	return -1
}
