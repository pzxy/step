package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
goroutine可能的切换点(可能切换 其他地方也能切换):
I/O,select
channel
等待锁
函数调用（有时）
runtime2.Gosched()
*/
func main() {
	goTest()
	fmt.Println(len(student))
}
func testGosched() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched() //主动交出控制权 若是不写这个 这个for循环是没机会交出控制权的。
			}
		}(i) //若是这里不将i传进去 就形成了闭包
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

type Student struct {
	name string
	age  int
}

var studentChan chan []*Student
var student []*Student

func goTest() {
	studentChan = make(chan []*Student, 0)
	go toChanValue()
	//func() {
	//	for {
	//		select {
	//		case v,ok := <-studentChan:
	//			if ok {
	//				student = v
	//				return
	//			}
	//		case <- time.NewTicker(time.Second * 2).C:
	//			return
	//		}
	//	}
	//}()
	time.Sleep(time.Second * 3)
	return
}

func toChanValue() {
	var student []*Student
	student = append(student, &Student{name: "小明", age: 12})
	studentChan <- student
}
