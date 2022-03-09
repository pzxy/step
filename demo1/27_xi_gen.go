package main

import (
	"fmt"
)

func main() {
	fmt.Println(gen(123456, 18))
}

func gen(seq int64, accountType int) int64 {
	accountType *= 1000000000000
	return int64(2000000000000000) + int64(accountType) + seq
}
