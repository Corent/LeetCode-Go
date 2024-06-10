package algorithm017

import "strings"

var keyboard = [][]uint8{
	{},                   // 0
	{},                   // 1
	{'a', 'b', 'c'},      //2
	{'d', 'e', 'f'},      //3
	{'g', 'h', 'i'},      //4
	{'j', 'k', 'l'},      //5
	{'m', 'n', 'o'},      //6
	{'p', 'q', 'r', 's'}, //7
	{'t', 'u', 'v'},      //8
	{'w', 'x', 'y', 'z'}, //9
}

var ss string
var chs []uint8
var ans []string

func letterCombinations(digits string) []string {
	digits = strings.ReplaceAll(digits, "0", "")
	digits = strings.ReplaceAll(digits, "1", "")
	if ans = make([]string, 0); len(digits) == 0 {
		return ans
	}
	ss, chs = digits, make([]uint8, len(digits))
	letterCombinationsCore(0)
	return ans
}

func letterCombinationsCore(idx int) {
	if idx == len(ss) {
		s := ""
		for _, c := range chs {
			s += string(c)
		}
		ans = append(ans, s)
		return
	}
	for _, c := range keyboard[ss[idx]-'0'] {
		chs[idx] = c
		letterCombinationsCore(idx + 1)
	}
}
