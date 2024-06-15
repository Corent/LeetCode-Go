package algorithm019

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	head := &ListNode{Val: 1}
	for ptr, i := head, 2; i <= 5; i++ {
		v := i
		ptr.Next = &ListNode{Val: v}
		ptr = ptr.Next
	}
	head = removeNthFromEnd(head, 2)
	for ; head != nil; head = head.Next {
		t.Log(head.Val)
	}
}
