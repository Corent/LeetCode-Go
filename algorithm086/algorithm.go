package algorithm086

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	less, more := &ListNode{}, &ListNode{}
	ptrLess, ptrMore := less, more
	for ptr := head; ptr != nil; {
		if ptr.Val < x {
			ptrLess.Next = ptr
			ptrLess = ptrLess.Next
			ptr = ptr.Next
			ptrLess.Next = nil
		} else {
			ptrMore.Next = ptr
			ptrMore = ptrMore.Next
			ptr = ptr.Next
			ptrMore.Next = nil
		}
	}
	ptrLess.Next = more.Next
	return less.Next
}
