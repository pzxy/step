package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	demo3()
}

/**
反射接受
*/
func demo3() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	var cases = createRecvCase2(ch1, ch2)

	go func(c chan int) {
		var count = 100
		for {
			ch1 <- count
			count++
			time.Sleep(time.Second)
		}
	}(ch1)

	go func(c chan int) {
		var count = 600
		for {
			ch2 <- count
			count++
			time.Sleep(time.Second)
		}
	}(ch2)

	for len(cases) > 0 {
		chosen, recv, ok := reflect.Select(cases)
		if recv.IsValid() {
			fmt.Println("recv:", chosen, recv, ok)
		}
	}
}

/**
反射 发，收
*/

func demo4() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	var cases = createRecvCase2(ch1, ch2)

	for len(cases) > 0 {
		chosen, recv, ok := reflect.Select(cases)
		if recv.IsValid() {
			fmt.Println("recv:", chosen, recv, ok)
		}
	}
}
func createRecvCase2(chs ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase
	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	return cases
}
