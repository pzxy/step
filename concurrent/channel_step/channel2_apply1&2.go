package main

import (
	"fmt"
	"time"
)

/**
典型应用场景1：
消息交流

他们将用户的请求放在一个 chan Job 中，这个 chan Job 就相当于一个待处理任务队列。
除此之外，还有一个 chan chan Job 队列，用来存放可以处理任务的 worker 的缓存队列。

类似以前看过的那个爬虫项目
*/

/**
场景2：
数据传递

有一道经典的使用 Channel 进行任务编排的题，
你可以尝试做一下：有四个 goroutine，编号为 1、2、3、4。
每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，
让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
*/

type Token struct{}

func newWorker(id int, ch chan Token, nextCh chan Token) {
	for {
		token := <-ch         // 取得令牌
		fmt.Println((id + 1)) // id从1开始
		time.Sleep(time.Second)
		nextCh <- token
	}
}
func main() {
	chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}

	// 创建4个worker
	for i := 0; i < 4; i++ {
		go newWorker(i, chs[i], chs[(i+1)%4])
	}

	//首先把令牌交给第一个worker
	chs[0] <- struct{}{}

	select {}
}
