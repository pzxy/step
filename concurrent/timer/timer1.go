package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	i := 1
	if i > 0 {
		fmt.Println("1")
	} else if i == -1 {
		fmt.Println("2")
	}

	delay()
}

func delay() {
	t := time.NewTimer(3 * time.Second)
	select {
	case <-t.C:
		log.Println("延迟执行")
	}
}

func waitChannel(conn <-chan string) {
	timer := time.NewTimer(time.Second)
	select {
	case <-conn:
		timer.Stop() //不需要timer时，要关闭定时器
		return
	case <-timer.C:
		log.Println("waitChannel timeout")
		return
	}
}
