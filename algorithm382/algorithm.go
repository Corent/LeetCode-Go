package algorithm382

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	Head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{Head: head}
}

func (this *Solution) GetRandom() int {
	node := this.Head
	point, val := 1, 0
	for node != nil {
		if rand.Intn(point) == 0 {
			val = node.Val
		}
		node, point = node.Next, point+1
	}
	return val
}
