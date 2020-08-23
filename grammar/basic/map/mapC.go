package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	map2bytedemo()
}

func demo1() {
	//初始化

	m := make(map[string]string) //make创建的值为不为nil 是有地址空间的 空值
	fmt.Println(m == nil)
	fmt.Printf("%p\n", m)

	mZ := map[string]string{"one": "矛盾", "two": "幽默", "three": "变化"}
	mX := map[string]string{
		"one":   "矛盾",
		"two":   "幽默",
		"three": "变化", //多了一个点 go编译器只看行 这里没","会找不到上面 "{"对应的"}"了, 加上","是告诉编译器下面还有属于这个模块的东西
	}
	fmt.Println("mZ :", mZ)
	fmt.Println("mX :", mX)

	//增 改

	mZ["one"] = "没有矛盾"
	mZ["four"] = "mount"
	fmt.Println("mZ :", mZ)

	//删

	delete(mZ, "four")
	fmt.Println("mZ :", mZ)

	// 查
	value, ok := mZ["one"]
	fmt.Println(value, ok) //没有返回空和false

	//遍历
	for key, value2 := range mX {
		fmt.Println(key, value2)
	}
}

func map2bytedemo() {
	m := make(map[int]int)
	m[1] = 2
	m[2] = 89
	s, _ := json.Marshal(m)
	fmt.Printf("%v \n", m)
	fmt.Printf("%v \n", s)

}
