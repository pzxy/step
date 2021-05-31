package main

import (
	"fmt"
	"reflect"
)

/**
我来借助一个例子，来演示一下，动态处理两个 chan 的情形。
因为这样的方式可以动态处理 case 数据，所以，你可以传入几百几千几万的 chan，
这就解决了不能动态处理 n 个 chan 的问题。
*/

func main() {
	var ch1 = make(chan int, 10)
	var ch2 = make(chan int, 10)

	// 创建SelectCase
	var cases = createCases(ch1, ch2)

	// 执行10次select
	for i := 0; i < 10; i++ {
		chosen, recv, ok := reflect.Select(cases) //这里是随机选的吗？
		if recv.IsValid() {                       // recv case
			fmt.Println("recv:", cases[chosen].Dir, recv, ok)
		} else { // send case
			fmt.Println("send:", cases[chosen].Dir, ok)
		}
	}
}

func createCases(chs ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase

	// 创建recv case
	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	// 创建send case
	for i, ch := range chs {
		v := reflect.ValueOf(i)
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: v,
		})
	}

	return cases
}

//func SeqSendEventReflectStep(m *manager, dic *di.Container) {
//	fmt.Println("enter SeqSendEventStep")
//	m.wg.Add(1)
//	defer m.wg.Done()
//	defer m.pool.Release()
//	_ = bootstrapContainer.LoggingClientFrom(dic.Get)
//
//	var cases []reflect.SelectCase
//
//	m.hashSeqEventMap.Range(func(key, value interface{}) bool {
//		cases = append(cases, reflect.SelectCase{
//			Dir:  reflect.SelectRecv,
//			Chan: reflect.ValueOf(value),
//		})
//		return true
//	})
//
//	for {
//		select {
//		case <-m.ctx.Done():
//			return
//		default:
//
//		}
//		_, recv, ok1 := reflect.Select(cases) //这里是随机选的吗？
//		if !recv.IsValid() {
//			continue
//		}
//		if !ok1 {
//			continue
//		}
//		event, ok2 := recv.Interface().(*SeqEvent)
//		if !ok2 {
//			continue
//		}
//		if len(event.EventCh) <= 0 {
//			continue
//		}
//		m.pool.JobQueue <- func() {
//			event.Lock()
//			defer event.Unlock()
//			if len(event.EventCh) <= 0 {
//				return
//			}
//			if e, ok := <-event.EventCh; ok {
//				i := rand.Intn(1000)
//				time.Sleep(time.Duration(i) * time.Millisecond)
//				fmt.Printf("chanLen:%d, dp:%s, eventId:%s \n\t", len(event.EventCh), e.Readings[0].ResourceName, e.Id)
//			}
//		}
//	}
//
//}
