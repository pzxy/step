package main

import "fmt"

//质数
func main() {
	for i := 2; i < 100; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				break
			}
			if i == j+1 {
				fmt.Println(i)
			}
		}
	}
}
