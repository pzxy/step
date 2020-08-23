package main

import (
	"fmt"
	"runtime"
)

func main() {

	test(2)

}

func test(skip int) {
	call(skip)
}

func call(skip int) {
	//一般用于日志，入参数表示调用该方法的第几层，0就是 这一行，1是call（）这个方法调用的地方，2是调用call（）方法的test（）调用的地方。以此类推
	pc, file, line, ok := runtime.Caller(skip)
	pcName := runtime.FuncForPC(pc).Name() //获取函数名
	fmt.Println(fmt.Sprintf("%v   %s   %d   %t   %s", pc, file, line, ok, pcName))
}
