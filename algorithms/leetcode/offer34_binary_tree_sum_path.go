package leetcode

/**
剑指 Offer 34. 二叉树中和为某一值的路径
难度 :中等
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。
从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]

提示：

节点总数 <= 10000
*/
//书中解法，前序遍历，回溯

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	p := make([]int, 0)
	return pathSumDo(root, 0, sum, &p)
}

func pathSumDo(r *TreeNode, v, sum int, path *[]int) (ret [][]int) {
	if r == nil {
		return
	}
	v += r.Val
	*path = append(*path, r.Val)
	if r.Left == nil && r.Right == nil {
		if v == sum {
			tmp := make([]int, len(*path))
			copy(tmp, *path)
			ret = append(ret, tmp)
		}
		*path = (*path)[:len(*path)-1]
		return
	}
	pathSumDo(r.Left, v, sum, path)
	pathSumDo(r.Right, v, sum, path)
	*path = (*path)[:len(*path)-1]
	return
}
