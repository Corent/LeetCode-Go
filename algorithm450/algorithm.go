package algorithm450

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if !containNode(root, key) {
		return root
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		successor := findMin(root.Right)
		successor.Right = delMin(root.Right)
		successor.Left = root.Left
		return successor
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}

func containNode(root *TreeNode, key int) bool {
	if root == nil {
		return false
	}
	if root.Val == key {
		return true
	}
	if root.Val >= key {
		return containNode(root.Left, key)
	}
	return containNode(root.Right, key)
}

func findMin(root *TreeNode) *TreeNode {
	cur := root
	for cur != nil {
		if cur.Left == nil {
			return cur
		}
		cur = cur.Left
	}
	return cur
}

func delMin(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root.Right
	}
	root.Left = delMin(root.Left)
	return root
}
