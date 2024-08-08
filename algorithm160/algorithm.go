package algorithm160

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	ptr := headA
	for ptr.Next != nil {
		ptr = ptr.Next
	}
	ptr.Next = headB
	ans := detectCycle(headA)
	ptr.Next = nil
	return ans
}

// algorithm 142
func detectCycle(head *ListNode) *ListNode {
	var fast, slow *ListNode
	for fast = head; fast != nil && fast != slow; {
		if fast = fast.Next; fast != nil {
			fast = fast.Next
		}
		if slow == nil {
			slow = head
		}
		slow = slow.Next
	}
	if fast == nil {
		return nil
	}

	for slow = head; fast != slow; slow = slow.Next {
		if fast = fast.Next; fast == slow {
			break
		}
	}
	return fast
}
