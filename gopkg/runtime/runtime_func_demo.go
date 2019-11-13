package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GC()
	fmt.Printf("%d",runtime.NumCPU())
}