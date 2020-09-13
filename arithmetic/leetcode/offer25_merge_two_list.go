package leetcode

/**
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var mergeHead *ListNode
	if l1.Val < l2.Val {
		mergeHead = l1
		mergeHead.Next = mergeTwoLists(l1.Next, l2)
	} else {
		mergeHead = l2
		mergeHead.Next = mergeTwoLists(l1, l2.Next)
	}
	return mergeHead
}
