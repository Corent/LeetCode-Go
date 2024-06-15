package algorithm020

var m = map[uint8]uint8{
	')': '(', ']': '[', '}': '{',
}

func isValid(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}
	stack := make([]uint8, 0)
	for i := 0; i < n; i++ {
		v, ok := m[s[i]]
		if !ok || len(stack) == 0 || v != stack[len(stack)-1] {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) > 0 && v == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
