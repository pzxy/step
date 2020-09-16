package bytedance

func demo1P2(add *ListNode, addend *ListNode) *ListNode {
	add = revList(add)
	addend = revList(addend)
	tmpAdd, tmpAddend := add, addend
	for tmpAdd != nil && tmpAddend != nil {
		tmpAdd.Val = tmpAddend.Val + tmpAdd.Val
		tmpAdd = tmpAdd.Next
		tmpAddend = tmpAddend.Next
	}
	return revList(add)
}

func revList(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}
	var head *ListNode
	pNode := l
	for pNode != nil {
		pNext := pNode.Next
		pNode.Next, head = head, pNode
		pNode = pNext
	}
	return head
}

//[0,1,0,2,1,0,1,3,2,1,2,1]
func demo1P3(num []int) int {
	length := len(num)
	if length <= 2 {
		return -1
	}
	signIdx := 0
	signVal := -1
	ret := 0
	for i, v := range num {
		if v != 0 && signIdx == 0 {
			signIdx = i
			signVal = v
			continue
		}
		if v != 0 {
			min := v
			if v > signIdx {
				min = signVal
			}
			ret += (i - signIdx - 1) * min
			signIdx = i
			signVal = v
			continue
		}
	}
	return ret
}
