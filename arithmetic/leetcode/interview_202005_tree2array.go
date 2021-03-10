package leetcode

import (
	"strconv"
	"strings"
)

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
			if r == nil { //添加进去下面两层
				ret = append(ret, -1)
				Q = append(Q, nil)
				Q = append(Q, nil)
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

func Serialize1(root *TreeNode) *TreeNode {
	// write code here

	return deserialize(serialize(root))
}

func serialize1(root *TreeNode) string {
	if root == nil {
		return "{}"
	}
	Q := make([]*TreeNode, 1)
	Q[0] = root
	var ret []string
	for len(Q) > 0 {
		node := Q[0]
		Q = Q[1:]

		if node == nil {
			ret = append(ret, "#")
			continue
		}

		Q = append(Q, node.Left)
		Q = append(Q, node.Right)
		ret = append(ret, strconv.Itoa(node.Val))
	}
	return "{" + strings.Join(ret, ",") + "}"
}

func deserialize(str string) *TreeNode {
	s := str[1 : len(str)-1]
	if s == "" {
		return nil
	}
	ss := strings.Split(s, ",")
	return deserializeDo(ss, 0)
}
func deserializeDo(s []string, idx int) *TreeNode {
	if len(s) == 0 || idx > len(s)-1 {
		return nil
	}
	if s[idx] == "#" {
		return nil
	}
	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(s[idx])
	root.Left = deserializeDo(s, idx*2+1)
	root.Right = deserializeDo(s, idx*2+2)
	return root
}
