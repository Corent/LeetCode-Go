package algorithm1345

func minJumps(arr []int) int {
	n, idx, vis := len(arr), make(map[int][]int), map[int]bool{0: true}
	for i, v := range arr {
		idx[v] = append(idx[v], i)
	}
	type pair struct{ idx, step int }
	q := []pair{{}}
	for {
		p := q[0]
		q = q[1:]
		i, step := p.idx, p.step
		if i == n-1 {
			return step
		}
		for _, j := range idx[arr[i]] {
			if !vis[j] {
				vis[j], q = true, append(q, pair{j, step + 1})
			}
		}
		delete(idx, arr[i])
		if !vis[i+1] {
			vis[i+1], q = true, append(q, pair{i + 1, step + 1})
		}
		if i > 0 && !vis[i-1] {
			vis[i-1], q = true, append(q, pair{i - 1, step + 1})
		}
	}
}
