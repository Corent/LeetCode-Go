package algorithm043

func multiply(num1 string, num2 string) string {
	if len(num1) < 1 || num1 == "0" || len(num2) < 1 || num2 == "0" {
		return "0"
	}
	len1, len2 := len(num1), len(num2)
	rst := make([]int, len1+len2)
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			rst[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
		}
	}
	var ans string
	for k := len1 + len2 - 1; k >= 0; k-- {
		if k > 0 {
			rst[k-1] += rst[k] / 10
		}
		rst[k] %= 10
		ans = string(byte(rst[k]+'0')) + ans
	}
	if ans[0] == '0' {
		ans = ans[1:]
	}
	return ans
}
