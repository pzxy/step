package main

import (
	"fmt"
	"time"
)

func main() {
	var num int
	intchan := make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			v, ok := <-intchan
			if !ok {
				break
			}
			fmt.Printf("输出值为 : %v\n", v)
		}
	}()

	for i := 1; i < 10; i++ {
		num = i
		intchan <- num
	}

	time.Sleep(time.Second)

}
