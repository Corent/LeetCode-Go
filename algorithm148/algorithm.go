package algorithm148

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var fast, slow = head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	mid := slow.Next
	slow.Next = nil
	return merge(sortList(head), sortList(mid))
}

func merge(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head, ptr *ListNode
	if list1.Val < list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}
	ptr = head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			ptr.Next = list1
			ptr = ptr.Next
			list1 = list1.Next
		} else {
			ptr.Next = list2
			ptr = ptr.Next
			list2 = list2.Next
		}
	}
	if list1 != nil {
		ptr.Next = list1
	} else if list2 != nil {
		ptr.Next = list2
	}
	return head
}
