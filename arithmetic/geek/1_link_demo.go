package main

import (
	"errors"
	"fmt"
	"step/utils/log"
)

/**
链表
主要有涉及的有找是否有中环，
链表反转
删除倒数n个节点
查找中间节点等
*/
func main() {
	//printLink(removeNthFromEnd(newLink(10), 7))
	//linkRev()
	//printLink(orderLinkMerge(newLink(10), newLink(10)))
	//linkDeleteN()
	//fmt.Println(linkMinNode(newLink(10)).v)

	fmt.Println(linkCircle(newLinkCircle(100)))
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
func newLinkCircle(n int) *node {
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
	p.next = root
	return root
}

func printLink(root *node) {
	for root != nil {
		fmt.Printf("%d\t", root.v)
		root = root.next
	}
}
func printLinkCircle(root *node) {
	p := root
	for root != nil {
		fmt.Printf("%d\t", root.v)
		root = root.next
		if p == root {
			fmt.Printf(" Second appearance : %d\t", root.v)
		}
	}
}

func revLink(head *node) *node {
	var pre, next *node

	for head != nil {
		next = head.next
		head.next = pre

		pre = head
		head = next
	}
	return pre
}

func linkDeleteN() {
	root := newLink(100)
	if !delNode4N(root, 11) {
		fmt.Println("error")
	}
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
		revRoot = revRoot.next
		root = revLink(revRoot)
		return true
	}

	preNode := returnNPre(revRoot, n)
	deleteNextNode(preNode)
	root = revLink(revRoot)
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
		return
	}
	if pre.next == nil {
		log.ErrLog(errors.New("pre next node is nil"))
		return
	}
	pre.next = pre.next.next
}

/**
删除倒数第n个点，不用链表反转
*/

func removeNthFromEnd(head *node, n int) *node {

	if head == nil {
		return nil
	}

	if n < 0 {
		return head
	}

	i, j := head, head
	z := 1

	for ; i.next != nil; i = i.next {
		if z <= n {
			z++
			continue
		}
		j = j.next
	}
	if z >= n {
		j.next = j.next.next
	}

	return head
}

/**
链表中环的检测
*/

/**
两个有序链表的合并
*/

func orderLinkMerge(head1 *node, head2 *node) *node {
	cur1 := head1
	cur2 := head2
	var head *node
	if cur1.v < cur2.v {
		head = cur1
		cur1 = cur1.next
	} else {
		head = cur2
		cur2 = cur2.next
	}
	curNode := head
	for cur1 != nil || cur2 != nil {
		if cur1.v < cur2.v {
			curNode.next = cur1
			cur1 = cur1.next
			curNode = curNode.next
		} else {
			curNode.next = cur2
			cur2 = cur2.next
			curNode = curNode.next
		}

		if cur1 == nil {
			curNode.next = cur2
			break
		}

		if cur2 == nil {
			curNode.next = cur1
			break
		}
	}
	return head
}

/**
求链表的中间节点
*/

func linkMinNode(head *node) *node {
	step := head
	twoStep := head
	for twoStep != nil {
		if twoStep = twoStep.next; twoStep == nil {
			return step
		}
		if twoStep = twoStep.next; twoStep == nil {
			return step
		}
		step = step.next
	}
	return step
}

/**
检测环
*/
func linkCircle(head *node) bool {

	fast, slow := head, head

	for fast != nil && slow != nil {
		if fast = fast.next; fast == nil {
			return false
		}
		fast = fast.next
		slow = slow.next
		if fast == slow {
			return true
		}
	}
	return false
}
