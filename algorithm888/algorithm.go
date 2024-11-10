package algorithm888

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	isSwaped := false
	if len(aliceSizes) > len(bobSizes) {
		aliceSizes, bobSizes, isSwaped = bobSizes, aliceSizes, true
	}
	sumA, sumB, set := 0, 0, make(map[int]struct{})
	for _, size := range aliceSizes {
		sumA += size
	}
	for _, size := range bobSizes {
		sumB, set[size] = sumB+size, struct{}{}
	}
	if sumA == sumB {
		return []int{0, 0}
	}
	diff := (sumB - sumA) >> 1
	for _, size := range aliceSizes {
		if _, ok := set[size+diff]; ok {
			if isSwaped {
				return []int{size + diff, size}
			}
			return []int{size, size + diff}
		}
	}
	return []int{0, 0}
}
