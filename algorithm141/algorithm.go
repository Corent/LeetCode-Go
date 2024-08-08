package algorithm141

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil {
		if fast = fast.Next; fast == nil {
			return false
		} else if fast == slow {
			return true
		}
		if fast = fast.Next; fast == nil {
			return false
		} else if fast == slow {
			return true
		}
		slow = slow.Next
	}
	return false
}
