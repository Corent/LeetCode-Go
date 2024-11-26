package algorithm1111

func maxDepthAfterSplit(seq string) []int {
	rst := make([]int, len(seq))
	for i, c := range seq {
		if c == '(' {
			rst[i] = i & 1
		} else {
			rst[i] = 1 - i&1
		}
	}
	return rst
}
