package main

import "fmt"

//实现接口
type Live interface { //接口
	run()
	eat()
}
type me struct { //结构体
	Name string
}

func (m *me) run() { //结构体实现方法   入参若是指针类型 初始化接结构体必须传指针
	fmt.Println(m.Name, "我跑步")
}

func (m me) eat() { //结构体实现方法  实现接口全部方法相当于实现接口
	fmt.Println(m.Name, "我吃饭")
}

func main() {
	var l Live = &me{"张三"} //实现了接口给接口赋值
	l.run()
	l.eat()
}
