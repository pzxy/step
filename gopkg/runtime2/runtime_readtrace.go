package main

import (
	"fmt"
	"runtime"
	"time"
)

type student struct {
	a int
}

func main() {
	runtime.StartTrace()
	s := student{}
	fmt.Printf("%v\n", s)
	trace := runtime.ReadTrace()
	runtime.Stack(trace, true)
	fmt.Printf("read :%s\n", trace)
	runtime.StartTrace()
	time.Sleep(time.Second * 3)
}
