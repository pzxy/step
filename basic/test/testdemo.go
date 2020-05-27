package main

import "fmt"

func main() {
	s := []int{3, 1, 8, 5, 2, 7, 9, 4, 6, 0}
	n := 5
	for _, v := range s {
		if v == n {
			continue
		}
		fmt.Println(v)
	}
}
