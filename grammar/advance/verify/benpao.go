package main

import "fmt"

func main() {
	B()
}
func B() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	for i, _ := range s {
		if i == 3 {

			s = append(s[:3], s[3+1:]...)
		}
	}
	fmt.Println(s)
}
