package main

import (
	"fmt"
	"strings"
)

//string 不可改只能基于原字符串进行裁剪、拼接等操作，从而生成一个新的字符串

//Go 语言会把所有被拼接的字符串依次拷贝到一个崭新且足够大的续内存空间中，
//并把持有相应指针值的string值作为结果返回

//一个string值会在底层与它的所有副本共用同一个字节数组。。
// 由于这里的字节数组永远不会被改变，所以这样做是绝对安全的。

//与string值相比，strings.Builder类型的值有哪些优势?
//1 已存在的内容不可变，但可以拼接更多的内容；
//2 减少了内存分配和内容拷贝的次数；
//3 可将内容重置，可重用值

func main() {
	var builder1 strings.Builder
	//拼接
	builder1.WriteString("我")
	builder1.WriteByte(97 - 32)

	fmt.Println(builder1.String())
	//重置后什么也不输出
	builder1.Reset()
	fmt.Println(builder1.String())
	// 省略若干代码。
	fmt.Println("Grow the builder ...")
	builder1.Grow(10) //扩容
	fmt.Printf("The length of contents in the builder is %d.\n", builder1.Len())

	//1.1副本不能调用本来的 方法
	var builder2 strings.Builder
	builder2.Grow(1)
	builder3 := builder2
	//builder3.Grow(1) // 这里会引发 panic。
	_ = builder3
	//1.2重启后位空时复制也没问题
	builder1.Reset()
	builder5 := builder1
	builder5.Grow(1) // 这里不会引发 panic。

	//指针可以被复制
	f2 := func(bp *strings.Builder) {
		(*bp).Grow(1) // 这里虽然不会引发 panic，但不是并发安全的。
		builder4 := *bp
		//builder4.Grow(1) // 这里会引发 panic。
		_ = builder4
	}
	f2(&builder1)

}
