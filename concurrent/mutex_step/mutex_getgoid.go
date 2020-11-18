package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(GoID())
}

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	//str := (*string)(unsafe.Pointer(&buf))
	//fmt.Println(*str)
	fmt.Println(string(buf[:n])) //学到了[]byte转string
	// 得到id字符串
	//strings.TrimPrefix(string(buf[:n]), "goroutine ") //返回除了"goroutine "的所有
	//idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]//以空格来切分
	idField := strings.Fields(string(buf[:n]))[1] //这里和上面那两行是等价的
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
