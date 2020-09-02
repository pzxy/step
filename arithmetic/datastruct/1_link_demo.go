package datastruct

import (
	"errors"
	"fmt"
	"step/grammar/codeskill/log"
)

/**
链表
主要有涉及的有找是否有中环，
链表反转
删除倒数n个节点
查找中间节点等
*/

/**
链表反转
*/
func linkRev() {
	printLink(newLink(6))
	fmt.Println()
	printLink(revLinkNode(newLink(6)))
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
	if head == nil {
		return nil
	}

	var revHead, curr, tmpNext *node

	curr = head
	for curr != nil {
		tmpNext = curr.next

		curr.next = revHead
		revHead = curr
		curr = tmpNext
	}

	return revHead
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
		j = j.next //往后走
	}
	if z >= n { //防止越界
		j.next = j.next.next //删除
	}

	return head
}

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
链表中环的检测
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

func revLinkNode(head *node) *node {
	if head == nil {
		return nil
	}
	var revHead, curr, tmpNext *node

	curr = head
	for curr != nil {
		tmpNext = curr.next

		curr.next = revHead
		revHead = curr
		curr = tmpNext
	}

	return revHead
}

/**
链表反转，每间隔三个，以前面试的的题，没写出来
每隔3数个反转链表 例:123 456 789 10，   321 654 987 10
*/

func recLinkNodeInterval3(head *node) *node {
	if head == nil {
		return nil
	}
	var headNode, preNode *node
	curr := head
	count := 1
	sliNodes := make([]*node, 0)
	for curr != nil {
		sliNodes = append(sliNodes, curr)
		curr = curr.next

		if count != 3 {
			count++
			continue
		}
		if preNode != nil {
			preNode.next = sliNodes[2]
		}
		sliNodes[2].next = sliNodes[1]
		sliNodes[1].next = sliNodes[0]
		sliNodes[0].next = curr
		if headNode == nil {
			headNode = sliNodes[2]
		}
		preNode = sliNodes[0]
		count = 1
		sliNodes = make([]*node, 0)
	}
	if headNode == nil {
		headNode = head
	}
	return headNode
}
