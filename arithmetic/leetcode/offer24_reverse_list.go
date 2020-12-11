package leetcode

/**
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

限制：

0 <= 节点个数 <= 5000

*/
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var preNode *ListNode
	pNode := head
	for pNode != nil {
		pNext := pNode.Next
		pNode.Next = preNode
		preNode = pNode
		pNode = pNext
	}
	return preNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var revHead, currNode, nextNode *ListNode
	currNode = head
	for currNode != nil {
		nextNode = currNode.Next

		currNode.Next, revHead = revHead, currNode

		currNode = nextNode
	}
	return revHead
}
