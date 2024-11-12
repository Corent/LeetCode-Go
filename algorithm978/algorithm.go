package algorithm978

func maxTurbulenceSize(arr []int) int {
	if len(arr) < 2 {
		return len(arr)
	}
	rst, dpOddBiggerPre, dpEvenBiggerPre := 1, 1, 1
	for i := 1; i < len(arr); i++ {
		dpOddBigger, dpEvenBigger := 1, 1
		if i&1 == 0 {
			if arr[i] > arr[i-1] {
				dpEvenBigger = dpEvenBiggerPre + 1
			} else if arr[i] < arr[i-1] {
				dpOddBigger = dpOddBiggerPre + 1
			}
		} else {
			if arr[i] > arr[i-1] {
				dpOddBigger = dpOddBiggerPre + 1
			} else if arr[i] < arr[i-1] {
				dpEvenBigger = dpEvenBiggerPre + 1
			}
		}
		dpOddBiggerPre, dpEvenBiggerPre = dpOddBigger, dpEvenBigger
		rst = maxInt(rst, maxInt(dpOddBigger, dpEvenBigger))
	}
	return rst
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}
