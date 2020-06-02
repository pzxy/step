package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "将有数字的字符串333变为字符串数组"
	fmt.Println(string2Array(s))
}

func string2Array(s string) []string {

	return strings.Split(s, "")
}
