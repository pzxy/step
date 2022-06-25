package leetcode

/**
zhi打印树
*/
func Print(pRoot *TreeNode) [][]int {
	ret := make([][]int, 0)
	if pRoot == nil {
		return ret
	}
	Q := make([]*TreeNode, 1)
	var tmpQ []*TreeNode
	Q[0] = pRoot
	left2Right := true

	for len(Q) > 0 {
		r := make([]int, 0)
		for i := len(Q) - 1; i >= 0; i-- {
			q := Q[i]
			if left2Right {
				if q.Left != nil {
					tmpQ = append(tmpQ, q.Left)
				}
				if q.Right != nil {
					tmpQ = append(tmpQ, q.Right)
				}
			} else {
				if q.Right != nil {
					tmpQ = append(tmpQ, q.Right)
				}
				if q.Left != nil {
					tmpQ = append(tmpQ, q.Left)
				}
			}
			r = append(r, q.Val)
		}
		if left2Right {
			left2Right = false
		} else {
			left2Right = true
		}
		ret = append(ret, r)
		Q = tmpQ
		tmpQ = make([]*TreeNode, 0)
	}
	// write code here
	return ret
}
