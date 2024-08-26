package algorithm401

import "strconv"

func readBinaryWatch(turnedOn int) []string {
	hours, minutes := make(map[int][]string), make(map[int][]string)
	for i := 0; i < 12; i++ {
		ones := countOne(i)
		hours[ones] = append(hours[ones], strconv.Itoa(i))
	}
	for i := 0; i <= 59; i++ {
		ones := countOne(i)
		v := strconv.Itoa(i)
		if i < 10 {
			v = "0" + v
		}
		minutes[ones] = append(minutes[ones], v)
	}
	rst := make([]string, 0)
	for i := 0; i <= turnedOn; i++ {
		hs, ms := hours[i], minutes[turnedOn-i]
		if len(hs) == 0 || len(ms) == 0 {
			continue
		}
		for _, h := range hs {
			for _, m := range ms {
				rst = append(rst, h+":"+m)
			}
		}
	}
	return rst
}

func countOne(n int) int {
	cnt := 0
	for n != 0 {
		n, cnt = n&(n-1), cnt+1
	}
	return cnt
}
