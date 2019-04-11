package parser

import (
	"godemo/crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="http://www.zhenai.com/zhenghun/meizhou/[^"]+"`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	//re := regexp.MustCompile(profileRe)
	matches := profileRe.FindAllSubmatch(contents, -1)
	//用户 信息
	result := engine.ParseResult{}
	for _, m := range matches {
		//name := string(m[2])
		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ProfileParser(string(m[2])), //这里用到了函数 所以不用先copy一个name了  因为若不是取地址的话 本身就是copy一个作为入参
		})
	}
	/*//下一页
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}*/

	return result
}
