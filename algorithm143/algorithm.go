package algorithm143

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	root1 := &ListNode{Next: head}
	cnt := 0
	for ptr := root1; ptr != nil; ptr = ptr.Next {
		cnt++
	}
	cnt--
	var ptr1, ptr2 *ListNode
	for ptr1, ptr2, cnt = head, root1, cnt>>1; ptr1 != nil; ptr1, cnt = ptr1.Next, cnt-1 {
		if cnt > 0 {
			continue
		}
		ptr2 = ptr2.Next
	}
	root2 := &ListNode{Next: ptr2.Next}
	ptr2.Next = nil
	root2.Next = reverse(root2.Next)
	head = merge(root1.Next, root2.Next)
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev, ptr *ListNode
	for ptr = head; ptr != nil; {
		next := ptr.Next
		ptr.Next = prev
		prev = ptr
		ptr = next
	}
	return prev
}

func merge(head1, head2 *ListNode) *ListNode {
	if head2 == nil {
		return head1
	}
	if head1 == nil {
		return head2
	}
	root, idx := &ListNode{Val: 0}, 0
	var ptr, ptr1, ptr2 *ListNode
	for ptr, ptr1, ptr2 = root, head1, head2; ptr1 != nil && ptr2 != nil; idx++ {
		if idx&1 == 0 {
			ptr.Next = ptr1
			ptr = ptr.Next
			ptr1 = ptr1.Next
		} else {
			ptr.Next = ptr2
			ptr = ptr.Next
			ptr2 = ptr2.Next
		}
		ptr.Next = nil
	}
	if ptr1 != nil {
		ptr.Next = ptr1
	} else if ptr2 != nil {
		ptr.Next = ptr2
	}
	return root.Next
}
