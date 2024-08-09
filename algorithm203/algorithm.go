package algorithm203

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	root := &ListNode{Next: head}
	for ptr := root; ptr.Next != nil; {
		if ptr.Next.Val == val {
			ptr.Next = ptr.Next.Next
		} else {
			ptr = ptr.Next
		}
	}
	return root.Next
}
