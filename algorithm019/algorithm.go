package algorithm019

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	h := &ListNode{Next: head}
	p1, p2, p3 := h, head, head
	for p3 != nil {
		p3, n = p3.Next, n-1
		if n >= 0 {
			continue
		}
		p1, p2 = p1.Next, p2.Next
	}

	if n > 0 {
		return h.Next
	}

	p1.Next = p2.Next

	return h.Next
}
