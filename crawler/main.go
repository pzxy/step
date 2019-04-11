package main

import (
	"godemo/crawler/engine"
	"godemo/crawler/persist"
	"godemo/crawler/scheduler"
	"godemo/crawler/zhenai/parser"
)

func main() {
	itemSaver, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(nil)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{}, //QueuedScheduler实现了Scheduler接口的全部的方法
		WorkerCount: 10,
		ItemChan:    itemSaver,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.PrintCityList,
	})
}
