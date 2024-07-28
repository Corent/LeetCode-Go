package algorithm061

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	length := 0
	for ptr := head; ptr != nil; ptr = ptr.Next {
		length++
	}
	if length == 1 {
		return head
	}
	header, ptr1, ptr2 := &ListNode{Next: head}, head, head
	for k %= length; k > 0; k-- {
		ptr1 = ptr1.Next
	}
	for ; ptr1.Next != nil; ptr1 = ptr1.Next {
		ptr2 = ptr2.Next
	}
	ptr1.Next = header.Next
	header.Next = ptr2.Next
	ptr2.Next = nil
	return header.Next
}
