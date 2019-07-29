package main

import "fmt"

//翻转单向链表
type link struct {
	next  *link
	value int
}

//每隔3个反转单数 例:123 456 7   321 654 7
func (l link) revlink() link {
	if &l != nil && l.next != nil && l.next.next != nil {
		a := l
		b := l.next
		c := l.next.next
		d := l.next.next.next
		var newNode link
		newNode = *c
		c.next = b
		c.next.next = &a
		c.next.next.next = d
		return newNode
	}
	return link{}
}

//反转整条链表
func (l *link) reverseLink() *link {
	//原先的前一个节点
	preLink := new(link)
	preLink = nil
	//原先的后一个节点
	nextLink := new(link)
	nextLink = nil
	for l != nil {
		//保存起来下一个操作的节点
		nextLink = l.next
		//link的下一个节点是 前一个节点
		l.next = preLink
		//更新前一个节点,下个节点原先的前一个节点就是这次的节点本身
		preLink = l
		//更新l 保证循环
		l = nextLink
	}
	return preLink
}
func main() {
	var node link
	node.value = 1
	node.next = &link{value: 2}
	node.next.next = &link{value: 3}
	node.next.next.next = &link{value: 4}
	node.next.next.next.next = &link{value: 5}
	node.next.next.next.next.next = &link{value: 6}
	node.next.next.next.next.next.next = &link{value: 7}

	node.reverseLink().printLink()

}

func (l *link) printLink() {
	for l != nil {
		fmt.Println(l.value)
		if l.next != nil {
			l = l.next
		} else {
			l = nil
		}
	}
}
