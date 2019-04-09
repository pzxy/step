package main

import (
	"godemo/crawler/engine"
	"godemo/crawler/persist"
	"godemo/crawler/scheduler"
	"godemo/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{}, //QueuedScheduler实现了Scheduler接口的全部的方法
		WorkerCount: 10,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.PrintCityList,
	})
}
