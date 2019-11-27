package main

import (
	"fmt"
	"runtime"
)

func main() {
	test1()
}

func test1() {
	f := runtime.Func{}
	fmt.Printf("%v", f.Entry())
	fmt.Printf("%v", f.Name())
}

func test2() {
	f := runtime.Func{}
	fmt.Printf("%v", f.Entry())
}