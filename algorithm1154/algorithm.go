package algorithm1154

import "strconv"

func dayOfYear(date string) int {
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	year, _ := strconv.Atoi(date[0:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])
	if year%1000 == 0 || year%100 != 0 && year%4 == 0 {
		days[1]++
	}
	rst := 0
	for month -= 2; month >= 0; month-- {
		rst += days[month]
	}
	return rst + day
}
