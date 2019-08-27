package _interface

import "fmt"

type me struct {
	s string
}
type me2 struct {
	me
}

func (m *me) test() {
	fmt.Println("实现接口")
}

func newTest() (i iftask) {
	m := &me{s: "s"}
	return m
}
