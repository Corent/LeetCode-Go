package algorithm023

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	mergeKListsCore(lists, 0, len(lists)-1)
	return lists[0]
}

func mergeKListsCore(lists []*ListNode, from, to int) {
	if from == to {
		return
	}
	mid := (from + to) / 2
	mergeKListsCore(lists, from, mid)
	mergeKListsCore(lists, mid+1, to)
	lists[from] = merge2Lists(lists[from], lists[mid+1])
}

func merge2Lists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head := &ListNode{}
	ptr := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
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
	return head.Next
}
