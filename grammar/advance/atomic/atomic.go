package main

import (
	"fmt"
	"sync"
	"time"
)

/**
传统同步机制:
WaitGroup
Mutex
Cond
*/

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {

	//只在下面的代码区加锁
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *atomicInt) get() int {
	return int(a.value)
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Second)
	fmt.Println(a)
}
