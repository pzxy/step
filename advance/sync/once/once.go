package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//在 sync.Once对象的生命周期内，只执行一次
//关键是sync.Once 是局部变量还是全局变量，他只能控制在自己的有效范围以内
//func main() {
//	for i:=0;i<10;i++{
//		 Test2()
//	}
//	time.Sleep(time.Second*3)
//}
func Test3() {
	for i := 0; i < 10; i++ {
		Test2()
	}
	time.Sleep(time.Second * 3)
}

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		fmt.Println("开始")
	//		cancel() //取消掉以后并不会让函数停止,只是通知外面去执行了
	//		fmt.Println("结束")
	//	}(i)
	//	contextTest()
	//}
	//<-ctx.Done()
	Test4()
}

func onetest() {
	fmt.Println("执行一次")
}

func onceTest() {
	var one sync.Once
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func(i int) {
			one.Do(onetest)
			wg.Done()
		}(i)
		go func(i int) {
			one.Do(onetest)
			wg.Done()
		}(i)
		fmt.Println(i)
	}
	wg.Wait()
}
func contextTest() {
	//var one sync.Once
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 10; i++ {
		go func(i int) {

			fmt.Println("开始")
			cancel() //取消掉以后并不会让函数停止,只是通知外面去执行了
			fmt.Println("结束")
		}(i)
	}
	<-ctx.Done()
	onceTest()
}

func Test4() {
	for i := 0; i < 10; i++ {
		Test2()
		fmt.Println("test3", i)
	}
	time.Sleep(3 * time.Second)
}

func Test2() {
	var one1 sync.Once
	for i := 0; i < 10; i++ {
		one1.Do(onetest)
		fmt.Println(i)
	}
}
