package algorithm386

func lexicalOrder(n int) []int {
	rst := []int{1}
	for i, num := 1, 1; i < n; i++ {
		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num == n {
				num /= 10
			}
			num++
		}
		rst = append(rst, num)
	}
	return rst
}
