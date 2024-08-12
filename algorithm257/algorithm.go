package algorithm257

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	ans  []string
	path []string
)

func binaryTreePaths(root *TreeNode) []string {
	ans, path = make([]string, 0), make([]string, 0)
	binaryTreePathsCore(root)
	return ans
}

func binaryTreePathsCore(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		path = append(path, strconv.Itoa(root.Val))
		ans = append(ans, strings.Join(path, "->"))
		path = path[:len(path)-1]
		return
	}
	path = append(path, strconv.Itoa(root.Val))
	if root.Left != nil {
		binaryTreePathsCore(root.Left)
	}
	if root.Right != nil {
		binaryTreePathsCore(root.Right)
	}
	path = path[:len(path)-1]
}
