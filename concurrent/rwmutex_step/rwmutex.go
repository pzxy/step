package main

import (
	"fmt"
	"sync"
	"time"
)

/**
go的读写锁是写优先。
Go 标准库中的 RWMutex 设计是 Write-preferring 方案。一个正在阻塞的 Lock 调用会排除新的 reader 请求到锁。

当 writer 请求锁的时候，是无法改变既有的 reader 持有锁的现实的，也不会强制这些 reader 释放锁，它的优先权只是限定后来的 reader 不要和它抢。


操作规则：
锁读锁时，如果readCount 大于 0，直接锁。如果为负数，说明有writer在等候，所以此时要通知reader阻塞等候。
解读锁时，readCount减去1，如果有等待的writer，则唤醒这个writer

锁写锁时，先锁住，readerCount设为负值(减去rwmutexMaxReaders)，等现有reder执行完以后再通知写操作进行。
解写锁时，先改变readerCount的值(加上rwmutexMaxReaders)，候再唤醒所有阻塞的reader，然后再解锁。
*/
func main() {
	var counter Counter3
	for i := 0; i < 10; i++ { // 10个reader
		go func() {
			for {
				fmt.Println(counter.Count()) // 计数器读操作
				time.Sleep(time.Millisecond * 500)
			}
		}()
	}

	for { // 一个writer
		counter.Incr() // 计数器写操作
		time.Sleep(time.Second)
	}
}

// 一个线程安全的计数器
type Counter3 struct {
	mu    sync.RWMutex
	count uint64
}

// 使用写锁保护
func (c *Counter3) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 使用读锁保护
func (c *Counter3) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}
