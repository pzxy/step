package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"runtime"
	"time"
)

/**
官方提供的信号量的实现。

如果你的资源数量并不是固定的，而是动态变化的，我建议你考虑一下这个信号量库。
*/
//https://godoc.org/golang.org/x/sync/semaphore
/**
在这段代码中，main goroutine 相当于一个 dispacher，
负责任务的分发。它先请求信号量，如果获取成功，就会启动一个 goroutine 去处理计算，
然后，这个 goroutine 会释放这个信号量（有意思的是，信号量的获取是在 main goroutine，
信号量的释放是在 worker goroutine 中），如果获取不成功，就等到有信号量可以使用的时候，再去获取。

需要提醒你的是，其实，在这个例子中，还有一个值得我们学习的知识点，就是最后的那一段处理（第 25 行）。
如果在实际应用中，你想等所有的 Worker 都执行完，就可以获取最大计数值的信号量。
*/
var (
	maxWorkers = runtime.GOMAXPROCS(0)                    // worker数量
	sema       = semaphore.NewWeighted(int64(maxWorkers)) //信号量
	task       = make([]int, maxWorkers*4)                // 任务数，是worker的四倍
)

func main() {
	ctx := context.Background()
	fmt.Println(maxWorkers, task)
	for i := range task {
		// 如果没有worker可用，会阻塞在这里，直到某个worker被释放
		if err := sema.Acquire(ctx, 1); err != nil {
			break
		}

		// 启动worker goroutine
		go func(i int) {
			defer sema.Release(1)
			time.Sleep(100 * time.Millisecond) // 模拟一个耗时操作
			task[i] = i + 1
		}(i)
	}

	// 请求所有的worker,这样能确保前面的worker都执行完
	if err := sema.Acquire(ctx, int64(maxWorkers)); err != nil { //可以一次请求多个资源。
		log.Printf("获取所有的worker失败: %v", err)
	}

	fmt.Println(task)
}
