package datastruct

import "fmt"

/**
浏览器前进后退用栈实现，算数运算用两个栈实现
栈可以用数组实现，也能用链表实现，数组实现是顺序栈，链表实现是链表栈
*/

/**
顺序栈
*/

type stack []int

func (s stack) push(v int) stack {
	return append(s, v)
}

func (s stack) pop() (stack, int) {
	l := len(s)
	if l == 0 {
		return s, 0
	}
	return s[:l-1], s[l-1]
}
func stackDemo() {
	s := make(stack, 0)
	s = s.push(10)
	s = s.push(11)
	s = s.push(12)
	fmt.Println(s)
	s, v := s.pop()
	s, v = s.pop()
	fmt.Println(s, v)
	s, v = s.pop()
	fmt.Println(s, v)
}

/**
链表栈
*/

func newStack2() *stack2 {
	return &stack2{nil, 0}
}

type stack2 struct {
	top *node2
	len int
}

type node2 struct {
	v   interface{}
	pre *node2
}

func (this *stack2) push(v interface{}) {
	n := &node2{v, this.top}
	this.top = n
	this.len++
}

func (this *stack2) peek() interface{} {
	if this.len == 0 {
		return nil
	}
	return this.top.v
}

func (this *stack2) pop() interface{} {
	if this.len == 0 {
		return nil
	}
	n := this.top
	this.len--
	this.top = this.top.pre
	return n.v
}

func stack2Demo() {
	stack := newStack2()
	stack.push(1)
	stack.push(2)
	stack.push(3)
	fmt.Println(stack.peek())
	fmt.Println(stack.pop())
	stack.push(4)
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
}
