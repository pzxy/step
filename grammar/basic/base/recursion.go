package main

import "fmt"

func add(a int, b int) int {
	if a == 0 {
		return b
	}
	return add(a-1, b+1)
}

func main() {
	fmt.Println(add(100000000, 0))
}

