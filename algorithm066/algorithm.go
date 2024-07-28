package algorithm066

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}
	carrier := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carrier
		carrier = digits[i] / 10
		digits[i] %= 10
	}
	if carrier > 0 {
		digits = append([]int{carrier}, digits...)
	}
	return digits
}
