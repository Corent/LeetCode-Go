package algorithm206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev, ptr *ListNode
	for ptr = head; ptr != nil; ptr = ptr.Next {
		next := ptr.Next
		ptr.Next = prev
		prev = ptr
		ptr = next
	}
	return prev
}
