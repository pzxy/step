package main

import (
	"container/ring"
	"fmt"
)

/**
双向循环链表
*/
func main() {
	r := ring.New(5) //r代表整个循环链表又代表第一个元素
	//向下赋值
	r.Value = 0
	r.Next().Value = 1
	r.Next().Next().Value = 2
	/*	r.Next().Next().Next().Value = 3
		r.Next().Next().Next().Next().Value = 4*/
	r.Prev().Value = 4
	r.Prev().Prev().Value = 3
	//打印指定位置 循环链表 相同位置
	fmt.Println("指定位置值:", r.Move(2).Value)
	fmt.Println("指定位置值:", r.Move(-3).Value)
	//打印全部
	r.Do(func(i interface{}) { //i代表当前执行元素的内容
		fmt.Println(i)
	})
	fmt.Println("**************")
	//追加链表后打印全部
	r1 := ring.New(1)
	r1.Value = 9
	r.Prev().Link(r1)

	r.Do(func(i interface{}) { //i代表当前执行元素的内容
		fmt.Println(i)
	})
	fmt.Println("**************")
	//删除元素
	r.Unlink(1)                //删除当前位置的后面一个元素 n=n%r.len()
	r.Do(func(i interface{}) { //i代表当前执行元素的内容
		fmt.Println(i)
	})
}
