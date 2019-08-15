package main

import (
	"fmt"
	"sync"
	"time"
)

//type Cond struct {
//	noCopy noCopy  // noCopy可以嵌入到结构中，在第一次使用后不可复制,使用go vet作为检测使用
//	// 根据需求初始化不同的锁，如*Mutex 和 *RWMutex
//	L Locker
//	notify  notifyList  // 通知列表,调用Wait()方法的goroutine会被放入list中,每次唤醒,从这里取出
//	checker copyChecker // 复制检查,检查cond实例是否被复制
//}

//- 1 使用cond := sync.NewCond()来创建,入参必须是Locker类型,参考用例
//- 2 传递cond时必须是引用,或者取地址,千万不要值传递
//- 3 然后加锁,使用wait()等待通知,最后再解锁,wait最好在for循环里面操作.
//- 4 使用Signal()和Broadcast()方法发送通知

//为什么先要锁定条件变量基于的互斥锁，才能调用它的Wait方法?
//>因为条件变量的Wait方法在阻塞当前的 goroutine 之前，会解锁它基于的互斥锁，所以在调用该Wait方法之前，我们必须先锁定那个互斥锁，否则在调用这个Wait方法时，就会引发一个不可恢复的 panic。
//>为什么条件变量的Wait方法要这么做呢？你可以想象一下，如果Wait方法在互斥锁已经锁定的情况下，阻塞了当前的 goroutine，那么又由谁来解锁呢？别的 goroutine 吗？
//先不说这违背了互斥锁的重要使用原则，即：成对的锁定和解锁，就算别的 goroutine 可以来解锁，那万一解锁重复了怎么办？由此引发的 panic 可是无法恢复的。
//如果当前的 goroutine 无法解锁，别的 goroutine 也都不来解锁，那么又由谁来进入临界区，并改变共享资源的状态呢？只要共享资源的状态不变，即使当前的 goroutine 因收到通知而被唤醒，也依然会再次执行这个Wait<方法，并再次被阻塞。
//所以说，如果条件变量的Wait方法不先解锁互斥锁的话，那么就只会造成两种后果：不是当前的程序因 panic 而崩溃，就是相关的 goroutine 全面阻塞。
//
//为什么要用for语句来包裹调用其Wait方法的表达式，用if不可以吗?
//>显然，if语句只会对共享资源的状态检查一次，而for语句却可以做多次检查，直到这个状态改变为止。那为什么要做多次检查呢？
//这主要是为了保险起见。如果一个 goroutine 因收到通知而被唤醒，但却发现共享资源的状态，依然不符合它的要求，那么就应该再次调用条件变量的Wait方法，并继续等待下次通知的到来。
//这种情况是很有可能发生的
func main() {
	cond2()
}
func cond2() {
	var locker = new(sync.Mutex)
	cond := sync.NewCond(locker)
	var wg sync.WaitGroup

	for i := 0; i < 40; i++ {
		wg.Add(1)
		go func(x int) {
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait()
			fmt.Println("我等到了通知,我输出:", x)
			wg.Done()
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("我现在通知其中一个........")
	cond.Signal()
	time.Sleep(time.Second)
	fmt.Println("我再次通知其中一个........")
	cond.Signal()
	time.Sleep(time.Second)
	fmt.Println("我通知所有的..........")
	cond.Broadcast()
	wg.Wait()
}
func cond1() {
	var locker = new(sync.Mutex)
	cond := sync.NewCond(locker) //cond最后返回的必须是指针,入参也要是指针
	var wg sync.WaitGroup

	for i := 0; i < 40; i++ {
		wg.Add(1)
		go func(x int) {
			cond.L.Lock() //必须加锁解锁
			defer cond.L.Unlock()
			cond.Wait()
			fmt.Println("我等到了通知,我输出:", x)
			wg.Done()
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("我现在通知其中一个........")
	cond.Signal()
	time.Sleep(time.Second)
	fmt.Println("我再次通知其中一个........")
	cond.Signal()
	time.Sleep(time.Second)
	fmt.Println("我通知所有的..........")
	cond.Broadcast()
	wg.Wait()
}
