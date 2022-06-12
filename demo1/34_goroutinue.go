package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		f2()
	}
	time.Sleep(time.Second)
}

var x, y int

func f2() {
	go func() {
		fmt.Println(x)
	}()

	go func() {
		fmt.Println(y)
	}()
	x, y = 123, 789
}
