package main

import (
	"fmt"
	"sync"
	"time"
)

// 一个[]byte的对象池，每个对象为一个[]byte
var bytePool2 = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func main() {

	a := time.Now().UnixNano()
	// 不使用对象池
	for i := 0; i < 100000000; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}
	b := time.Now().UnixNano()

	// 使用对象池
	for i := 0; i < 100000000; i++ {
		obj := bytePool2.Get().(*[]byte)
		_ = obj
		bytePool2.Put(obj)
	}
	c := time.Now().UnixNano()

	fmt.Println("without pool ", b-a)
	fmt.Println("with    pool ", c-b)
}
