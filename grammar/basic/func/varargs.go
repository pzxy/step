package main

import (
	"bytes"
	"fmt"
)

/**
可变参数
*/
//例如:
//fmt.Println() 全部可变
//fmt.Printf() 部分可变

//实现字符串拼接
func joinStrings(s ...string) string {
	//创建缓冲流
	var b bytes.Buffer
	for _, val := range s {
		//将值读入缓冲流
		b.WriteString(val)
	}
	//转换成字符串返回
	return b.String()
}
func main() {
	fmt.Println(joinStrings("白马啸西风", "雪中悍刀行"))
}
