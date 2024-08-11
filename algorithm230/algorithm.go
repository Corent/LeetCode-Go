package algorithm230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	K, idx, ans int
)

func kthSmallest(root *TreeNode, k int) int {
	K, idx = k-1, -1
	kthSmallestCore(root)
	return ans
}

func kthSmallestCore(root *TreeNode) {
	if root == nil || idx == K {
		return
	}
	if root.Left != nil {
		kthSmallestCore(root.Left)
	}
	if idx == K {
		return
	}
	if idx = idx + 1; idx == K {
		ans = root.Val
		return
	}
	if root.Right != nil {
		kthSmallestCore(root.Right)
	}
}
