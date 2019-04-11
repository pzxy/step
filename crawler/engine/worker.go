package engine

import (
	"godemo/crawler/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url) //通过url返回全部数据
	if err != nil {                   //错误处理 避免一个爬去错误影响后续的请求
		log.Printf("Fetcher :error "+"fetching url %s,%v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body, r.Url), nil //处理数据 request 放到requests中 item打印出来
}
