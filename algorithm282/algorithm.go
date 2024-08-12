package algorithm282

import (
	"fmt"
	"strconv"
)

var (
	Target int
	ans    []string
)

func addOperators(num string, target int) []string {
	Target, ans = target, make([]string, 0)
	addOperatorsCore(num, "", 0, 0)
	return ans
}

func addOperatorsCore(num, tmp string, rst, prevNum int) {
	if rst == Target && len(num) == 0 {
		ans = append(ans, tmp)
		return
	}
	for i := 1; i <= len(num); i++ {
		curStr := num[:i]
		if len(curStr) > 1 && curStr[0] == '0' {
			return
		}
		curNum, _ := strconv.Atoi(curStr)
		next := num[i:]
		if len(tmp) != 0 {
			addOperatorsCore(next, fmt.Sprintf("%s*%d", tmp, curNum), (rst-prevNum)+prevNum*curNum, prevNum*curNum)
			addOperatorsCore(next, fmt.Sprintf("%s+%d", tmp, curNum), rst+curNum, curNum)
			addOperatorsCore(next, fmt.Sprintf("%s-%d", tmp, curNum), rst-curNum, -curNum)
		} else {
			addOperatorsCore(next, curStr, curNum, curNum)
		}
	}
}
