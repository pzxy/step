package main

import "fmt"

/**
函数式编程
*/

func addr() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(addr()(i))
	}
}
