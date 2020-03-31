package main

import "fmt"

func main() {
	s := make([]uint, 0)
	s = []uint{1, 2, 3, 4, 5}
	s = append(s, 1, 2, 3, 4)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	for _, i := range s {
		fmt.Println(i)
	}
}
