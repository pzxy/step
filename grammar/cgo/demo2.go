package main

/**
前面使用了标准库中已有的函数。现在我们先自定义一个叫作SayHello的C函数来实现打印，然后从Go语言环境中调用这个SayHello()函数：
*/

/*
#include <stdio.h>
static void SayHello(const char* s) {
puts(s);
}
*/
import "C"

func main() {
	C.SayHello(C.CString("Hello, World\n"))
}
