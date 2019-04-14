package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Id   int
	Name string
}

func main() {
	//利用反射判断变量类型
	s := "change"
	fmt.Println(reflect.TypeOf(s))

	//获取结构体类型
	peo := People{23, "wonkung"}
	//获取peo的值
	v := reflect.ValueOf(peo)
	//获取属性个数  若v不是结构体类型 panic
	fmt.Println(v.NumField())
}
