package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Id   int
	Name string `xml:name`
}

func main() {
	//利用反射判断变量类型
	s := "change"
	fmt.Println(reflect.TypeOf(s))

	//获取结构体类型
	peo := People{23, "wonkung"}
	//获取peo的值
	v := reflect.ValueOf(peo)
	fmt.Println("v", v)

	//获取属性个数  若v不是结构体类型 panic
	fmt.Println(v.NumField())
	//获取指定位置值
	fmt.Println("获取指定位置值", v.FieldByIndex([]int{0}))
	//获取指定类型
	fmt.Println(v.FieldByName("Name"))

	//设置属性的值
	peo1 := new(People)                          //这是指针类型
	v1 := reflect.ValueOf(peo1).Elem()           //获取指针类型的值
	fmt.Println(v1.FieldByName("Name").CanSet()) //能否被设置
	if v1.FieldByName("Name").CanSet() {
		v1.FieldByName("Name").SetString("王阳明")
	}
	if v1.FieldByName("Id").CanSet() {
		v1.FieldByName("Id").SetInt(99)
	}
	fmt.Println(peo1)

	//获取标记
	t := reflect.TypeOf(People{}) //获取结构体类型
	fmt.Println(t.FieldByName("Name"))
	name, _ := t.FieldByName("Name")
	fmt.Println(name.Tag) //xml:name
	fmt.Println(name.Tag.Get("xml"))
}
