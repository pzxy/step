package parser

import (
	"fmt"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func printCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)        //匹配规则
	matches := re.FindAllSubmatch(contents, -1) //匹配返回[][][]byte   []byte相当于string
	result := engine.ParseResult{}
	for _, m := range matches { //便利生成三层数组
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	fmt.Printf("Matches found : %d\n", len(matches))
}
