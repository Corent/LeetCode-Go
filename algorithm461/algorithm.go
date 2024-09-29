package algorithm461

func hammingDistance(x int, y int) int {
	z, dis := x^y, 0
	for z != 0 {
		z, dis = z&(z-1), dis+1
	}
	return dis
}
