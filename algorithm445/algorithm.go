package algorithm445

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1, l2 = reverse(l1), reverse(l2)
	root := &ListNode{Val: 0}
	ptr, carry := root, 0
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		v := v1 + v2 + carry
		ptr.Next, carry = &ListNode{Val: v % 10}, v/10
		ptr = ptr.Next
	}
	if carry > 0 {
		ptr.Next = &ListNode{Val: carry}
	}
	return reverse(root.Next)
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev, ptr *ListNode
	for ptr = head; ptr != nil; {
		next := ptr.Next
		ptr.Next = prev
		prev = ptr
		ptr = next
	}
	return prev
}
