package main

import (
	"fmt"
	"runtime"
)

var (
	gChan chan int = make(chan int)
)

func loop(id int) {
	index := 0
	for index < 5 {
		gChan <- id + index
		index++
	}
}

// always print 10 first
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	gChan = make(chan int)

	go loop(1)
	go loop(5)
	go loop(10)
	go loop(15)

	count := 0
	for {
		select {
		case output := <-gChan:
			fmt.Println(output)
			count++
		}
		if count > 15 {
			break
		}
	}
}
