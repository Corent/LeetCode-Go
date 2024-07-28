package algorithm067

func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	carrier, ans := uint8(0), ""
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		cha, chb := uint8('0'), uint8('0')
		if i >= 0 {
			cha = a[i]
		}
		if j >= 0 {
			chb = b[j]
		}
		sum := (cha - '0') + (chb - '0') + carrier
		carrier = sum / 2
		ans = string(uint8(sum%2)+'0') + ans
	}
	if carrier != 0 {
		ans = "1" + ans
	}
	return ans
}
