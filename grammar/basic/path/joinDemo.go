package main

import (
	"fmt"
	"path"
)

func main() {
	testJoin()
}

//输出一个 目录类型的结构
func testJoin() {
	pathStr := path.Join("a", "b", fmt.Sprintf("%v%v", "c", ".txt"))
	fmt.Printf("path.join : %v", pathStr)
}
