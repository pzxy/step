package leetcode

import (
	"strconv"
	"strings"
)

//0, 3, 9, 10, -1, 5, -1
func Array2Tree(data1 []int) *TreeNode {
	return array2TreeDo(data1, 0)
}

//思路：和树转数组不同，数组不方便一开始确定大小，为了不越界，所以要append。所以层级遍历树。
//这里直接前序列还原
func array2TreeDo(data []int, rootIdx int) *TreeNode {
	if rootIdx >= len(data) {
		return nil
	}
	if data[rootIdx] == -1 {
		return nil
	}
	root := &TreeNode{Val: data[rootIdx]}
	root.Left = array2TreeDo(data, rootIdx*2+1)
	root.Right = array2TreeDo(data, rootIdx*2+2)
	return root
}

func serialize2(r *TreeNode) string {
	if r == nil {
		return ""
	}
	Q := make([]*TreeNode, 1)
	tmp := make([]*TreeNode, 0)
	Q[0] = r
	var s []string
	hadNode := true
	for len(Q) > 0 && hadNode {
		hadNode = false
		tmp, Q = Q, make([]*TreeNode, 0)
		for len(tmp) > 0 {
			q := tmp[0]
			tmp = tmp[1:]
			if q == nil {
				s = append(s, "#")
				Q = append(Q, nil)
				Q = append(Q, nil)
				continue
			}
			if q.Left != nil || q.Right != nil {
				hadNode = true
			}
			s = append(s, strconv.Itoa(q.Val))
			Q = append(Q, q.Left)
			Q = append(Q, q.Right)
		}
	}
	ret := strings.Join(s, ",")
	return ret
}
