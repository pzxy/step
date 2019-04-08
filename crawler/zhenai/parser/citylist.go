package parser

import (
	"godemo/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func PrintCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)        //匹配规则
	matches := re.FindAllSubmatch(contents, -1) //匹配返回[][][]byte   []byte相当于string
	result := engine.ParseResult{}
	limit := 1
	for _, m := range matches { //便利生成三层数组
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity, //nil不安全
		})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
