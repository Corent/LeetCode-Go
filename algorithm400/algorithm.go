package algorithm400

import "strconv"

func findNthDigit(n int) int {
	tenPowN, digits, sum := 1, 1, 9
	for sum < n {
		digits++
		tenPowN *= 10
		sum += tenPowN * 9 * digits
	}
	sum -= tenPowN * 9 * digits
	n -= sum + 1
	s := strconv.Itoa(n/digits + tenPowN)
	return int(s[n%digits] - '0')
}
