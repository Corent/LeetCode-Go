package algorithm433

func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	bankSet, charSet := make(map[string]struct{}), []byte{'A', 'C', 'G', 'T'}
	for _, b := range bank {
		bankSet[b] = struct{}{}
	}
	level, visited, queue := 0, map[string]struct{}{}, make([]string, 0)
	queue, visited[startGene] = append(queue, startGene), struct{}{}
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			size--
			curr := queue[0]
			queue = queue[1:]
			if curr == endGene {
				return level
			}
			currArray := []byte(curr)
			for i := 0; i < len(currArray); i++ {
				old := currArray[i]
				for _, c := range charSet {
					if old == c {
						continue
					}
					currArray[i] = c
					next := string(currArray)
					_, ok1 := visited[next]
					_, ok2 := bankSet[next]
					if !ok1 && ok2 {
						queue, visited[next] = append(queue, next), struct{}{}
					}
				}
				currArray[i] = old
			}
		}
		level++
	}
	return -1
}
