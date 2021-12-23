package main

import "fmt"

func main() {
	if s,err := aaaaaa();err != nil {
		fmt.Println("1")
	}else {
		fmt.Println(s)
	}
}

func aaaaaa() (string, error) {
	return "s", nil
}
