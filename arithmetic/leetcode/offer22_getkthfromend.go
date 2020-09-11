package leetcode

/**
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，
本题从1开始计数，即链表的尾节点是倒数第1个节点。
例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。

示例：

给定一个链表: 1->2->3->4->5, 和 k = 2.

返回链表 4->5.
*/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return nil
	}
	fast := head
	for i := 1; i < k; i++ {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		head = head.Next
	}

	return head
}

func getKthFromEnd2(head *ListNode, k int) *ListNode {
	var res []*ListNode //使用数组的办法可以获取到单链表的前面的值
	for head != nil {
		res = append(res, head)
		head = head.Next
	}
	l := len(res)
	if l >= k && k > 0 {
		return res[l-k]
	}
	return nil
}
