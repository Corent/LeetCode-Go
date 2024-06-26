package algorithm009

func isPalindrome(x int) bool {
	if x < 0 || x > 0 && x%10 == 0 {
		return false
	}
	if x < 10 {
		return true
	}
	reverse := 0
	for x > reverse {
		reverse = reverse*10 + x%10
		x = x / 10
	}
	return x == reverse || x == reverse/10
}
