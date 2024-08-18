package algorithm328

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	oddRoot, evenRoot := &ListNode{}, &ListNode{}
	oddPtr, evenPtr := oddRoot, evenRoot
	for idx, ptr := 1, head; ptr != nil; idx = idx + 1 {
		if idx&1 == 1 {
			oddPtr.Next = ptr
			oddPtr = oddPtr.Next
			ptr = ptr.Next
			oddPtr.Next = nil
		} else {
			evenPtr.Next = ptr
			evenPtr = evenPtr.Next
			ptr = ptr.Next
			evenPtr.Next = nil
		}
	}
	oddPtr.Next = evenRoot.Next
	return oddRoot.Next
}
