package algorithm687

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans = 0

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	longestUnivaluePathCore(root)
	return ans
}

func longestUnivaluePathCore(root *TreeNode) {
	if root == nil {
		return
	}
	path := 0
	if root.Left != nil && root.Val == root.Left.Val {
		path += 1 + findUnivaluePath(root.Left)
	}
	if root.Right != nil && root.Val == root.Right.Val {
		path += 1 + findUnivaluePath(root.Right)
	}
	ans = maxInt(ans, path)
	if root.Left != nil {
		longestUnivaluePathCore(root.Left)
	}
	if root.Right != nil {
		longestUnivaluePathCore(root.Right)
	}
}

func findUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftPath, rightPath := 0, 0
	if root.Left != nil && root.Val == root.Left.Val {
		leftPath = 1 + findUnivaluePath(root.Left)
	}
	if root.Right != nil && root.Val == root.Right.Val {
		rightPath = 1 + findUnivaluePath(root.Right)
	}
	return maxInt(leftPath, rightPath)
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}
