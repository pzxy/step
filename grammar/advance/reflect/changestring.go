package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := "Hello"
	c := a + " 2021"
	var b []byte
	tmp := (*string)(unsafe.Pointer(&b))
	*tmp = c
	b[0] = 'h'
	fmt.Println(c)
}

func changeString1() {
	a := "Kylin Lab"
	c := "Hello" + a //跳过编译器检查

	var s []byte

	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&c))
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	sliceHeader.Data = strHeader.Data
	sliceHeader.Len = strHeader.Len
	sliceHeader.Cap = strHeader.Len

	s[0] = 'h'
	fmt.Println(c) //“hello Kylin Lab”
}

func changeString3() {
	a := " 2021"
	c := "Hello" + a
	var b []byte
	tmp := (*string)(unsafe.Pointer(&b))
	*tmp = c

	tmp2 := (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&b)) + unsafe.Sizeof(uintptr(1)) + unsafe.Sizeof(1)))
	*tmp2 = len(b)

	b[0] = 'h'
	fmt.Println(c)
}

func changeString4() {
	a := "Kylin Lab"
	c := "Hello" + a
	var b []byte
	tmp := (*string)(unsafe.Pointer(&b))
	*tmp = c

	//这里为了内存对齐
	bPtr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	tmp2 := (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&b)) + unsafe.Offsetof(bPtr.Cap))) //这个地址就是*reflect.SliceHeader结构中，cap的起始地址
	*tmp2 = len(b)
	b[0] = 'h'
	fmt.Println(c)
}

func changeString2() {
	/**
	关于string，我们还要啰嗦一点，Go语言中string变量的内容默认是不会被修改的，
	而我们通过给string变量整体赋新值的方式来改变它的内容时，实际上会重新分配它的底层数组。
	而string类型字面量的底层数组会被分配到只读数据段，执行阶段会发生错误。

	而运行时动态拼接而成的string变量，它的底层数组不在只读数据段，而是由Go语言在语法层面阻止对字符串内容的修改行为。

	若我们利用unsafe让一个[]byte复用这个字符串c的底层数组，就可以绕过Go语法层面的限制，修改底层数组的内容了。
	*/
	a := "2021"
	c := "Hello " + a
	var b []byte
	tmp := (*string)(unsafe.Pointer(&b))
	*tmp = c
	tmp2 := (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&b)) + 16))
	*tmp2 = len(b)
	b[0] = 'h'
	fmt.Println(c)
	fmt.Println(string(b))
}
