package leetcode

import (
	"sync"
)

/**
请实现 copyRandomList 函数，复制一个复杂链表。
在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，
还有一个 random 指针指向链表中的任意节点或者 null。

输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

*/
//思路将每个node和新的node，对应起来。key:oldNode,value:newNode
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	if head.Next == nil && head.Random == nil {
		return &Node{Val: head.Val, Next: head.Next, Random: head.Random}
	}

	var nodeMap sync.Map
	var ret, r, h *Node        //return, temp ret, temp head
	ret = &Node{Val: head.Val} //为了先存下头节点
	nodeMap.Store(head, ret)

	h = head.Next
	r = ret

	for h != nil {
		r.Next = &Node{Val: h.Val}
		r = r.Next
		nodeMap.Store(h, r)
		h = h.Next

	}
	h = head
	r = ret
	for h != nil {
		if v, ok := nodeMap.Load(h.Random); ok {
			if v == nil {
				continue
			}
			r.Random = v.(*Node)
			h = h.Next
			r = r.Next
			continue
		}
		h = h.Next
		r = r.Next
	}
	return ret
}

//方法2，就是说n1 -> n1' -> n2 -> n2' 这样存放。思想就是这样。这个不想实现了
