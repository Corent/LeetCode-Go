package algorithm109

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nums []int

func sortedListToBST(head *ListNode) *TreeNode {
	nums = make([]int, 0)
	for ptr := head; ptr != nil; ptr = ptr.Next {
		nums = append(nums, ptr.Val)
	}
	return sortedListToBSTCore(0, len(nums)-1)
}

func sortedListToBSTCore(from, to int) *TreeNode {
	if from > to {
		return nil
	}
	mid := (from + to) >> 1
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedListToBSTCore(from, mid-1)
	root.Right = sortedListToBSTCore(mid+1, to)
	return root
}
