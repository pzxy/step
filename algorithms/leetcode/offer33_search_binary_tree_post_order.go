package leetcode

/**
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。
如果是则返回true，否则返回false。
假设输入的数组的任意两个数字都互不相同。

参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true


提示：

数组长度 <= 1000

思路：[1,3,2,6,5]，搜索树：左子树的节点都小于根结点，又子树的节点都大于根结点。
5是根节点，找到比5大的节点6，那么数组中，6到5之前的树都是右子树。数组头到6的树都是左子树。
然后依次递归。
*/

func verifyPostorder(postorder []int) bool {
	length := len(postorder)
	if length == 0 {
		return true
	}
	root := postorder[length-1]
	var idx int
	var isFind bool
	for i, v := range postorder {
		if v > root {
			idx = i
			isFind = true
			break
		}
	}
	var left, right []int

	if isFind { //isFind 为了处理 只有左子树的情况[4,6,5,7]
		left = postorder[:idx]
		right = postorder[idx : length-1]
	} else {
		left = postorder[:length-1]
	}

	for _, v := range left {
		if v > root {
			return false
		}
	}
	for _, v := range right {
		if v < root {
			return false
		}
	}
	return verifyPostorder(left) && verifyPostorder(right)
}
