package algorithm472

func findAllConcatenatedWordsInADict(words []string) []string {
	hash, rst := make(map[string]struct{}), make([]string, 0)
	for _, w := range words {
		hash[w] = struct{}{}
	}
	for _, w := range words {
		n := len(w)
		dp := make([]bool, n+1)
		dp[0] = true
		for i := 0; i < n; i++ {
			if !dp[i] {
				continue
			}
			for j := i + 1; j <= n; j++ {
				if j-i < n {
					if _, ok := hash[w[i:j]]; ok {
						dp[j] = true
					}
				}
			}
			if dp[n] {
				rst = append(rst, w)
				break
			}
		}
	}
	return rst
}
