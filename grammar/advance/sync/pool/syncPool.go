package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//使用方法:
//创建一个sync.Pool对象
//给New属性赋值
//用的时候直接用Get方法取 ,然后断言类型

//GC每两秒清理一次本地池

//粗浅原理:
//一个sync.Pool中有许多本地池(localPool),
//一个本地池对应一个协程
//一个本地池中有私有(private)和共享(share)两种属性
//	私有(private) 只能放置一个对象,只能被自己的协程调用
//	共享(share)能放置多个对象,能被其他协程调用
//Get方法获取对象顺序:
//  从自己本地池(localPool)私有(private)中拿,(不用加锁)
//  从自己以及其他本地池(localPool)中共享(share)中拿,(需要加锁)
//  从其他线程上拿
//  自己New方法创建
func main() {
	testPool2()
}
func testPool2() {
	p := sync.Pool{
		New: func() interface{} {
			c := make(chan int)
			return c
		},
	}
	ctx, cancel := context.WithCancel(context.Background())
	intChan := p.Get().(chan int)
	defer close(intChan)
	p.Put(intChan)

	go func() {
		for {
			c, ok := <-intChan
			if !ok {
				break
			}
			fmt.Printf("读出的数据 %d\n", c)
		}
		cancel()
	}()

	for i := 0; i < 100; i++ {
		intChan <- i
	}

	ctx.Done()

}
func testPool() {
	p := sync.Pool{
		New: func() interface{} {
			return make(chan int)
		},
	}
	ctx, cancel := context.WithCancel(context.Background())
	c := p.Get().(chan int)
	p.Put(c)
	go func() {
		for i := 1; i < 10; i++ {
			c <- i
			time.Sleep(time.Millisecond * 300)
		}
		cancel()
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	<-ctx.Done()
	close(c)
}
