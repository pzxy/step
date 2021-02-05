package leetcode

//0, 3, 9, 10, -1, 5, -1
func Tree2Array(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return tree2ArrayDo(root)
}

//思路，层级遍历，Q->层，tmp->每层，nonNilP一层全部空节点，则停止循环
func tree2ArrayDo(root *TreeNode) (ret []int) {
	Q := []*TreeNode{root}
	var tmp []*TreeNode
	var nonNilP = true //队列存在非空节点？继续:结束；
	for len(Q) > 0 && nonNilP {
		nonNilP = false
		tmp, Q = Q, make([]*TreeNode, 0)
		for len(tmp) > 0 {
			r := tmp[0]
			tmp = tmp[1:]
			if r == nil {
				ret = append(ret, -1)
				continue
			}
			ret = append(ret, r.Val)
			if r.Left != nil || r.Right != nil {
				nonNilP = true
			}
			Q = append(Q, r.Left)
			Q = append(Q, r.Right)
		}
	}
	return ret
}
