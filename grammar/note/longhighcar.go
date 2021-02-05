package main

import (
	"fmt"
)

func main() {
	testData := []int{6, 4, 8, 0, 0, 7, 9}
	res := tranTree2(testData)

	result := inOrder(res)
	fmt.Printf("result =%v", result)
}

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func tranTree2(data1 []int) *TreeNode1 {
	head := &TreeNode1{}
	root := head
	for i := 0; i <= len(data1)/2; i++ {
		if data1[i] == 0 {
			continue
		}
		root.Val = data1[i]
		if data1[2*i+1] != 0 {
			root.Left = &TreeNode1{Val: data1[2*i+1]}
		}
		if data1[2*i+2] != 0 {
			root.Right = &TreeNode1{Val: data1[2*i+2]}
		}
	}
	return root
}

func inOrder4TranTree(root *TreeNode1, data []int, rootIdx int) {
	if rootIdx >= len(data) {
		return
	}
	root.Left = &TreeNode1{}
	inOrder4TranTree(root.Left, data, rootIdx*2+1)
	root.Val = data[rootIdx]
	root.Right = &TreeNode1{}
	inOrder4TranTree(root.Right, data, rootIdx*2+2)
}

func TranArray(node1 *TreeNode1) (ret []int) {
	if node1 == nil {
		return
	}
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
				ret = append(ret, 0)
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
	return
}

func tranTree(data1 []int) *TreeNode1 {

	var tran func(data []int, index int) *TreeNode1
	tran = func(data []int, index int) *TreeNode1 {
		root := &TreeNode1{}
		if index < len(data) {
			if data[index] == -1 {
				return nil
			}
			root = &TreeNode1{}
			root.Val = data[index]
			root.Left = tran(data, 2*index+1)
			root.Right = tran(data, 2*index+2)
			fmt.Printf("val:=%v", root.Val)
			return root
		}

		return root
	}

	res := tran(data1, 0)

	return res
}

func inOrder(root *TreeNode1) []int {
	p := root
	stack := []*TreeNode1{}
	ans := []int{}
	for len(stack) != 0 || p != nil {
		if p != nil {
			stack = append(stack, p)
			p = p.Left

		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, node.Val)
			p = node.Right

		}

	}
	return ans

}
