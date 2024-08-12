package algorithm273

import "strings"

var (
	units = []string{"", "Thousand", "Million", "Billion"}
	words = map[int]string{
		0:          " ",
		1:          "One",
		2:          "Two",
		3:          "Three",
		4:          "Four",
		5:          "Five",
		6:          "Six",
		7:          "Seven",
		8:          "Eight",
		9:          "Nine",
		10:         "Ten",
		11:         "Eleven",
		12:         "Twelve",
		13:         "Thirteen",
		14:         "Fourteen",
		15:         "Fifteen",
		16:         "Sixteen",
		17:         "Seventeen",
		18:         "Eighteen",
		19:         "Nineteen",
		20:         "Twenty",
		30:         "Thirty",
		40:         "Forty",
		50:         "Fifty",
		60:         "Sixty",
		70:         "Seventy",
		80:         "Eighty",
		90:         "Ninety",
		100:        "Hundred",
		1000:       "Thousand",
		1000000:    "Million",
		1000000000: "Billion",
	}
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	rst, i := "", 0
	for num > 0 {
		if num%1000 > 0 {
			rst = helper(num%1000) + units[i] + getWord(0) + rst
		}
		num /= 1000
		i++
	}
	return strings.TrimSpace(rst)
}

func helper(num int) string {
	if num >= 100 {
		return getWord(num/100) + getWord(0) + getWord(100) + getWord(0) + helper(num%100)
	}
	if num > 20 {
		return getWord((num/10)*10) + getWord(0) + helper(num%10)
	}
	if num > 0 {
		return getWord(num) + getWord(0)
	}
	if num == 0 {
		return ""
	}
	return ""
}

func getWord(num int) string {
	word, ok := words[num]
	if ok {
		return word
	}
	return " "
}
