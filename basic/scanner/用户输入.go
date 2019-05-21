package main

import (
	"fmt"
)

func main() {
	var name, age string
	fmt.Println("输入姓名年龄..")
	//fmt.Scanln(&name,&age)//控制台 必须在一行输入
	fmt.Scanf("%s\n%s", &name, &age) //控制台 输入回车  输入回车
	fmt.Println("输入内容为", name, age)

}
