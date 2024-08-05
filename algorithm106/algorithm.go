package algorithm106

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var Inorder, Postorder []int

func buildTree(inorder []int, postorder []int) *TreeNode {
	Inorder, Postorder = inorder, postorder
	return buildTreeCore(0, len(inorder)-1, 0, len(postorder)-1)
}

func buildTreeCore(from1, to1, from2, to2 int) *TreeNode {
	if to1-from1 != to2-from2 || from1 > to1 {
		return nil
	}
	if from2 == to2 {
		return &TreeNode{Val: Postorder[from2]}
	}
	root, idx := &TreeNode{Val: Postorder[to2]}, -1
	for i := from1; i <= to1; i++ {
		if Inorder[i] == root.Val {
			idx = i
			break
		}
	}
	if idx < 0 {
		return nil
	}
	leftLen := idx - from1
	root.Left = buildTreeCore(from1, idx-1, from2, from2+leftLen-1)
	root.Right = buildTreeCore(idx+1, to1, from2+leftLen, to2-1)
	return root
}
