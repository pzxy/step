package main

import (
	"fmt"
)

func main() {

	var s []int
	s = append(s, 3)
	s = append(s, 3)
	s = append(s, 3)
	s = append(s, 3)
	s1 := make([]int,0)
	s1 = s
	demo12(s1)

	fmt.Println(s)
}

func demo12(p []int) {
	p = append(p,p[2:]...)
	fmt.Println("p",p)
}
