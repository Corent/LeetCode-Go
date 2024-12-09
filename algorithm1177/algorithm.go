package algorithm1177

func canMakePaliQueries(s string, queries [][]int) []bool {
	n := len(s)
	cnts := make([]int, n+1)
	for i := 0; i < n; i++ {
		cnts[i+1] = cnts[i] ^ (1 << (s[i] - 'a'))
	}
	rst := make([]bool, len(queries))
	for i, querie := range queries {
		l, r, k := querie[0], querie[1], querie[2]
		bits, x := 0, cnts[r+1]^cnts[l]
		for x > 0 {
			x &= x - 1
			bits++
		}
		rst[i] = bits <= 2*k+1
	}
	return rst
}
