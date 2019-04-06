package main

import (
	"fmt"
	"regexp"
)

const text = `My email is wonkung@163.com
My email is wonasdng@163
My email is wonasdng@163.cm.cn`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`) //must含义 写的肯定是对的 不会出错 表达式加()表示提取
	//match := re.FindAllString(text,-1)//findstring 说明参数是string
	match := re.FindAllStringSubmatch(text, -1) //提取regex中()的数据

	fmt.Println(match)
}
