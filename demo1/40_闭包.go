package main

import "fmt"

func main() {
	c := 30
	a40(c)()
	fmt.Println(c)
}

func a40(c int) func() {
	c = c + 1
	return func() {
		fmt.Println(c)
	}
}
