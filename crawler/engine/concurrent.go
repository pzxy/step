package engine

import "log"

type ConcurrentEngine struct { //并发引擎结构体 两个字段:1调度字段 2管道数
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface { //调度器接口 向别的地方分发    SimplerScheduler实现了 这个接口
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item: #%d:%v", itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
