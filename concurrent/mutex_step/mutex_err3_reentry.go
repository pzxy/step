package main

import (
	"fmt"
	"sync"
)

/**
场景3"：重入
Mutex 不是可重入的锁
可重入：
当一个线程获取锁时，如果没有其它线程拥有这个锁，
那么，这个线程就成功获取到这个锁。之后，如果其它线程再请求这个锁，就会处于阻塞等待的状态。
但是，如果拥有这把锁的线程再请求这把锁的话，不会阻塞，而是成功返回，
所以叫可重入锁（有时候也叫做递归锁）。
只要你拥有这把锁，你可以可着劲儿地调用，
比如通过递归实现一些算法，调用者不会阻塞或者死锁。


原因：
想想也不奇怪，因为 Mutex 的实现中没有记录哪个 goroutine 拥有这把锁。
理论上，任何 goroutine 都可以随意地 Unlock 这把锁，所以没办法计算重入条件，毕竟，“臣妾做不到啊”！
*/
func foo3(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}

func bar(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

func main() {
	l := &sync.Mutex{}
	foo3(l)
}
