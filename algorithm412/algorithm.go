package algorithm412

import "strconv"

func fizzBuzz(n int) []string {
	rst := make([]string, n)
	for i := 0; i < n; i++ {
		num, s := i+1, ""
		if num%3 == 0 {
			s = "Fizz"
		}
		if num%5 == 0 {
			s += "Buzz"
		}
		if len(s) == 0 {
			s = strconv.Itoa(num)
		}
		rst[i] = s
	}
	return rst
}
