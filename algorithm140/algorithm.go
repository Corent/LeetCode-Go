package algorithm140

import "strings"

var (
	n          int
	m          [][]bool
	dp         []bool
	wordSet    map[string]struct{}
	ans, words []string
)

func wordBreak(s string, wordDict []string) []string {
	ans, words = make([]string, 0), make([]string, 0)
	if !wordBreakI(s, wordDict) {
		return ans
	}
	wordBreakCore(s, 0)
	return ans
}

func wordBreakCore(s string, from int) {
	if from == n {
		ans = append(ans, strings.Join(words, " "))
		return
	}
	if from > 0 && !dp[from-1] {
		return
	}
	for i := 0; i < n-from; i++ {
		if m[from][from+i] {
			words = append(words, s[from:from+i+1])
			wordBreakCore(s, from+i+1)
			words = words[:len(words)-1]
		}
	}
}

func wordBreakI(s string, wordDict []string) bool {
	n = len(s)
	m, wordSet = make([][]bool, n), make(map[string]struct{})
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}
	for i := 0; i < n; i++ {
		m[i] = make([]bool, n)
		for j := i; j < n; j++ {
			_, m[i][j] = wordSet[s[i:j+1]]
		}
	}

	dp = make([]bool, n)
	for j := 0; j < n; j++ {
		if j == 0 {
			dp[j] = m[0][0]
			continue
		}
		if m[0][j] {
			dp[j] = true
			continue
		}
		for i := 0; i < j; i++ {
			if dp[i] && m[i+1][j] {
				dp[j] = true
				break
			}
		}
	}
	return dp[n-1]
}
