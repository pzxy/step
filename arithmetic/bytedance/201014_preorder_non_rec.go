package bytedance

import "fmt"

/**
非递归打印前序遍历
*/

func preOrderNonRec(root *TreeNode) {
	stack := []*TreeNode{root}
	r := root
	for len(stack) > 0 {
		r = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if r != nil {
			fmt.Println(r.Val)
			stack = append(stack, r.Right)
			stack = append(stack, r.Left)
		}

	}
}

/**
非递归打印中序遍历
*/
func lastOrderNonRec(root *TreeNode) {
	stack := []*TreeNode{root}
	r := root
	for len(stack) > 0 {
		r = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if r != nil {
			if r.Right != nil || r.Left == nil {
				fmt.Println(r.Val)
				continue
			}
			if r.Right == nil && r.Left != nil {
				stack = append(stack, r.Left)
				stack = append(stack, r)
			}
			if r.Right == nil && r.Left != nil {
				stack = append(stack, r.Left)
				stack = append(stack, r.Right)
			}
		}

	}
}
