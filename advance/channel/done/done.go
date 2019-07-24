package main

import (
	"fmt"
	"sync"
)

/**
使用channel等待任务结束：
（一）可以自己写方法，增加一个管道 例如管道c<-true 来通知调用者我的任务结束了。
（二）也可以：
	1 var wg sync.WaitGroup
	2 wg.Add(20)增加任务
	3 wg.Done()任务结束 我的任务结束了（在做任务的结束的地方调用）
	4 wg.Wait()等待任务结束
*/
func doWorker(id int, c chan int, w worker) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func() //抽象一下
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	c := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, c.in, c)
	return c
}
func chanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	wg.Add(20)
	//发数据
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	//发数据
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	//wait for all of them
	wg.Wait()
}
func main() {
	chanDemo()
}
