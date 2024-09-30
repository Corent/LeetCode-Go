package algorithm467

func findSubstringInWraproundString(s string) int {
	sSize, count := make([]int, len(s)), [26]int{}
	for i := 0; i < len(s); i++ {
		sSize[i] = int(s[i] - 'a')
	}
	num := 0
	for i := 0; i < len(sSize); i++ {
		if i > 0 && (sSize[i-1]+1)%26 == sSize[i] {
			num++
		} else {
			num = 1
		}
		if num > count[sSize[i]] {
			count[sSize[i]] = num
		}
	}
	rst := 0
	for _, cnt := range count {
		rst += cnt
	}
	return rst
}
