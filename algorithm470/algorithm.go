package algorithm470

import "math/rand"

func rand7() int {
	return rand.Intn(7) + 1
}

var m map[int]int = map[int]int{
	2:  1,
	3:  2,
	5:  3,
	8:  4,
	10: 5,
	15: 6,
	18: 7,
	20: 8,
	24: 9,
	30: 10,
}

func rand10() int {
	v, ok := m[rand7()*rand7()]
	for !ok {
		v, ok = m[rand7()*rand7()]
	}
	return v
}
