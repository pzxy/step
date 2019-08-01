package main

import "fmt"

//翻转单向链表
type link struct {
	next  *link
	value int
}

//每隔3个反转单数 例:123 456 7   321 654 7
func (head *link) revlink() *link {
	var newHead *link
	newHead = head.next.next
	for head != nil && head.next != nil && head.next.next != nil {
		//拷贝三个值
		a := head
		b := head.next
		c := head.next.next
		//下一个路径
		var d *link
		if head.next.next.next != nil {
			d = head.next.next.next
		}
		//总是指向下次循环的头结点

		//拼接互相链接的三个

		c.next = b
		c.next.next = a
		//从新开始

		head = d
	}
	return newHead
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
