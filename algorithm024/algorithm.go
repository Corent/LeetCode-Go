package algorithm024

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h := &ListNode{Next: head}
	ptr1, ptr2, ptr3 := h, head, head.Next
	for ptr3 != nil {
		ptr1.Next = ptr3
		ptr2.Next = ptr3.Next
		ptr3.Next = ptr2
		ptr2, ptr3 = ptr3, ptr2
		ptr3, ptr2, ptr1 = ptr3.Next, ptr2.Next, ptr1.Next
		if ptr3 == nil {
			break
		}
		ptr3, ptr2, ptr1 = ptr3.Next, ptr2.Next, ptr1.Next
	}
	return h.Next
}
