package algorithm1400

func canConstruct(s string, k int) bool {
	if len(s) == 0 || len(s) < k {
		return false
	}
	hash := [256]int{}
	for _, c := range s {
		hash[c]++
	}
	oddCnt := 0
	for _, cnt := range hash {
		oddCnt += cnt & 1
	}
	return oddCnt <= k
}
