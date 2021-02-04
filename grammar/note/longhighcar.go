package main

import (
	"fmt"
)

func main() {
	//a := 0
	//fmt.Scan(&a)
	//fmt.Printf("%d\n", a)

	testData := []int{0, 1, 2}
	res := tranTree2(testData)

	result := inOrder(res)
	fmt.Printf("result =%v", res)
}

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func tranTree2(data1 []int) *TreeNode1 {
	rootIdx := 0
	root := &TreeNode1{}
	inOrder4TranTree(root, data1, rootIdx)
	return root
}

func inOrder4TranTree(root *TreeNode1, data []int, rootIdx int) {
	if rootIdx > len(data)/2 {
		return
	}
	root.Val = data[rootIdx]
	root.Left = &TreeNode1{Val: data[rootIdx*2+1]}
	root.Right = &TreeNode1{Val: data[rootIdx*2+2]}
}

func TranArray(node1 *TreeNode1) (ret []int) {
	if node1 == nil {
		return
	}
	Q := []*TreeNode1{node1}
	for len(Q) > 0 {
		r := Q[0]
		ret = append()
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
