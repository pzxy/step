package main

import "fmt"


func main(){
	wtf()
}

func wtf(){
	s := [...]string{//还能这样写? 这是表
		0:"00",
		1:"11",
		2:"22",
		3:"33",
		4:"44",
	}
	fmt.Println(s)
	//[0 1 2 3 4]

	for i := range s {//还能这样写？ 只返回一个的话，是索引
		fmt.Println(i)
	}
}
