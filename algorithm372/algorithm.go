package algorithm372

// https://www.bilibili.com/video/BV1k84y1q74n

const MOD = 1337

func superPow(a int, b []int) int {
	return superPowCore(a, b, len(b)-1)
}

func superPowCore(a int, b []int, idx int) int {
	if idx == -1 {
		return 1
	}
	return qpow(superPowCore(a, b, idx-1), 10) * qpow(a, b[idx]) % MOD
}

func qpow(a int, b int) int {
	rst := 1
	a %= MOD
	for b != 0 {
		if b&1 != 0 {
			rst = rst * a % MOD
		}
		a = a * a % MOD
		b >>= 1
	}
	return rst
}
