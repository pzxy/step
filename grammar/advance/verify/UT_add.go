package main

import "fmt"

func main() {
	T_add2()
}

//无️符号溢出,类似截断后的值
func U_add1() {
	var x uint8
	var y uint8
	x = 200
	y = 200
	fmt.Printf("%b,%v", 400, x+y)
}

//补码，正溢出，和减少了 2^8 ,所以值为-126
func T_add1() {
	var x int8
	var y int8
	x = 65
	y = 65
	fmt.Printf("%b,%v", 130, x+y)
}

//补码，负溢出，和增加了 2^8,所以值为，126。。除此之外用补码截取也能算出来。
func T_add2() {
	var x int8
	var y int8
	x = -65
	y = -65
	fmt.Printf("%b,%v", -130, x+y)
}

