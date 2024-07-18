package algorithm038

import "fmt"

func countAndSay(n int) string {
	if n < 1 {
		return ""
	}
	ans := "1"
	for n > 1 {
		ch, cnt, next := ans[0], 1, ""
		for i := 1; i < len(ans); i++ {
			if ch == ans[i] {
				cnt++
			} else {
				next += fmt.Sprintf("%d%s", cnt, string(ch))
				ch, cnt = ans[i], 1
			}
		}
		next += fmt.Sprintf("%d%s", cnt, string(ch))
		ans = next
		n--
	}
	return ans
}
