package algorithm105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var Preorder, Inorder []int

func buildTree(preorder []int, inorder []int) *TreeNode {
	Preorder, Inorder = preorder, inorder
	return buildTreeCore(0, len(preorder)-1, 0, len(inorder)-1)
}

func buildTreeCore(from1, to1, from2, to2 int) *TreeNode {
	if to1-from1 != to2-from2 || from1 > to1 {
		return nil
	}
	if from1 == to1 {
		return &TreeNode{Val: Preorder[from1]}
	}
	root, idx := &TreeNode{Val: Preorder[from1]}, -1
	for i := from2; i <= to2; i++ {
		if Inorder[i] == root.Val {
			idx = i
			break
		}
	}
	if idx < 0 {
		return nil
	}
	leftLen := idx - from2
	root.Left = buildTreeCore(from1+1, from1+leftLen, from2, idx-1)
	root.Right = buildTreeCore(from1+leftLen+1, to1, idx+1, to2)
	return root
}
