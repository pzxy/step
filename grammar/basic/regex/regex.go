package main

import (
	"fmt"
	"regexp"
)

const text = `My email is wonkung@163.com
My email is wonasdng@163
My email is wonasdng@163.cm.cn`
const (
	url2 = `http://update.file.test.chipparking.com/robot/robot.bin?e=1571383314&token=lg7QUwopX7lpjq6xH4O98cocrkiTWYVo0W2PjJF7:eAl1ge1o2Um98zKGLItDuO5Ipn4=`
	url3 = `D://dt_test_file.bin`
)

func main() {
	regex := `(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`
	//math := `([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`
	re := regexp.MustCompile(regex) //must含义 写的肯定是对的 不会出错 表达式加()表示提取
	//match := re.FindAllString(text,-1)//findstring 说明参数是string
	matchs := re.FindAllStringSubmatch(url3, -1) //提取regex中()的数据
	for _, m := range matchs {
		fmt.Println(m[1])
	}
}
