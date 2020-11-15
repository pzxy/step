package main

import (
	"fmt"
	"reflect"
	"time"
)

/**
场景五：
任务编排
共5种，分别是 Or-Done 模式、扇入模式、扇出模式、Stream 和 Map-Reduce。
*/

/**
Or-Done 模式
我们会使用“信号通知”实现某个任务执行完成后的通知机制，
在实现时，我们为这个任务定义一个类型为 chan struct{}类型的 done 变量，等任务结束后，
我们就可以 close 这个变量，然后，其它 receiver 就会收到这个通知。
*/
func or(channels ...<-chan interface{}) <-chan interface{} {
	// 特殊情况，只有零个或者1个chan
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) { //每个goroutinue都在等管道关闭。一旦关闭了其中一个，那么orDone就会关闭，调用orDone的地方就不会阻塞了。
		case 2: // 2个也是一种特殊情况
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default: //超过两个，二分法递归处理
			m := len(channels) / 2
			select {
			case <-or(channels[:m]...):
			case <-or(channels[m:]...):
			}
		}
	}()

	return orDone
}

//测试

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()

	<-or(
		sig(10*time.Second),
		sig(20*time.Second),
		sig(30*time.Second),
		sig(40*time.Second),
		sig(50*time.Second),
		sig(01*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))
}

//实现2

func or2(channels ...<-chan interface{}) <-chan interface{} {
	//特殊情况，只有0个或者1个
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		// 利用反射构建SelectCase
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		// 随机选择一个可用的case
		reflect.Select(cases)
	}()

	return orDone
}
