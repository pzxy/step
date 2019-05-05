package main

import "fmt"

/**
闭包可以用在许多地方。它的最大用处有两个，
一个是可以读取函数内部的变量，
二 就是让这些变量的值始终保持在内存中，不会在调用后被自动清除。
*/
func main() {
	//字符传
	str := "hello world "
	//匿名函数
	foo := func() {
		//匿名函数中访问str
		str = "hello closure"
	}
	foo()
	//accumulator accumulator2是两个不同的函数实例
	accumulator := Accumulate(1)
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Printf("%p\n", accumulator)
	accumulator2 := Accumulate(10)
	fmt.Println(accumulator2())
	fmt.Println(accumulator())
	fmt.Println(accumulator2())
	fmt.Printf("%p\n", accumulator2)

}

//累加器
func Accumulate(value int) func() int {
	//返回一个闭包
	return func() int {
		value++
		return value
	}
}

//总结 闭包一般都是函数要返回匿名函数  匿名函数中捕获了变量 这个变量在此函数实例中是一直存在的 变量被捕获了
