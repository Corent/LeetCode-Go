package algorithm108

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var NUMS []int

func sortedArrayToBST(nums []int) *TreeNode {
	NUMS = nums
	return sortedArrayToBSTCore(0, len(nums)-1)
}

func sortedArrayToBSTCore(from, to int) *TreeNode {
	if from > to {
		return nil
	}
	mid := (from + to) >> 1
	root := &TreeNode{Val: NUMS[mid]}
	root.Left = sortedArrayToBSTCore(from, mid-1)
	root.Right = sortedArrayToBSTCore(mid+1, to)
	return root
}
