package datastruct228

import (
	"fmt"
	"strconv"
)

type node struct {
	next *node
	v    int
}

//create link
func newNode(n int) *node {
	var head *node
	var curr *node
	for i := 0; i < n; i++ {
		if head == nil {
			head = &node{v: i}
			curr = head
			continue
		}
		curr.next = &node{v: i}
		curr = curr.next
	}
	return head
}

//print link
func printLink(head *node) {
	for head != nil {
		fmt.Print(strconv.Itoa(head.v) + " ")
		head = head.next
	}
	fmt.Println()
	fmt.Println()
}

//link invert
func invertLink(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	var revHead, curr, temNext *node
	curr = head
	for curr != nil {
		temNext = curr.next
		curr.next = revHead
		revHead = curr
		curr = temNext
	}
	return revHead
}
