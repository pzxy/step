package leetcode

import "fmt"

/**
给定一棵二叉搜索树，请找出其中第k大的节点。

示例 1:

输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
  2
输出: 4
示例 2:

输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 4

限制：

1 ≤ k ≤ 二叉搜索树元素个数
*/
func kthLargest2(root *TreeNode, k int) int { //很快
	if root == nil {
		return -1
	}
	var count, kthVal int
	var inOrder func(root *TreeNode, k int)
	inOrder = func(root *TreeNode, k int) {
		if root == nil {
			return
		}
		inOrder(root.Right, k)
		count++
		if count == k {
			kthVal = root.Val
		}
		if count >= k {
			return
		}
		inOrder(root.Left, k)
	}
	inOrder(root, k)
	return kthVal
}

var skip int
var res int

func kthLargest(root *TreeNode, k int) int {
	skip = k
	res = 0
	dfs4kthLargest(root)
	return res
}

func dfs4kthLargest(root *TreeNode) {
	if root != nil {
		dfs4kthLargest(root.Right)
		if skip <= 0 {
			return
		}
		skip--
		if skip == 0 {
			res = root.Val
			return
		}
		fmt.Println(root.Val)
		dfs4kthLargest(root.Left)
	}
}
