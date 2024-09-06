package algorithm402

func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}
	stack := make([]int32, 0)
	for _, c := range num {
		for len(stack) > 0 && stack[len(stack)-1] > c && k > 0 {
			stack, k = stack[:len(stack)-1], k-1
		}
		if !(c == '0' && len(stack) == 0) {
			stack = append(stack, c)
		}
	}
	for len(stack) > 0 && k > 0 {
		stack, k = stack[:len(stack)-1], k-1
	}
	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}
