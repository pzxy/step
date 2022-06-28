package leetcode

import (
	"strconv"
	"strings"
)

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param root TreeNode类
 * @return TreeNode类
 */
func Serialize(root *TreeNode) *TreeNode {
	// write code here
	if root == nil {
		return root
	}

	return deserialize(serialize(root))
}

func serialize(r *TreeNode) string {
	if r == nil {
		return ""
	}
	Q := make([]*TreeNode, 1)
	Q[0] = r
	var s []string
	for len(Q) > 0 {
		q := Q[0]
		Q = Q[1:]
		if q == nil {
			s = append(s, "#")
			continue
		}
		s = append(s, strconv.Itoa(q.Val))
		Q = append(Q, q.Left)
		Q = append(Q, q.Right)

	}
	ret := strings.Join(s, ",")
	return ret
}
