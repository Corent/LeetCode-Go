package algorithm958

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue, preNode := []*TreeNode{root}, root
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if preNode == nil && node != nil {
			return false
		}
		preNode = node
		if node != nil {
			queue = append(queue, node.Left, node.Right)
		}
	}
	return true
}
