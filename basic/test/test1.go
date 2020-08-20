package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(string('a'))
	fmt.Println(string("a"))
	fmt.Println(string(97))
}

func string2Array(s string) []string {

	return strings.Split(s, "")
}
