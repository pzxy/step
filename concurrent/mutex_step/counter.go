package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
*
正常模式下，waiter 都是进入先入先出队列，被唤醒的 waiter 并不会直接持有锁，而是要和新来的 goroutine 进行竞争。
新来的 goroutine 有先天的优势，它们正在 CPU 中运行，可能它们的数量还不少，所以，在高并发情况下，被唤醒的 waiter 可能比较悲剧地获取不到锁，
这时，它会被插入到队列的前面。

如果 waiter 获取不到锁的时间超过阈值 1 毫秒，那么，这个 Mutex 就进入到了饥饿模式。
在饥饿模式下，Mutex 的拥有者将直接把锁交给队列最前面的 waiter。
新来的 goroutine 不会尝试获取锁，即使看起来锁没有被持有，它也不会去抢，也不会 spin，它会乖乖地加入到等待队列的尾部。

如果拥有 Mutex 的 waiter 发现下面两种情况的其中之一，它就会把这个 Mutex 转换成正常模式:
1. 此 waiter 已经是队列中的最后一个 waiter 了，没有其它的等待锁的 goroutine 了；
2. 此 waiter 的等待时间小于 1 毫秒。正常模式拥有更好的性能，因为即使有等待抢锁的 waiter，goroutine 也可以连续多次获取到锁。

饥饿模式是对公平性和性能的一种平衡，它避免了某些 goroutine 长时间的等待锁。在饥饿模式下，优先对待的是那些一直在等待的 waiter。
*/
func main() {
	mutex4()
}
func mutex2() { //有锁
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)
	var count int
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

/*
*
go run -race counter.go

运时会有问题，只有在运行到时才会发现。
*/
func mutex1() {
	/**
	按理说应该是10000才对，但是很小概率会出现。
	这是为什么呢？
	其实，这是因为，count++ 不是一个原子操作，
	它至少包含几个步骤，比如读取变量 count 的当前值，对这个值加 1，把结果再保存到 count 中。
	因为不是原子操作，就可能有并发的问题。

	比如，10 个 goroutine 同时读取到 count 的值为 9527，接着各自按照自己的逻辑加 1，值变成了 9528，
	然后把这个结果再写回到 count 变量。但是，实际上，此时我们增加的总数应该是 10 才对，这里却只增加了 1，
	好多计数都被“吞”掉了。这是并发访问共享数据的常见错误。
	*/
	/**
	go run -race counter.go
	这个警告不但会告诉你有并发问题，
	而且还会告诉你哪个 goroutine 在哪一行对哪个变量有写操作，
	同时，哪个 goroutine 在哪一行对哪个变量有读操作，就是这些并发的读写访问，引起了 data race。
	*/
	wg := sync.WaitGroup{}
	wg.Add(10)
	var count int
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
	/**
	运行 go tools compile -race -S counter.go，可以查看计数器例子的代码，重点关注一下 count++ 前后的编译后的代码：
	*/
}

func mutex3() {
	// 封装好的计数器
	var counter Counter

	var wg sync.WaitGroup
	wg.Add(10)

	// 启动10个goroutine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 执行10万次累加
			for j := 0; j < 100000; j++ {
				counter.Incr() // 受到锁保护的方法
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count())
}

/**
原子操作
go run -race counter.go
时不会有问题
*/

func mutex4() {

	wg := sync.WaitGroup{}
	wg.Add(10)
	var count int32
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				atomic.AddInt32(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(atomic.LoadInt32(&count))
	/**
	运行 go tools compile -race -S counter.go，可以查看计数器例子的代码，重点关注一下 count++ 前后的编译后的代码：
	*/
}

// 线程安全的计数器类型
type Counter struct {
	CounterType int
	Name        string

	mu    sync.Mutex
	count uint64
}

// 加1的方法，内部使用互斥锁保护
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 得到计数器的值，也需要锁保护
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
