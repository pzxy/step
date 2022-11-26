package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sync"
)

//去面试了,那人让拿出json数据中所有uri,他的答案是用匿名结构体的思想去做,我用的匹配正则表达式的思想,他说我不对,
//1他说换行匹配的问题,问题是在json文件中,特别的goland中的json文件,是不让换行的
//2他说数据最后处理的问题,放到map中的就迎刃而解了
//3他说一看我代码就没错误处理方式不对,我当时被他虎住了没想那么多,回来想了想,写个demo还要多么严禁的错误处理.真是上纲上线
//4他说用什么方法都可以,我用的这个他说不行,真是有意思

var wg sync.WaitGroup

func main() {
	data, err := ioutil.ReadFile("/ago/pzxy/WorkSpace/Go/src/step/basic/json/re.json")
	if err != nil {
		panic(err)
	}
	//配置正则表达式
	reAttach := regexp.MustCompile(`"(attachPath)":"([^"]+)"`)
	str := reAttach.FindAllStringSubmatch(string(data), -1)
	reNameAttach := regexp.MustCompile(`"(attachName)":"([^"]+)"`)
	str2 := reNameAttach.FindAllStringSubmatch(string(data), -1)
	op(str, str2)
	wg.Wait()
}

func onlyUri(str [][]string) {
	var attachMap sync.Map
	for _, v := range str {
		attachMap.Store(v[2], v[1])
	}
	attachMap.Range(func(key, value interface{}) bool {
		fmt.Println(value, key)
		return true
	})
	wg.Done()
}
func op(str ...[][]string) {
	var allStr [][][]string
	allStr = append(allStr, str...)
	for _, v := range allStr {
		wg.Add(1)
		go onlyUri(v)
	}
}
