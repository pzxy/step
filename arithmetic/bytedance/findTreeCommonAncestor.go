package bytedance

func findCommonNode(root *TreeNode, n1 *TreeNode, n2 *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == n1.Val || root.Val == n2.Val {
		return root
	}
	left := findCommonNode(root.Left, n1, n2)
	right := findCommonNode(root.Right, n1, n2)

	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
