package main

import "fmt"

func main() {
	a := [6]int{1, 9, 5, 6, 3, 4}
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)
}
