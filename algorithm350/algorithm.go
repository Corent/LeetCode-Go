package algorithm350

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int][2]int)
	for _, n := range nums1 {
		cnts, ok := m[n]
		if !ok {
			cnts = [2]int{0, 0}
		}
		cnts[0]++
		m[n] = cnts
	}
	for _, n := range nums2 {
		cnts, ok := m[n]
		if !ok {
			continue
		}
		cnts[1]++
		m[n] = cnts
	}
	rst := make([]int, 0)
	for n, cnts := range m {
		for i := 0; i < minInt(cnts[0], cnts[1]); i++ {
			rst = append(rst, n)
		}
	}
	return rst
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}
