package main

import (
	"fmt"
	"sync"
)

/**
场景2：锁被复制， 造成死锁
*/

type Counter2 struct {
	sync.Mutex
	Count int
}

func main() { //fatal error: all goroutines are asleep - deadlock!
	var c Counter2
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo2(c) // 复制锁
}

// 这里Counter的参数是通过复制的方式传入的
func foo2(c Counter2) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}

/**
解决办法：
你肯定不想运行的时候才发现这个因为复制 Mutex 导致的死锁问题，
那么你怎么能够及时发现问题呢？可以使用 vet 工具，
把检查写在 Makefile 文件中，在持续集成的时候跑一跑，这样可以及时发现问题，及时修复。我们可以使用 go vet 检查这个 Go 文件：
go vet mutex_err2_copy.go


输出：
./mutex_err2_copy.go:22:7: call of foo2 copies lock value: command-line-arguments.Counter2
./mutex_err2_copy.go:26:13: foo2 passes lock by value: command-line-arguments.Counter2
*/
