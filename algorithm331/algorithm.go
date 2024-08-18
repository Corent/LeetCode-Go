package algorithm331

func isValidSerialization(preorder string) bool {
	if len(preorder) == 0 {
		return false
	}
	n, stack := len(preorder), make([]uint8, 0)
	for i := 0; i < n; i++ {
		if preorder[i] == ',' {
			continue
		}
		for preorder[i] == '#' && len(stack) > 0 && stack[len(stack)-1] == '#' {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		for preorder[i] >= '0' && preorder[i] <= '9' && i+1 < n && preorder[i+1] >= '0' && preorder[i+1] <= '9' {
			i++
		}
		stack = append(stack, preorder[i])
	}
	return len(stack) == 1 && stack[0] == '#'
}
