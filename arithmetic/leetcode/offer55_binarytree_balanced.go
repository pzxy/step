package leetcode

import "math"

/**
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回false 。
*/

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return do(root)
}

func do(root *TreeNode) bool {
	r1 := bF4IsBalancedMaxDepth(root.Left)
	r2 := bF4IsBalancedMaxDepth(root.Right)
	if math.Abs(float64(r1-r2)) > 1 {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil && root.Right != nil {
		return do(root.Left) && do(root.Right)
	}
	if root.Left != nil {
		return do(root.Left)
	}
	return do(root.Right)
}

func bF4IsBalancedMaxDepth(root *TreeNode) (ret int) {
	if root == nil {
		return
	}
	Q := make([]*TreeNode, 0)
	Q = append(Q, root)
	for len(Q) > 0 {
		ret++
		tmp := make([]*TreeNode, 0)
		for _, v := range Q {
			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
		}
		Q = make([]*TreeNode, len(tmp))
		copy(Q, tmp)
	}
	return
}
