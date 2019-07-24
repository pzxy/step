package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
向channel中放值都是阻塞的，若是没有接受的话就会panic，
但是使用select来调度的话就不会有这个问题

需要注意的是，在select中为nil的chan是阻塞的是不会调度的哪里的 并不是chan中的值是nil 是指channel本身是nil
*/
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

//这里会有消费过慢的问题导致数据被覆盖掉
func select1() {
	c1, c2 := generator(), generator()
	w := createWorker(0)
	n := 0
	haseValue := false
	for {
		var activeWorker chan<- int //这个activeWorker本身是nil
		if haseValue {
			activeWorker = w
		}
		select {
		case n = <-c1:
			haseValue = true
		case n = <-c2:
			haseValue = true
		case activeWorker <- n: //刚开始activeWorker为nil 可以忽略这里，然后有值后activeWorker成了w了
			haseValue = false
		}
	}
}

//改进select1 增加缓冲区
func select2() {
	c1, c2 := generator(), generator()
	w := createWorker(0)
	var values []int
	for {
		var activeWorker chan<- int //这个activeWorker本身是nil
		var activeValue int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue: //刚开始activeWorker为nil 可以忽略这里，然后有值后activeWorker成了w了
			values = values[1:]
		}
	}
}

//增加计时器
func select3() {
	c1, c2 := generator(), generator()
	w := createWorker(0)
	tm := time.After(time.Second * 10)
	tick := time.Tick(time.Second)
	var values []int
	for {
		var activeWorker chan<- int //这个activeWorker本身是nil
		var activeValue int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue: //刚开始activeWorker为nil 可以忽略这里，然后有值后activeWorker成了w了
			values = values[1:]
		case <-time.After(800 * time.Millisecond): //每次select调度间隔800毫秒 若没有调度 则打印输出
			fmt.Println("timeout...")
		case <-tick: //每过1秒 定时
			fmt.Println("tick...")
		case <-tm:
			fmt.Println("886") //全局时间
			return
		}
	}
}
func main() {
	select3()
}
