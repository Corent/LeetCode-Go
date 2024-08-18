package algorithm337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	ifRobMap, ifRobNotMap = make(map[*TreeNode]int), make(map[*TreeNode]int)
)

func rob(root *TreeNode) int {
	ifRobMap, ifRobNotMap = make(map[*TreeNode]int), make(map[*TreeNode]int)
	return robYesOrNot(root)
}

func robYesOrNot(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ifRobYes, ok := ifRobMap[root]
	if !ok {
		ifRobYes = root.Val + robNot(root.Left) + robNot(root.Right)
		ifRobMap[root] = ifRobYes
	}
	ifRobNot, ok := ifRobNotMap[root]
	if !ok {
		ifRobNot = robYesOrNot(root.Left) + robYesOrNot(root.Right)
		ifRobNotMap[root] = ifRobNot
	}
	if ifRobYes > ifRobNot {
		return ifRobYes
	}
	return ifRobNot
}

func robNot(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ifRobNot, ok := ifRobNotMap[root]
	if ok {
		return ifRobNot
	}
	rst := robYesOrNot(root.Left) + robYesOrNot(root.Right)
	ifRobNotMap[root] = rst
	return rst
}
