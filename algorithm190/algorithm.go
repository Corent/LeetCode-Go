package algorithm190

func reverseBits(num uint32) uint32 {
	ans := uint32(0)
	for i := 0; i < 32; i++ {
		ans |= ((num >> i) & 1) << (31 - i)
	}
	return ans
}
