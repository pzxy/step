package leetcode

/**
删除有序单向链表中的重复的节点。
例如： 1 2 3 3 4
删除后：1 2 4
*/
func deleteDuplication(pHead **ListNode) {
	if pHead == nil || *pHead == nil {
		return
	}
	var pPreNode *ListNode
	var pNode *ListNode = *pHead

	for pNode != nil {
		pNext := pNode.Next
		ifDelete := false
		if pNext != nil && pNext.Val == pNode.Val {
			ifDelete = true
		}

		if !ifDelete {
			pPreNode = pNode
			pNode = pNode.Next
			continue
		}
		value := pNode.Val
		p2Del := pNode
		for p2Del != nil && p2Del.Val == value {
			pNext = p2Del.Next
			p2Del = pNext
		}
		if pPreNode == nil {
			*pHead = pNext
		} else {
			pPreNode.Next = pNext
		}
		pNode = pNext
	}
	return

}

func deleteDuplication2(pHead **ListNode) { //这样**ListNode写的好处就是 *pHead可以随便改。
	if pHead == nil && *pHead == nil {
		return
	}
	var preNode *ListNode
	pNode := *pHead
	for pNode != nil {
		pNext := pNode.Next
		var ifDelete bool
		if pNext != nil && pNext.Val == pNode.Val {
			ifDelete = true
		}
		if !ifDelete {
			preNode = pNode
			pNode = pNode.Next
			continue
		}
		value := pNode.Val
		for pNode != nil && pNode.Val == value {
			pNode = pNode.Next
		}
		if preNode == nil {
			*pHead = pNode
		} else {
			preNode.Next = pNode
		}
	}
}
