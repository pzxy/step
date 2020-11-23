package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			c <- i
		}
		return
	}()
	//使用 range时，写的goroutine退出时，触发死锁。直接panic
	for v := range c {
		fmt.Print(v)
	}

	select {}
}
