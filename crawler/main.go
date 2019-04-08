package main

import (
	"godemo/crawler/engine"
	"godemo/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.PrintCityList,
	})
}
