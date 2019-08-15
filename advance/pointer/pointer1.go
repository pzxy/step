package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//临时量和返回结果值是不可寻址的,除非用变量接收一下
	//for语句中的range左边的返回值是可以寻址的.
	people := People{"小明"}
	peoP := &people
	fmt.Printf("%T %v\n", peoP, peoP)
	//uintptr与指针类型是不能直接互相转换的.uintptr是放置地址数据类型的准确值.例如824634204608 这里做的转换 &{小明}->824634204608

	fmt.Printf("%T %v\n", unsafe.Pointer(peoP), unsafe.Pointer(peoP))

	peoPtr := uintptr(unsafe.Pointer(peoP)) //这是结构体的地址值824634204608
	fmt.Printf("%T %v\n", peoPtr, peoPtr)

	namePtr := peoPtr + unsafe.Offsetof(peoP.name) //结构体地址值 + name值在结构体中的地址偏移量,
	fmt.Printf("%T %v\n", namePtr, namePtr)

	nameP := (*string)(unsafe.Pointer(namePtr))
	fmt.Printf("%T %v\n", nameP, nameP)

}

type People struct {
	name string
}
