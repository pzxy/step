package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)

	}()
	go func() {
		fmt.Println(<-c)
		return
	}()

}
