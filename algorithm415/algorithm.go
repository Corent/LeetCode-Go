package algorithm415

func addStrings(num1 string, num2 string) string {
	if num1 == "0" {
		return num2
	}
	if num2 == "0" {
		return num1
	}
	rst, carry := "", uint8(0)
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; {
		n1, n2 := uint8(0), uint8(0)
		if i >= 0 {
			n1, i = num1[i]-'0', i-1
		}
		if j >= 0 {
			n2, j = num2[j]-'0', j-1
		}
		sum := n1 + n2 + carry
		carry, rst = sum/10, string(sum%10+'0')+rst
	}
	if carry > 0 {
		rst = string(carry+'0') + rst
	}
	return rst
}
