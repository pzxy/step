package leetcode

/**
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树[1,2,2,3,4,4,3] 是对称的。

  1
 / \
 2  2
/ \ / \
3 4 4 3
但是下面这个[1,2,2,null,3,null,3] 则不是镜像对称的:

  1
 / \
 2  2
  \  \
   3  3

示例 1：

输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：

输入：root = [1,2,2,null,3,null,3]
输出：false
*/

func isSymmetric(root *TreeNode) bool {
	//如果是对称树，那么镜像树和原树是相同的。
	//原树的前序遍历（左中右）应该和镜像树（右中左）遍历相同
	//打印的时候为空也打印，因为如果是数字相同的不完全二叉树，即使不对称，输出也是一样的。所以为空也要打印出来。
	return checkSymmetric(root, root)
}

func checkSymmetric(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}

	return checkSymmetric(t1.Left, t2.Right) && checkSymmetric(t1.Right, t2.Left)
}
