package datastruct228

import "testing"

//链表反转
func Test1(t *testing.T) {
	head := newNode(10)
	link := invertLink(head)
	printLink(link)
}
