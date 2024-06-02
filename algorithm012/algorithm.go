package algorithm012

var roman = [][]string{
	{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
	{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
	{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
	{"", "M", "MM", "MMM"},
}

func intToRoman(num int) string {
	s := ""
	for digit := 0; num != 0; digit++ {
		remain := num % 10
		s = roman[digit][remain] + s
		num /= 10
	}
	return s
}
