package main

import (
	"fmt"
	"strings"
)

func main() {
	var a []string
	a = append(a, "1")
	a = append(a, "2")
	a = append(a, "3")

	var b []string
	b = append(b, "2")
	b = append(b, "2")
	b = append(b, "1")

	fmt.Println(Demo1(a, b))
}
func Demo1(a, b []string) bool {

	aa := strings.Join(a, "-")
	fmt.Println(aa)
	bb := strings.Join(b, "-")
	fmt.Println(bb)

	return strings.ContainsAny(aa, bb)
}
func Demo2(a, b []string) bool {



	return strings.ContainsAny(aa, bb)
}
