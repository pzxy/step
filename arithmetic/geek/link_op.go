package main

import (
	"fmt"
	"step/utils/log"
)

func main() {

	linkRev()
	linkDeleteN()
}

/**
链表反转
*/
func linkRev() {
	printLink(newLink(6))
	fmt.Println()
	printLink(revLink(newLink(6)))
}

type node struct {
	v    int
	next *node
}

func newLink(n int) *node {
	root := &node{
		v: 0,
	}
	p := root
	for i := 1; i < n; i++ {
		n := &node{
			v: i,
		}
		p.next = n
		p = n
	}
	return root
}

func printLink(root *node) {
	for {
		if root == nil {
			return
		}
		fmt.Print(root.v)
		root = root.next
	}
}

func revLink(root *node) *node {
	p := root.next
	pre := root
	root.next = nil
	var nextTemp *node
	for {
		if p == nil {
			return pre
		}
		nextTemp = p.next
		p.next = pre
		pre = p
		p = nextTemp
	}
}

func linkDeleteN() {
	root := newLink(10)
	delNode4N(root, 5)
	printLink(root)
}

/**
删除倒数第n个点
*/

func delNode4N(root *node, n int) bool {

	if n <= 0 || root == nil {
		return false
	}

	revRoot := revLink(root)

	if n == 1 {
		root = revRoot.next
		return true
	}

	preNode := returnNPre(revRoot, n)
	deleteNextNode(preNode)
	root = revLink(preNode)
	return true
}

func returnNPre(root *node, n int) *node {
	if n == 2 {
		return root
	}
	return returnNPre(root.next, n-1)
}

func deleteNextNode(pre *node) {
	if pre == nil {
		log.ErrLog(errors.New("pre node is nil"))
	}
	pre.next = pre.next.next
	pre.next.next = nil
}
