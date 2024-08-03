package algorithm083

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	root, node := &ListNode{}, head
	tmp := root
	for node != nil && node.Next != nil {
		if node.Val == node.Next.Val {
			for node.Next != nil && node.Val == node.Next.Val {
				node = node.Next
			}
		} else {
			tmp.Next = node
			tmp = tmp.Next
			node = node.Next
		}
	}
	tmp.Next = node
	return root.Next
}
