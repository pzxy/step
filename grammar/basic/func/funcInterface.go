package main

import "fmt"

/**
函数实现接口
*/

//创建一个接口  结构体 和 函数均可实现接口
type Invoker interface {
	Call(interface{})
}

//函数定义为类型
type FuncCaller func(interface{})

//实现Invoker的Call
func (f FuncCaller) Call(p interface{}) {
	//调用f()函数本体
	f(p) //不是f.Call(p)
}

//函数来源是命名函数,匿名函数或者闭包
//FuncCaller函数 无需被实例化  只需要把函数类型转换为FuncCall类型即可

func main() {
	//声明接口变量
	var invoker Invoker
	//将匿名函数转化为FuncCaller类型 再赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("将匿名函数转化为FuncCaller类型 再赋值给接口", v)
	})
	invoker.Call("点用接口")
	//总结 函数实现接口  就是调用函数本体 然后匿名函数再赋值给接口 让接口调用
}
