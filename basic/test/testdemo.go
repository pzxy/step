package main

import (
	"fmt"
)

func main() {
	var s []string
	s = nil
	fmt.Print(len(s))
}

func demo2asd(){
	s := [...]string{
		0:"0",
		1:"1",
		2:"2",
		3:"3",
		4:"4",
	}
	fmt.Println(s)
}
