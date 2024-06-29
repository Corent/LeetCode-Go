package algorithm025

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k < 2 {
		return head
	}
	h := &ListNode{Next: head}
	ptr1, ptr2, ptr3 := h, head, head.Next
	for ptr3 != nil {
		n := k - 2
		for ; n > 0; n-- {
			if ptr3 == nil {
				break
			}
			ptr3 = ptr3.Next
		}
		if n > 0 || ptr3 == nil {
			break
		}
		tmp := ptr3.Next
		ptr3.Next = nil
		subHead, subTail := reverseList(ptr2)
		ptr1.Next = subHead
		subTail.Next = tmp
		ptr1 = subTail
		ptr2 = subTail.Next
		if ptr2 == nil {
			break
		}
		ptr3 = ptr2.Next
	}
	return h.Next
}

// returns head, tail
func reverseList(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}
	tail := head
	ptr1, ptr2 := head, head.Next
	for ptr1.Next = nil; ptr2 != nil; {
		tmp := ptr2.Next
		ptr2.Next = ptr1
		ptr1, ptr2 = ptr2, tmp
	}
	return ptr1, tail
}
