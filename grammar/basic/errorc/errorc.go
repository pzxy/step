package main

import (
	"errors"
	"fmt"
)

/**
defer: 最后执行
panic: 向上级返回错误
recover:捕获panic 处理后 panic就此结束 不在向上传递
*/
func main() {
	testErr2()

}
func testErr2() {
	err := errors.New("又创建了一个错误")
	defer func() {
		if error := recover(); error != nil {
			fmt.Printf("出现的错误是:%v\n", error)
		}
	}()
	panic(err)
}
func testErr() {
	err := errors.New("创建错误")
	fmt.Errorf("打印错误")

	defer func() { //放在在栈中 多个defer 按照先进后出执行    一般用于关闭连接等收尾工作
		if error := recover(); error != nil { //recover()捕获panic
			fmt.Println("出现错误,错误是:", error)
		}
	}()

	panic(err) //结束程序 输出错误 若是有defer 执行defer后面的    一般程序尽量不用panic
	//os.Exit(0) //程序立即结束 不打印defer后面内容
}
