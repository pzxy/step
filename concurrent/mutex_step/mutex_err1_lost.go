package main

import (
	"fmt"
	"sync"
)

/**
场景1：没有成对出现
*/

func foo() {
	var mu sync.Mutex
	defer mu.Unlock()
	fmt.Println("hello world!")
}
