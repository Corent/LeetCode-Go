package algorithm021

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	header := &ListNode{}
	ptr := header
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			ptr.Next = list1
			ptr, list1 = ptr.Next, list1.Next
		} else {
			ptr.Next = list2
			ptr, list2 = ptr.Next, list2.Next
		}
	}
	if list1 != nil {
		ptr.Next = list1
	} else if list2 != nil {
		ptr.Next = list2
	}
	return header.Next
}
