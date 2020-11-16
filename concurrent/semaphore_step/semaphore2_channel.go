package main

import "sync"

/**
channel实现信号量


据之前的 Channel 类型的介绍以及 Go 内存模型的定义，你应该能想到，
使用一个 buffer 为 n 的 Channel 很容易实现信号量，比如下面的代码，
我们就是使用 chan struct{}类型来实现的。在初始化这个信号量的时候，
我们设置它的初始容量，代表有多少个资源可以使用。
它使用 Lock 和 Unlock 方法实现请求资源和释放资源，正好实现了 Locker 接口。
*/

// Semaphore 数据结构，并且还实现了Locker接口
type semaphore2 struct {
	sync.Locker
	ch chan struct{}
}

// 创建一个新的信号量
func NewSemaphore(capacity int) sync.Locker {
	if capacity <= 0 {
		capacity = 1 // 容量为1就变成了一个互斥锁
	}
	return &semaphore2{ch: make(chan struct{}, capacity)}
}

// 请求一个资源
func (s *semaphore2) Lock() {
	s.ch <- struct{}{}
}

// 释放资源
func (s *semaphore2) Unlock() {
	<-s.ch
}
