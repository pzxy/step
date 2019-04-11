package parser

import (
	"godemo/crawler/engine"
	"godemo/crawler/model"
	"regexp"
	"strconv"
)

/*
<a href="http://album.zhenai.com/u/1212636832" target="_blank">渔人不悔改</a>
<td width="180"><span class="grayL">性别：</span>男士</td>
<td><span class="grayL">居住地：</span>四川阿坝</td>
<td width="180"><span class="grayL">年龄：</span>34</td>
<td><span class="grayL">月   薪：</span>12001-20000元</td>
<td width="180"><span class="grayL">婚况：</span>离异</td>
<td width="180"><span class="grayL">身   高：</span>172</td>
*/

/**
<div class="purple-btns" data-v-bff6f798="">
<div class="m-btn purple" data-v-bff6f798>[^<]+</div>
<div class="m-btn purple" data-v-bff6f798>170cm</div>
	<div class="m-btn purple" data-v-bff6f798>未婚</div>
	<div class="m-btn purple" data-v-bff6f798>30岁</div>
	<div class="m-btn purple" data-v-bff6f798>射手座(11.22-12.21)</div>
	<div class="m-btn purple" data-v-bff6f798>170cm</div>
	<div class="m-btn purple" data-v-bff6f798>58kg</div>
	<div class="m-btn purple" data-v-bff6f798>工作地:阿坝红原</div>
	<div class="m-btn purple" data-v-bff6f798>月收入:3千以下</div>
	<div class="m-btn purple" data-v-bff6f798>操作工人</div>
	<div class="m-btn purple" data-v-bff6f798>高中及以下</div>
</div>
*/
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div>`)                                                                                         //年龄
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([0-9]+)cm</div>`)                                                                                    //身高
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)kg</div>`)                                                                                     //体重
var marriageRe = regexp.MustCompile(`<div class="m-content-box" data-v-bff6f798><div class="purple-btns" data-v-bff6f798><div class="m-btn purple" data-v-bff6f798>([^<]+)</div>`) //婚况
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	/**
	匹配整数部分
	*/
	profile := model.Profile{}
	profile.Name = name                                     //名字
	age, err := strconv.Atoi(extractSting(contents, ageRe)) //年龄
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractSting(contents, heightRe)) //身高
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractSting(contents, weightRe)) //体重
	if err == nil {
		profile.Weight = weight
	}

	/**
	匹配字符部分
	*/
	profile.Marriage = extractSting(contents, marriageRe) //婚况
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractSting([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}
	return result
}

func extractSting(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func ProfileParser(name string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, url, name)
	}
}
