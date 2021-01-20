package main

import (
	"fmt"
	"math/rand"
	"time"
)

//向一个为nil的管道发数据并不会有问题。
//go看起来有点奇怪的特性：向nil channel读写数据会永久堵塞该goroutine,正确使用该特性！！
func main() {
	var c chan int
	go func() {
		for {
			time.Sleep(time.Second)
			c <- rand.Intn(100)
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(<-c)
		}

	}()
	time.Sleep(time.Second * 10)
}
