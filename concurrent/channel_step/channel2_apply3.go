package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/**
场景3：
信号通知

chan 类型有这样一个特点：
chan 如果为空，那么，receiver 接收数据的时候就会阻塞等待，
直到 chan 被关闭或者有新的数据到来。
利用这个机制，我们可以实现 wait/notify 的设计模式。

传统的并发原语 Cond 也能实现这个功能，但是，
Cond 使用起来比较复杂，容易出错，而使用 chan 实现 wait/notify 模式就方便很多了。

除了正常的业务处理时的 wait/notify，我们经常碰到的一个场景，就是程序关闭的时候，
我们需要在退出之前做一些清理（doCleanup 方法）的动作。
这个时候，我们经常要使用 chan。

比如，使用 chan 实现程序的 graceful shutdown，
在退出之前执行一些连接关闭、文件 close、缓存落盘等一些动作。
*/

func main() {

}
func demo1() {
	go func() {
		// 执行业务处理
	}()

	// 处理CTRL+C等中断信号
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	// 执行退出之前的清理动作
	//doCleanup()

	fmt.Println("优雅退出")
}

/**
有时候，doCleanup 可能是一个很耗时的操作，比如十几分钟才能完成，如果程序退出需要等待这么长时间，
用户是不能接受的，所以，在实践中，我们需要设置一个最长的等待时间。
只要超过了这个时间，程序就不再等待，可以直接退出。

所以，退出的时候分为两个阶段：
1. closing，代表程序退出，但是清理工作还没做；
2. closed，代表清理工作已经做完。
*/
//所以上面的例子可以改为。
func demo2() {
	var closing = make(chan struct{})
	var closed = make(chan struct{})

	go func() {
		// 模拟业务处理
		for {
			select {
			case <-closing:
				return
			default:
				// ....... 业务计算
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// 处理CTRL+C等中断信号
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	close(closing)
	// 执行退出之前的清理动作
	go doCleanup(closed)

	select {
	case <-closed:
	case <-time.After(time.Second):
		fmt.Println("清理超时，不等了")
	}
	fmt.Println("优雅退出")
}

func doCleanup(closed chan struct{}) {
	time.Sleep((time.Minute))
	close(closed)
}
