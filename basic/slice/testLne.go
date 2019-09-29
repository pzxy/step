package main

import "fmt"

func main() {
	sli := make([]int, 4)
	sli[0] = 1

	fmt.Println(len(sli), cap(sli))
	fmt.Println(sli)
}
