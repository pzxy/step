package leetcode

/**
输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

例如：

给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度3 。

提示：

节点总数 <= 10000
*/

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return bF4maxDepth(root)
}
func bF4maxDepth(root *TreeNode) (ret int) {
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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	dF4maxDepth(root, 1)
	return max
}

var max int

func dF4maxDepth(root *TreeNode, m int) {
	if root == nil {
		return
	}
	if m > max {
		max = m
	}
	dF4maxDepth(root.Right, m+1)
	dF4maxDepth(root.Left, m+1)

}
