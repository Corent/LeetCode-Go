package algorithm002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{}
	ptr, carry := head, 0
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}
		sum := carry + v1 + v2
		carry, ptr.Next = sum/10, &ListNode{Val: sum % 10}
		ptr = ptr.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		ptr.Next = &ListNode{Val: carry}
	}
	return head.Next
}
