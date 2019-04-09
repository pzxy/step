package parser

import (
	"godemo/crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="http://www.zhenai.com/zhenghun/meizhou/[^"]+"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	//re := regexp.MustCompile(profileRe)
	matches := profileRe.FindAllSubmatch(contents, -1)
	//用户 信息
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult { return ParseProfile(c, name) }, //匿名函数
		})
	}
	//下一页
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
