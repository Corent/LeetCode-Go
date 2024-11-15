package algorithm1028

import (
	"regexp"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nodes [][2]int

func recoverFromPreorder(traversal string) *TreeNode {
	if len(traversal) == 0 {
		return nil
	}
	reg1, _ := regexp.Compile("-+")
	reg2, _ := regexp.Compile("\\d+")
	s1, s2 := reg1.Split(traversal, -1), reg2.Split(traversal, -1)
	nodes = make([][2]int, len(s1))
	for i := range s1 {
		v, _ := strconv.Atoi(s1[i])
		nodes[i] = [2]int{v, len(s2[i])}
	}
	return recoverFromPreorderCore(0, len(nodes)-1)
}

func recoverFromPreorderCore(from, to int) *TreeNode {
	if from > to {
		return nil
	}
	node := &TreeNode{Val: nodes[from][0]}
	if from == to {
		return node
	}
	idx1 := from + 1
	for idx1 <= to && nodes[idx1][1] != nodes[from][1]+1 {
		idx1++
	}
	if idx1 > to {
		return node
	}
	idx2 := idx1 + 1
	for idx2 <= to && nodes[idx2][1] != nodes[from][1]+1 {
		idx2++
	}
	if idx2 > to {
		node.Left = recoverFromPreorderCore(idx1, to)
	} else {
		node.Left = recoverFromPreorderCore(idx1, idx2-1)
		node.Right = recoverFromPreorderCore(idx2, to)
	}
	return node
}
