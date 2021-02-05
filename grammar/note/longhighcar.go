package main

import (
	"fmt"
)

func main() {
	testData := []int{0, 3, 9, 10, -1, 5, -1}
	res := tranTree(testData)
	result := tranArray(res)
	fmt.Printf("result =%v", result)
}

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func tranTree(data1 []int) *TreeNode1 {
	return tranTreeDo(data1, 0)
}

func tranTreeDo(data []int, rootIdx int) *TreeNode1 {
	if rootIdx >= len(data) || data[rootIdx] == -1 {
		return nil
	}
	root := &TreeNode1{Val: data[rootIdx]}
	root.Left = tranTreeDo(data, rootIdx*2+1)
	root.Right = tranTreeDo(data, rootIdx*2+2)
	return root
}

func tranArray(node1 *TreeNode1) (ret []int) {
	if node1 == nil {
		return
	}
	return tranArrayDo(node1)
}

func tranArrayDo(node1 *TreeNode1) (ret []int) {
	Q := []*TreeNode1{node1}
	var nonNilP = true //队列存在非空节点？继续:结束；
	for len(Q) > 0 && nonNilP {
		tmp := Q
		Q = make([]*TreeNode1, 0)
		nonNilP = false
		for len(tmp) > 0 {
			q := tmp[0]
			tmp = tmp[1:]
			if q == nil {
				ret = append(ret, -1)
				continue
			}
			ret = append(ret, q.Val)

			if q.Left == nil {
				Q = append(Q, nil)
			} else {
				nonNilP = true
				Q = append(Q, q.Left)
			}
			if q.Right == nil {
				Q = append(Q, nil)
			} else {
				nonNilP = true
				Q = append(Q, q.Right)
			}
		}
	}
	return ret
}
