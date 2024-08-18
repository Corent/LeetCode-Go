package algorithm332

import (
	"slices"
	"sort"
)

var (
	nodes, path, ans []string
	matrix           [][]int
	Tickets          [][]string
	nameToId         map[string]int
	idToName         map[int]string
)

func findItinerary(tickets [][]string) []string {
	Tickets, path, ans = tickets, make([]string, 0), nil
	nodeSet := make(map[string]struct{})
	for _, ticket := range tickets {
		nodeSet[ticket[0]], nodeSet[ticket[1]] = struct{}{}, struct{}{}
	}
	nodes = make([]string, 0)
	for k, _ := range nodeSet {
		nodes = append(nodes, k)
	}
	sort.Strings(nodes)
	nameToId, idToName = make(map[string]int), make(map[int]string)
	for id, name := range nodes {
		nameToId[name], idToName[id] = id, name
	}
	n := len(nodes)
	matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for _, ticket := range tickets {
		x, y := nameToId[ticket[0]], nameToId[ticket[1]]
		matrix[x][y]++
	}

	dfs(nameToId["JFK"])

	return ans
}

func dfs(idx int) {
	if ans != nil {
		return
	}
	path = append(path, idToName[idx])
	if len(path) == len(Tickets)+1 {
		ans = slices.Clone[[]string](path)
		return
	}
	for next := 0; next < len(nodes); next++ {
		if matrix[idx][next] > 0 {
			matrix[idx][next]--
			dfs(next)
			matrix[idx][next]++
		}
	}
	path = path[:len(path)-1]
}
