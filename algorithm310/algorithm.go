package algorithm310

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	m := make(map[int]map[int]struct{})
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		s, ok := m[x]
		if !ok || s == nil {
			s = make(map[int]struct{})
		}
		s[y] = struct{}{}
		m[x] = s
		s, ok = m[y]
		if !ok || s == nil {
			s = make(map[int]struct{})
		}
		s[x] = struct{}{}
		m[y] = s
	}

	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if len(m[i]) == 1 {
			queue = append(queue, i)
		}
	}

	for n > 2 {
		size := len(queue)
		n -= size
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			next := GetElementInSet(m[node])
			delete(m[next], node)
			delete(m, node)
			if len(m[next]) == 1 {
				queue = append(queue, next)
			}
		}
	}
	return queue
}

func GetElementInSet(m map[int]struct{}) int {
	for k := range m {
		return k
	}
	return -1
}
