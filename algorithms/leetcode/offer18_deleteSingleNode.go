package leetcode

/**
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。

返回删除后的链表的头节点。

注意：此题对比原题有改动

示例 1:

输入: head = [4,5,1,9], val = 5
输出: [4,1,9]
解释: 给定你链表中值为5的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
示例 2:

输入: head = [4,5,1,9], val = 1
输出: [4,5,9]
解释: 给定你链表中值为1的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.


题目保证链表中节点的值互不相同
*/
//offer中的题和这里不太一样，offer中给了要删除的节点a,
//他是用下一个节点b的值来覆盖a。
//这样不用得到a前面的节点，就可以删除a。
func deleteNode(head *ListNode, val int) *ListNode {
	srcHead := head
	if head == nil {
		return nil
	}

	if head.Val == val {
		return head.Next
	}

	for head.Next != nil {
		if head.Next.Val != val {
			head = head.Next
			continue
		}
		head.Next = head.Next.Next
		break
	}

	return srcHead
}
