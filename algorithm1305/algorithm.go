package algorithm1305

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) (ans []int) {
	stack1, stack2 := []*TreeNode{}, []*TreeNode{}
	if root1 != nil {
		for root1.Left != nil {
			stack1 = append(stack1, root1)
			root1 = root1.Left
		}
		stack1 = append(stack1, root1)
	}
	if root2 != nil {
		for root2.Left != nil {
			stack2 = append(stack2, root2)
			root2 = root2.Left
		}
		stack2 = append(stack2, root2)
	}
	for len(stack1) > 0 || len(stack2) > 0 {
		m, n := len(stack1), len(stack2)
		if m > 0 && n > 0 {
			i, j := stack1[m-1], stack2[n-1]
			if i.Val < j.Val {
				ans = append(ans, i.Val)
				stack1 = stack1[:m-1]
				if i.Right != nil {
					i = i.Right
					for i.Left != nil {
						stack1 = append(stack1, i)
						i = i.Left
					}
					stack1 = append(stack1, i)
				}
			} else {
				ans = append(ans, j.Val)
				stack2 = stack2[:n-1]
				if j.Right != nil {
					j = j.Right
					for j.Left != nil {
						stack2 = append(stack2, j)
						j = j.Left
					}
					stack2 = append(stack2, j)
				}
			}
		} else {
			for ; m > 0; m = len(stack1) {
				i := stack1[m-1]
				ans = append(ans, i.Val)
				stack1 = stack1[:m-1]
				if i.Right != nil {
					i = i.Right
					for i.Left != nil {
						stack1 = append(stack1, i)
						i = i.Left
					}
					stack1 = append(stack1, i)
				}
			}
			for ; n > 0; n = len(stack2) {
				i := stack2[n-1]
				ans = append(ans, i.Val)
				stack2 = stack2[:n-1]
				if i.Right != nil {
					i = i.Right
					for i.Left != nil {
						stack2 = append(stack2, i)
						i = i.Left
					}
					stack2 = append(stack2, i)
				}
			}
		}
	}
	return
}
