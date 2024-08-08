package algorithm147

type ListNode struct {
	Val  int
	Next *ListNode
}

// 每次一把head位置的元素往后挪动插入
func insertionSortList(head *ListNode) *ListNode {
	root := &ListNode{}
	for head != nil {
		ptr, next := root, head.Next
		for ptr.Next != nil && ptr.Next.Val < head.Val {
			ptr = ptr.Next
		}
		head.Next = ptr.Next
		ptr.Next = head
		head = next
	}
	return root.Next
}
