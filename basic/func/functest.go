package main

import (
	"fmt"
)

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
func fib() func() int {
	var a, b = 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type iAdder func(int) (int, iAdder)
type iFib func(int, int) (int, int, iFib)

func fib2(a int, b int) iFib {
	return func(a2 int, b2 int) (aa int, bb int, ifib iFib) {
		a2, b2 = b2, a2+b2
		return a2, b2, fib2(a2, b2)
	}
}

func adder2(base int) iAdder {
	return func(i2 int) (i int, adder iAdder) {
		return base + i2, adder2(base + i2)
	}
}

func main() {
	//闭包累加一
	/*	for i := 0; i < 10; i++ {
		fmt.Println(addr()(i))
	}*/
	//闭包累加二
	/*	a := adder2(0)
		for i := 0; i < 10; i++ {
			var s int
			s,a = a(i)
			fmt.Println(s)
		}*/
	//斐波那契 yi
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	//斐波那契 二
	/*f := fib2(0, 1)
	for i := 0; i < 10; i++ {
		fmt.Println(f)
	}*/
}
