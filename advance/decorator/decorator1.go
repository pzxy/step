package main

import "fmt"

//装饰器，每个函数执行的时候，先执行自己想执行的那部分。
//但是这个装饰器不适用于任何函数情况。go不支持范行，因为是静态编译语言。

func decorator(f func(s string)) func(s string) {

	return func(s string) {
		fmt.Println("打印日志：函数前。。。")
		f(s)
		fmt.Println("打印日志：函数后。。。")
	}
}

func hello1(s string) {
	fmt.Println(s)
}

func main() {
	decorator(hello1)("hello world")
}
