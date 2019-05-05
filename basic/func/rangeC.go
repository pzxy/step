package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() { //这个管道只能放进去一个数字
		c <- 1 //只有消费一个 管道为空才能再次放进去
		c <- 2
		c <- 3
		c <- 4
		close(c)
	}()

	time.Sleep(time.Second)
	for v := range c {
		fmt.Println(v)
	}
}
