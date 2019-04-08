package main

import (
	"godemo/crawler/engine"
	"godemo/crawler/scheduler"
	"godemo/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimplerScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.PrintCityList,
	})
}
