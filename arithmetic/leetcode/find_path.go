package leetcode

/**
输入一颗二叉树的根节点和一个整数，按字典序打印出二叉树中结点值的和为输入整数的所有路径。
路径定义为从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。


{10,5,12,4,7},22


[[10,5,7],[10,12]]

*/
func FindPath(root *TreeNode, expectNumber int) [][]int {
	// write code here
	if root == nil {
		return [][]int{}
	}
	ret := make([][]int, 0)
	tmp := make([]int, 0)
	do1(root, expectNumber, tmp, &ret)
	return ret
}

func do1(root *TreeNode, k int, tmp []int, ret *[][]int) {
	if root == nil {
		return
	}
	tmp = append(tmp, root.Val)
	k -= root.Val
	if k == 0 && root.Left == nil && root.Right == nil {
		*ret = append(*ret, tmp)
	}
	do1(root.Left, k, tmp, ret)
	do1(root.Right, k, tmp, ret)
}
