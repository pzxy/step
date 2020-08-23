package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	i32 := new(int32)

	//加法操作 向地址中原值做加法操作,并且返回新值
	atomic.AddInt32(i32, 32)
	addInt := atomic.AddInt32(i32, 2)
	fmt.Println("addInt : ", addInt)

	//保存指针的类型做加法操作
	uiptr := new(uintptr)
	uiptrval := atomic.AddUintptr(uiptr, 12)
	uiptrval = atomic.AddUintptr(uiptr, 23)
	fmt.Println("uiptrval : ", uiptrval)

	//减法操作 u开头的类型不能做减法操作
	addInt = atomic.AddInt32(i32, -32)
	fmt.Println("addInt : ", addInt)

	//CAS 交换值 old值必须和地址中的旧值一样
	bool := atomic.CompareAndSwapInt32(i32, 2, 23)
	fmt.Println("bool : ", bool)
	fmt.Println("addInt : ", *i32)

	//获取值 将地址中的值取出来
	addInt = atomic.LoadInt32(i32)
	fmt.Println("addInt : ", addInt)

	//存值或者更改值
	atomic.StoreInt32(i32, 100)
	fmt.Println("StoreInt32 : ", *i32)

	//存值或者更改值 并把旧值返回
	oldInt := atomic.SwapInt32(i32, 200)
	fmt.Println("oldInt : ", oldInt)

	//原子值 可以放任意类型 但是不能放置nil类型
	var atomicValue atomic.Value
	sli := []string{"1", "2", "3"}
	atomicValue.Store(sli) //存储进去,这样存储进去是不安全的,切片是引用传递,可以在外面更改切片的值,可以拷贝一份再放进去
	//更改他
	sli[0] = "更改他"
	fmt.Println("已经被更改", atomicValue.Load())
	//复制一份
	copySlice := make([]string, len(sli))
	copy(copySlice, sli)
	atomicValue.Store(copySlice)
	//更改他 发现无法被更改回来.
	sli[0] = "1"
	fmt.Println(atomicValue.Load())

}
