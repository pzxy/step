package leetcode

/**
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]

示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。


说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。
*/

/**
* Definition for TreeNode.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

func (t *TreeNode) insert(data int) {
	//如果存在如何插入?
	for t != nil {
		if t.Val > data {
			if t.Left == nil {
				t.Left = &TreeNode{Val: data}
				return
			}
			t = t.Left
			continue
		}
		if t.Val < data {
			if t.Right == nil {
				t.Right = &TreeNode{Val: data}
				return
			}
			t = t.Right
			continue
		}
	}
}

//二叉搜索树才能这样写
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return nil
	}
	for root != nil {
		if (root.Val > p.Val && root.Val < q.Val) || (root.Val > q.Val && root.Val < p.Val) || root.Val == p.Val || root.Val == q.Val {
			return root
		}

		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		}
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		}
	}
	return root
}

//二叉搜索树才能这样写
func recLowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil || p == nil || q == nil {
		return nil
	}

	if root.Val > p.Val && root.Val > q.Val {
		return recLowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return recLowestCommonAncestor(root.Right, p, q)

	}
	return root
}

//普通的二叉树
/**
从当前根节点的子树中后续遍历，只要p q两个节点不在同一子树中，那么当前节点就是跟节点。

*/
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	var r *TreeNode
	var f func(root, p, q *TreeNode) bool
	f = func(root, p, q *TreeNode) bool {
		if root == nil {
			return false
		}
		left := f(root.Left, p, q)
		right := f(root.Right, p, q)
		min := root == p || root == q
		if (left && right) || (left && min) || (right && min) {
			r = root
		}
		return left || min || right
	}

	f(root, p, q)

	return r
}

func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor3(root.Left, p, q)
	right := lowestCommonAncestor3(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func postorder(root, p, q *TreeNode) bool {
	if root == nil {
		return false
	}

	left := postorder(root.Left, p, q)
	right := postorder(root.Right, p, q)
	min := root == p || root == q
	if (left && right) || (left && min) || (right && min) {

	}
	return left || min || right

}
