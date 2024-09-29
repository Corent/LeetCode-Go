package algorithm449

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	builder := &strings.Builder{}
	this.serializeCore(root, builder)
	s := builder.String()
	return s[:len(s)-1]
}

func (this *Codec) serializeCore(root *TreeNode, builder *strings.Builder) {
	if root == nil {
		return
	}
	builder.WriteString(strconv.Itoa(root.Val) + ",")
	this.serializeCore(root.Left, builder)
	this.serializeCore(root.Right, builder)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	arr := strings.Split(data, ",")
	return this.deserializeCore(arr, 0, len(arr)-1)
}

func (this *Codec) deserializeCore(arr []string, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}
	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(arr[lo])

	index := hi + 1
	for i := lo + 1; i <= hi; i++ {
		if v, _ := strconv.Atoi(arr[i]); v > root.Val {
			index = i
			break
		}
	}
	root.Left = this.deserializeCore(arr, lo+1, index-1)
	root.Right = this.deserializeCore(arr, index, hi)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
