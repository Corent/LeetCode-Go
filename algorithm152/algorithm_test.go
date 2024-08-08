package algorithm152

import "testing"

func TestMaxProduct(t *testing.T) {
	nums := []int{0, 10, 10, 10, 10, 10, 10, 10, 10, 10, -10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 0}
	t.Log(maxProduct(nums))
}
