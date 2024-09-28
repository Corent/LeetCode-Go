package algorithm427

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

var (
	Grid [][]int
)

func construct(grid [][]int) *Node {
	Grid = grid
	return build(0, 0, len(grid))
}

func build(x, y, length int) *Node {
	node := &Node{}
	if length == 1 {
		node.Val, node.IsLeaf = Grid[x][y] == 1, true
		return node
	}
	length >>= 1
	topLeft := build(x, y, length)
	topRight := build(x, y+length, length)
	bottomLeft := build(x+length, y, length)
	bottomRight := build(x+length, y+length, length)
	node.IsLeaf = topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf &&
		(topLeft.Val && topRight.Val && bottomLeft.Val && bottomRight.Val ||
			!topLeft.Val && !topRight.Val && !bottomLeft.Val && !bottomRight.Val)
	if node.IsLeaf {
		node.Val = topLeft.Val
	} else {
		node.Val = true
		node.TopLeft = topLeft
		node.TopRight = topRight
		node.BottomLeft = bottomLeft
		node.BottomRight = bottomRight
	}
	return node
}
