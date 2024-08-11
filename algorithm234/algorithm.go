package algorithm234

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	fast, slow, n := head, head, 0
	for fast.Next != nil {
		fast = fast.Next
		if n&1 == 1 {
			slow = slow.Next
		}
		n++
	}
	if fast == slow {
		return true
	}
	node1, node2 := head, slow.Next
	slow.Next = nil
	for node2 = reverse(node2); node2 != nil; {
		if node1.Val != node2.Val {
			return false
		}
		node1, node2 = node1.Next, node2.Next
	}
	return true
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
