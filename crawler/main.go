package main

import (
	"godemo/crawler/engine"
	"godemo/crawler/scheduler"
	"godemo/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimplerScheduler{}, //scheduler包下的SimplerScheduler实现了Scheduler接口
		WorkerCount: 10,
	}
	e.Run(engine.Request{ //这里调用engine.ConcurrentEngine中的run()方法
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.PrintCityList,
	})
}
