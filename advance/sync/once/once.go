package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("开始")
			cancel() //取消掉以后并不会让函数停止,只是通知外面去执行了
			fmt.Println("结束")
		}(i)
	}
	<-ctx.Done()
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
}
