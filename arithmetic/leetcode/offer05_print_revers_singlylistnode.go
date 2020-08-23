package leetcode

import "fmt"

/**
很明显反转打印是先进后出，先进后出想到了栈，递归就是栈，栈想到了递归。就能写出来代码了

但是呢，递归会栈溢出，用循环会更好，这里我们先用递归了
*/
func printReversSinglyListNode(head *ListNode) {
	if head == nil {
		return
	}
	if head.Next != nil {
		printReversSinglyListNode(head.Next)
	}
	fmt.Println(head.Val)
}

func (n *ListNode) newListNode() {
	for i := 2; i <= 100; i++ {
		n.Next = &ListNode{Val: i}
		n = n.Next
	}
}
