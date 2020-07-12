package main

import "fmt"

func main() {
	s := [...]string{
		0: "00",
		1: "11",
		2: "22",
		3: "33",
		4: "44",
	}
	fmt.Println(cap(s[0:2:2]))
}

func wtf() {
	s := [...]string{ //还能这样写? 这是表
		0: "00",
		1: "11",
		2: "22",
		3: "33",
		4: "44",
	}
	fmt.Println(s)
	//[0 1 2 3 4]

	for i := range s { //还能这样写？ 只返回一个的话，是索引
		fmt.Println(i)
	}
}

func wtf2() {
	s := [...]string{
		0: "00",
		1: "11",
		2: "22",
		3: "33",
		4: "44",
	}
	fmt.Println(s[:2:5]) //表示0-2，5表示cap大小。
}
