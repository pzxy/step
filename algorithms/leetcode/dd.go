package leetcode

func d(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var ret []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			ret = append(ret, root.Val)
			stack = append(stack, root)
			root = root.Left
			continue
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right

	}
	return ret
}
