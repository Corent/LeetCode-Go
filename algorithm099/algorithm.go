package algorithm099

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var previous, mistake1, mistake2 *TreeNode

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	previous, mistake1, mistake2 = nil, nil, nil
	recoverTreeCore(root)
	if mistake1 != nil && mistake2 != nil {
		mistake1.Val, mistake2.Val = mistake2.Val, mistake1.Val
	}
}

func recoverTreeCore(root *TreeNode) {
	if root == nil {
		return
	}
	recoverTreeCore(root.Left)
	if previous != nil && previous.Val >= root.Val {
		if mistake1 == nil {
			mistake1, mistake2 = previous, root
		} else {
			mistake2 = root
		}
	}
	previous = root
	recoverTreeCore(root.Right)
}
