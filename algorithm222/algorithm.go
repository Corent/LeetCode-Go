package algorithm222

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right, hLeft, hRight := root, root, 0, 0
	for left != nil {
		left, hLeft = left.Left, hLeft+1
	}
	for right != nil {
		right, hRight = right.Right, hRight+1
	}
	if hLeft == hRight {
		return int(math.Pow(2, float64(hLeft))) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
