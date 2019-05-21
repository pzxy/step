package main

import (
	"fmt"
	"runtime"
)

func customer(ch chan int) {
	for {
		data := <-ch
		if data == 0 {
			break
		}
		fmt.Println(data)
	}
}
func main() {
	ch := make(chan int)

	fmt.Println("asd")
	for {
		var dumy string
		//模拟进程持续进行
		fmt.Scan(&dumy)
		if dumy == "quit" {
			//main本身算一个goroutine 剩下的多少个goroutine 发送多少个0
			for i := 0; i < runtime.NumGoroutine()-1; i++ {
				ch <- 0
			}
			continue
		}
		go customer(ch)
		//输出 goroutine数量
		fmt.Println("goroutines:", runtime.NumGoroutine())
	}
}
