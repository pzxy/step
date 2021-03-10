package main

import "fmt"

func main() {
	fmt.Println(NumberOf1(-1))
}

func NumberOf1(n int) int {
	// write code here
	var num uint64
	if n < 0 {

		num = uint64(^(-n) + 1)
	} else {
		num = uint64(n)
	}
	var count int
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}
