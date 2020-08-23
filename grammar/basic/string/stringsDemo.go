package main

import (
	"fmt"
	"strings"
)

func main() {
	hasSuffixDemo("head_xxxx_tail")
}

func hasSuffixDemo(s string) {
	fmt.Printf("strings.HasPrefix : %v \n", strings.HasPrefix(s, "head"))
	//s是否是tail结尾
	fmt.Printf("strings.HasSuffix : %v \n", strings.HasSuffix(s, "tail"))
}
