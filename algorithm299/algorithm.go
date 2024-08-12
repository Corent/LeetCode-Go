package algorithm299

import "fmt"

func getHint(secret string, guess string) string {
	aCnt, bCnt, n := 0, 0, len(secret)
	ds, dg := [10]int{}, [10]int{}
	for i := 0; i < n; i++ {
		x, y := secret[i]-'0', guess[i]-'0'
		if x == y {
			aCnt++
		}
		ds[x], dg[y] = ds[x]+1, dg[y]+1
	}
	for i := 0; i < 10; i++ {
		bCnt += minInt(ds[i], dg[i])
	}
	return fmt.Sprintf("%dA%dB", aCnt, bCnt-aCnt)
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}
