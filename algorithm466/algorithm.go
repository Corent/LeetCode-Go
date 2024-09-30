package algorithm466

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	repeatCnt, nextIdx := make([]int, n1+1), make([]int, n1+1) //  重复次数，下一次匹配开始位置
	j, cnt := 0, 0
	for k := 1; k <= n1; k++ {
		for i := 0; i < len(s1); i++ {
			if s1[i] == s2[j] {
				j++
				if j == len(s2) {
					j, cnt = 0, cnt+1
				}
			}
		}
		repeatCnt[k], nextIdx[k] = cnt, j
		for start := 0; start < k; start++ { //  判断之前是否已经有从j开始的匹配，有的话可以直接计算返回
			if nextIdx[start] == j {
				interval := k - start
				repeat := (n1 - start) / interval
				patternCnt := (repeatCnt[k] - repeatCnt[start]) * repeat
				remainCnt := repeatCnt[start+(n1-start)%interval]
				return (patternCnt + remainCnt) / n2
			}
		}
	}
	return repeatCnt[n1] / n2
}
