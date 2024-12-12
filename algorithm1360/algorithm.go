package algorithm1360

import (
	"strconv"
	"strings"
)

func daysBetweenDates(date1 string, date2 string) int {
	rst := dateToInt(date1) - dateToInt(date2)
	if rst < 0 {
		rst = -rst
	}
	return rst
}

func dateToInt(date string) int {
	ss := strings.Split(date, "-")
	if len(ss) != 3 {
		return 0
	}
	year, _ := strconv.Atoi(ss[0])
	month, _ := strconv.Atoi(ss[1])
	day, _ := strconv.Atoi(ss[2])
	rst, months := day-1, []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for month != 0 {
		month--
		rst += months[month]
		if month == 2 && isLeapYear(year) {
			rst += 1
		}
	}
	rst += 365 * (year - 1971)
	rst += (year-1)/4 - 1971/4
	rst -= (year-1)/100 - 1971/100
	rst += (year-1)/400 - 1971/400
	return rst
}

func isLeapYear(year int) bool {
	return year%400 == 0 || year%100 != 0 && year%4 == 0
}
