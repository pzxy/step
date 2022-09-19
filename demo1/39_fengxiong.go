package main

import (
	"fmt"
	"time"
)

func funcA() {
	fmt.Println("A")
}
func funcB() {
	fmt.Println("B")
}
func funcC() {
	fmt.Println("C")
}
func funcD() {
	fmt.Println("D")
}

func main() {
	go func() {
		fmt.Println("start")
		for i := 0; i < 100; i++ {
			funcA()
		}
		go funcB()
		funcC()
		funcC()
		funcD()
		go func() {
			fmt.Println("test")
		}()
		fmt.Println("goroutine")
		for i := 0; i < 50; i++ {
			go funcD()
		}
	}()
	for {
		time.Sleep(time.Second * 1000)
	}
}

//func main() {
//	go func() {
//		fmt.Println("hello")
//		go func() {
//			fmt.Println("world")
//		}()
//		for i := 0; i < 50; i++ {
//			fmt.Println("i:", i)
//		}
//		fmt.Println("test")
//	}()
//
//	select {}
//}
