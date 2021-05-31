package main

import (
	"log"
	"sync"
	"time"
)

/**
条件变量主要是看条件的变化呢。
Cond对象不可复制使用

关于wait为什么加锁的分析。
1. wait外面加锁
首先是条件变量本身，我们要依靠条件变量来判断需不需要再判断wait了
以为是每次条件改变后都会通知，如果改变条件变量的时候不加锁，那么就可能进入死锁。
为什么会进入死锁？
因为现在的用法是(demo2)，wait后收到通知，当条件不符合要求后（ready==9，只要不等于10 ，就继续wait），要继续调用wait等待。
但是在 (不符合要求---->继续wait等待前)这段时间（此时ready==9），刚好条件符合要求了（ready加到了10，按理说ready不会改变的，但是没加锁），
此时如果已经发了最后的通知了，那么由于此时wait还没调用（未把此协程加入等待队列），那么他就永远不会收到通知，就会进入死锁。

1. 分析结果
条件改变时加锁是为了，判断条件时是否wait时，保证条件不会改变。
wait外面加锁也是为了拿到锁，不让条件改变(我们认为此时wait中并没有解锁加锁这些操作)。

2. wait里面为什么解锁加锁？
其实此时wait外面已经有锁了，
wait中解锁是为了：
1是为了条件能拿到锁改变条件给自己发通知。
2是因为wait可能不止一个协程里面有可能有好多个，他们也需要拿到锁加入等待队列。
wait中加锁是为了：保证条件不在发生变化。

waitGroup和Cond的区别：
WaitGroup 是主 goroutine 等待确定数量的子 goroutine 完成任务；
而 Cond 是等待某个条件满足，这个条件的修改可以被任意多的 goroutine 更新，
而且 Cond 的 Wait 不关心也不知道其他 goroutine 的数量，只关心等待条件。
而且 Cond 还有单个通知的机制，也就是 Signal 方法。

Cond中copyChecker和noCopy的区别：
noCopy和copyChecker都是为了检查Cond对象是否被复制。
noCopy是静态检查，通过go vet 命令行参数检查。
copyChecker是运行时检查用的。

*/
func main() {
	demo1()
}

/**
正常用法
*/
func demo1() {
	c := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 100; i++ {
		go func(i int) {
			//time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			// 加锁更改等待条件
			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("运动员#%d 已准备就绪\n", i)
			// 广播唤醒所有的等待者
			c.Broadcast()
		}(i)
	}
	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	c.L.Unlock()

	//所有的运动员是否就绪
	time.Sleep(time.Second)
	log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")
}

/**
触发死锁
*/
func demo2() {
	c := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 10; i++ {
		go func(i int) {
			//time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			// 加锁更改等待条件
			//c.L.Lock()
			if ready == 9 {
				time.Sleep(time.Second * 2)
			}
			ready++
			//c.L.Unlock()

			log.Printf("运动员#%d 已准备就绪\n", i)
			// 广播唤醒所有的等待者
			c.Broadcast()
		}(i)
	}
	c.L.Lock()
	for ready != 10 {
		if ready == 9 {
			time.Sleep(time.Second * 5)
		}
		//log.Printf("readfy：%v\n", ready)
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	c.L.Unlock()

	//所有的运动员是否就绪
	time.Sleep(time.Second)
	log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")
}
