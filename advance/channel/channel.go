package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		/*n,ok := <-c
		if !ok{
			break
		}*/
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}
func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 3
}

func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 3
	close(c)
	//time.Sleep(time.Millisecond)
}
func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	//发数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}
func main() {
	//chanDemo()
	//bufferedChannel() //缓冲区
	//缓冲区 通知接收方我发完了
	channelClose()
}
