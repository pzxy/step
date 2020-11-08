package main

import (
	"flag"
	"fmt"
)

//定义一个字符串变量，并制定默认值以及使用方式
var species = flag.String("species", "gopher", "the species we are studying")

//这里的species是个指针,指向 flag.String(name,value,usage)中的value
//定义一个int型字符
var nums = flag.Int("ins", 1, "ins nums")

func main() {

	// 上面定义了两个简单的参数，在所有参数定义生效前，需要使用flag.Parse()来解析参数
	//其实这里解析只是获取命令行参数来解析,若是不使用命令行参数的话,这句不写也行
	flag.Parse()
	//这里就是运行的时候使用了命令行 参数 这个3把原来的值3覆盖掉了
	//$ go run flag-test.go -ins 3 -species biaoge
	//a string flag: biaoge
	//ins num: 3

	// 测试上面定义的函数
	fmt.Println("a string flag:", *species)
	fmt.Println("ins num:", *nums)
	if *nums == 4 {
		fmt.Println("ok")
	}
	//默认打印指令的具体信息,可以这个函数打桩,就是重写
	flag.Usage()
	flag.Usage = func() {
		fmt.Println("重写flag.Usage()方法")
	}
	flag.Usage()
}

//$ go run flag-test.go
//a string flag: gopher
//ins num: 1
